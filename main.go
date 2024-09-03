package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v64/github"
	"golang.org/x/net/html"
)

const credlyBaseURL = "https://www.credly.com/"

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
	content, _, _, err := ghClient.Repositories.GetContents(ctx, ghUsername, ghUsername, "README.md", nil)
	if err != nil {
		log.Fatal(err)
	}

	readme, err := content.GetContent()
	if err != nil {
		log.Fatal(err)
	}

	// Find the badges section in the README.md
	startIndex := strings.Index(readme, "<!--START_BADGES:badges-->")
	if startIndex == -1 {
		log.Fatal("START_BADGES section not found in README.md")
	}

	endIndex := strings.Index(readme, "<!--END_BADGES:badges-->")
	if endIndex == -1 {
		log.Fatal("END_BADGES section not found in README.md")
	}

	// Adjust endIndex to point to the end of the <!--END_BADGES:badges--> tag
	endIndex += len("<!--END_BADGES:badges-->")

	urlString := credlyBaseURL + "users/" + credlyUsername + "/badges"
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodGet, parsedURL.String(), http.NoBody)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		log.Fatal(err)
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

	var badgeMarkdown strings.Builder
	for _, badge := range badges {
		badgeMarkdown.WriteString(fmt.Sprintf("<img src=\"%s\" alt=\"%s\" />\n", badge.ImageSrc, badge.Alt))
	}

	readme = readme[:startIndex] + "<!--START_BADGES:badges-->\n" + badgeMarkdown.String() + "<!--END_BADGES:badges-->" + readme[endIndex:]

	_, _, err = ghClient.Repositories.UpdateFile(ctx, ghUsername, ghUsername, "README.md", &github.RepositoryContentFileOptions{
		Branch:  github.String("main"),
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

	log.Println("Badges updated successfully!")
}
