package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"

	"github.com/mikejoh/go-credly/internal/credly"
	"github.com/mikejoh/go-credly/internal/readme"
)

type credlyBadgesOptions struct {
	credlyUsername string
	ghToken        string
	ghUsername     string
	branch         string
	commitMessage  string
}

func main() {
	cdOpts := &credlyBadgesOptions{}

	flag.StringVar(&cdOpts.credlyUsername, "credly-username", "", "Credly username")
	flag.StringVar(&cdOpts.ghToken, "gh-token", "", "GitHub token")
	flag.StringVar(&cdOpts.ghUsername, "gh-username", "", "GitHub username")
	flag.StringVar(&cdOpts.branch, "branch", "main", "Branch to commit the changes")
	flag.StringVar(&cdOpts.commitMessage, "commit-message", "Update Credly badges!", "Commit message")
	flag.Parse()

	if cdOpts.credlyUsername == "" {
		cdOpts.credlyUsername = os.Getenv("INPUT_CREDLY_USERNAME")
		if cdOpts.credlyUsername == "" {
			log.Fatal("Username is not provided. Please provide it as a command-line argument or set the CREDLY_USERNAME environment variable.")
		}
	}

	if cdOpts.ghToken == "" {
		cdOpts.ghToken = os.Getenv("INPUT_GITHUB_TOKEN")
		if cdOpts.ghToken == "" {
			log.Fatal("GitHub token is not provided. Please provide it as a command-line argument or set the GITHUB_TOKEN environment variable.")
		}
	}

	if cdOpts.ghUsername == "" {
		cdOpts.ghUsername = os.Getenv("GITHUB_ACTOR")
		if cdOpts.ghUsername == "" {
			log.Fatal("GitHub username is not provided. Please provide it as a command-line argument or set the GITHUB_USERNAME environment variable.")
		}
	}

	if cdOpts.branch == "" {
		cdOpts.branch = "main"
	}

	ctx := context.Background()

	credlyClient := credly.NewClient()
	profileReadme := readme.NewReadme(cdOpts.ghUsername, cdOpts.ghUsername)

	err := profileReadme.Fetch(ctx)
	if err != nil {
		log.Fatal(err)
	}

	body, err := credlyClient.FetchUserPage(ctx, cdOpts.credlyUsername)
	if err != nil {
		log.Fatal(err)

	}

	badges, err := credly.ExtractBadges(body)
	if err != nil {
		log.Fatal(err)
	}

	if len(badges) == 0 {
		log.Fatalf("no badges found for the provided username %s. Exiting...", cdOpts.credlyUsername)
	}

	err = profileReadme.WriteBadges(badges)
	if err != nil {
		if errors.Is(err, readme.ErrFilesAreEqual) {
			log.Printf("no changes between the fetched %s and the updated detected. Exiting...", profileReadme.Filename())
			return
		}
		log.Fatal(err)
	}

	err = profileReadme.Update(ctx, "main")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("credly badges in %s updated successfully!", profileReadme.Filename())
}
