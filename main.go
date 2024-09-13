package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v64/github"
	"golang.org/x/net/html"
)

const (
	credlyBaseURL = "https://www.credly.com/"
	readmeFile    = "README.md"
	badgeStart    = "<!--START_BADGES:badges-->"
	badgeEnd      = "<!--END_BADGES:badges-->"
	defaultBranch = "main"
)

var ErrFilesAreEqual = errors.New("files are equal")

var (
	credlyUsername string
	ghToken        string
	ghUsername     string
	commitMessage  string
)

type Badge struct {
	ImageSrc string
	Alt      string
}

func main() {
	flag.StringVar(&credlyUsername, "credly-username", "", "Credly username")
	flag.StringVar(&ghToken, "gh-token", "", "GitHub token")
	flag.StringVar(&ghUsername, "gh-username", "", "GitHub username")
	flag.StringVar(&commitMessage, "commit-message", "Update Credly badges!", "Commit message")
	flag.Parse()

	if credlyUsername == "" {
		credlyUsername = os.Getenv("INPUT_CREDLY_USERNAME")
		if credlyUsername == "" {
			log.Fatal("Username is not provided. Please provide it as a command-line argument or set the CREDLY_USERNAME environment variable.")
		}
	}

	if ghToken == "" {
		ghToken = os.Getenv("INPUT_GITHUB_TOKEN")
		if ghToken == "" {
			log.Fatal("GitHub token is not provided. Please provide it as a command-line argument or set the GITHUB_TOKEN environment variable.")
		}
	}

	if ghUsername == "" {
		ghUsername = os.Getenv("GITHUB_ACTOR")
		if ghUsername == "" {
			log.Fatal("GitHub username is not provided. Please provide it as a command-line argument or set the GITHUB_USERNAME environment variable.")
		}
	}

	ghClient := github.NewClient(nil).WithAuthToken(ghToken)

	ctx := context.Background()

	// TODO: This only works on personal GitHub spaces. Fix so that we can get the README from any repository
	content, _, _, err := ghClient.Repositories.GetContents(ctx, ghUsername, ghUsername, readmeFile, nil)
	if err != nil {
		log.Fatal(err)
	}

	readme, err := content.GetContent()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	body, err := fetchCredlyUserPage(ctx, credlyUsername)
	if err != nil {
		log.Fatal(err)

	}

	badges, err := extractBadges(body)
	if err != nil {
		log.Fatal(err)
	}

	readme, err = updateReadme(readme, badges)
	if err != nil {
		if errors.Is(err, ErrFilesAreEqual) {
			log.Printf("No changes between the fetched %s and the updated detected. Exiting...", readmeFile)
			return
		}
		log.Fatal(err)
	}

	// Commit the changes
	_, _, err = ghClient.Repositories.UpdateFile(ctx, ghUsername, ghUsername, readmeFile, &github.RepositoryContentFileOptions{
		Branch:  github.String(defaultBranch),
		Message: github.String(commitMessage),
		Committer: &github.CommitAuthor{
			Name:  github.String("github-actions[bot]"),
			Email: github.String("41898282+github-actions[bot]@users.noreply.github.com"),
		},
		Author: &github.CommitAuthor{
			Name:  github.String("github-actions[bot]"),
			Email: github.String("41898282+github-actions[bot]@users.noreply.github.com"),
		},
		Content: []byte(readme),
		SHA:     content.SHA,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Credly badges in %s updated successfully!", readmeFile)
}

// updateReadme updates the README.md file with the provided badges
func updateReadme(readme string, badges []Badge) (string, error) {
	originalReadme := readme

	var badgeMarkdown strings.Builder
	for _, badge := range badges {
		badgeMarkdown.WriteString(fmt.Sprintf("<img src=\"%s\" alt=\"%s\" />\n", badge.ImageSrc, badge.Alt))
	}

	startIndex, endIndex, err := findStartAndEndIndex(readme)
	if err != nil {
		return "", err
	}

	readme = readme[:startIndex] + badgeStart + "\n" + badgeMarkdown.String() + badgeEnd + readme[endIndex:]

	if originalReadme == readme {
		return originalReadme, ErrFilesAreEqual
	}

	return readme, nil
}

// extractBadges extracts the badges from the provided Credly user page
func extractBadges(b []byte) ([]Badge, error) {
	doc, err := html.Parse(strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}

	var badges []Badge

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "cr-standard-grid-item-content c-badge c-badge--medium" {
					var badge Badge
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						if c.Type == html.ElementNode && c.Data == "img" {
							for _, a := range c.Attr {
								if a.Key == "src" {
									badge.ImageSrc = a.Val
								}
								if a.Key == "alt" {
									badge.Alt = a.Val
								}
							}
						}
					}
					badges = append(badges, badge)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return badges, nil
}

// findStartAndEndIndex finds the start and end index of the badges section in the README.md
func findStartAndEndIndex(readme string) (start, end int, err error) {
	startIndex := strings.Index(readme, badgeStart)
	if startIndex == -1 {
		return 0, 0, fmt.Errorf("%s not found in %s", badgeStart, readmeFile)
	}

	endIndex := strings.Index(readme, badgeEnd)
	if endIndex == -1 {
		return 0, 0, fmt.Errorf("%s section not found in %s", badgeEnd, readmeFile)
	}

	endIndex += len(badgeEnd)

	return startIndex, endIndex, nil
}

// fetchCredlyUserPage fetches the Credly user page for the provided username
func fetchCredlyUserPage(ctx context.Context, username string) ([]byte, error) {
	urlString := credlyBaseURL + "users/" + username + "/badges"

	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, parsedURL.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("failed to fetch Credly user (%s) page: %s", username, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
