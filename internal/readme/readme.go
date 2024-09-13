package readme

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	gh "github.com/google/go-github/v64/github"
	"github.com/mikejoh/go-credly/internal/credly"
)

var ErrFilesAreEqual = errors.New("files are equal")

type GitHubReadme struct {
	githubClient *gh.Client
	repoContent  *gh.RepositoryContent
	readme       string
	fileName     string
	repo         string
	owner        string
	badgeStart   string
	badgeEnd     string
}

func NewReadme(owner, repo string) *GitHubReadme {
	fileName := "README.md"
	badgeStart := "<!--START_BADGES:badges-->"
	badgeEnd := "<!--END_BADGES:badges-->"

	return &GitHubReadme{
		githubClient: gh.NewClient(nil),
		repoContent:  &gh.RepositoryContent{},
		readme:       "",
		fileName:     fileName,
		repo:         repo,
		owner:        owner,
		badgeStart:   badgeStart,
		badgeEnd:     badgeEnd,
	}
}

func (gr *GitHubReadme) WithGitHubClient(client *gh.Client) *GitHubReadme {
	gr.githubClient = client
	return gr
}

func (gr *GitHubReadme) WithFileName(filename string) *GitHubReadme {
	gr.fileName = filename
	return gr
}

func (gr *GitHubReadme) WithBadgeStart(badgeStart string) *GitHubReadme {
	gr.badgeStart = badgeStart
	return gr
}

func (gr *GitHubReadme) WithBadgeEnd(badgeEnd string) *GitHubReadme {
	gr.badgeEnd = badgeEnd
	return gr
}

func (gr *GitHubReadme) Fetch(ctx context.Context) error {
	content, _, _, err := gr.githubClient.Repositories.GetContents(ctx, gr.repo, gr.repo, gr.fileName, nil)
	if err != nil {
		return err
	}

	readmeString, err := content.GetContent()
	if err != nil {
		return err
	}

	gr.repoContent = content
	gr.readme = readmeString

	log.Println("readme content fetched and saved")

	return nil
}

func (gr *GitHubReadme) WriteBadges(badges []credly.Badge) error {
	originalReadme := gr.readme

	var badgeMarkdown strings.Builder
	for _, badge := range badges {
		badgeMarkdown.WriteString(fmt.Sprintf("<img src=\"%s\" alt=\"%s\" />\n", badge.ImageSrc, badge.Alt))
	}

	startIndex, endIndex, err := findStartAndEndIndex(gr.readme, gr.badgeStart, gr.badgeEnd)
	if err != nil {
		return err
	}

	gr.readme = gr.readme[:startIndex] + gr.badgeStart + "\n" + badgeMarkdown.String() + gr.badgeEnd + gr.readme[endIndex:]

	if originalReadme == gr.readme {
		return ErrFilesAreEqual
	}

	return nil
}

func (gr *GitHubReadme) Get() string {
	return gr.readme
}

func (gr *GitHubReadme) Update(ctx context.Context, branch string) error {
	_, _, err := gr.githubClient.Repositories.UpdateFile(ctx, gr.repo, gr.owner, gr.Filename(), &gh.RepositoryContentFileOptions{
		Branch:  gh.String(branch),
		Message: gh.String("Update Credly badges"),
		Committer: &gh.CommitAuthor{
			Name:  gh.String("github-actions[bot]"),
			Email: gh.String("41898282+github-actions[bot]@users.noreply.github.com"),
		},
		Author: &gh.CommitAuthor{
			Name:  gh.String("github-actions[bot]"),
			Email: gh.String("41898282+github-actions[bot]@users.noreply.github.com"),
		},
		Content: []byte(gr.readme),
		SHA:     gr.repoContent.SHA,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("readme updated")

	return nil
}

func (gr *GitHubReadme) Filename() string {
	return gr.fileName
}

func findStartAndEndIndex(readme, startString, endString string) (start, end int, err error) {
	if startString == "" || endString == "" {
		return 0, 0, errors.New("startString and endString cannot be empty")
	}

	if readme == "" {
		return 0, 0, errors.New("readme cannot be empty")
	}

	startIndex := strings.Index(readme, startString)
	if startIndex == -1 {
		return 0, 0, fmt.Errorf("%s not found in readme", startString)
	}

	endIndex := strings.Index(readme, endString)
	if endIndex == -1 {
		return 0, 0, fmt.Errorf("%s not found in readme", endString)
	}

	endIndex += len(endString)

	return startIndex, endIndex, nil
}
