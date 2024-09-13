package credly

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

const credlyBaseURL = "https://www.credly.com/"

type Badge struct {
	ImageSrc string
	Alt      string
}

type Credly struct {
	baseURL string
	client  http.Client
}

func NewClient() *Credly {
	return &Credly{
		baseURL: credlyBaseURL,
		client:  http.Client{},
	}
}

func (c *Credly) WithHTTPClient(client http.Client) *Credly {
	c.client = client
	return c
}

func (c *Credly) WithBaseURL(baseURL string) *Credly {
	c.baseURL = baseURL
	return c
}

// FetchCredlyUserPage fetches the Credly user page for the provided username.
func (c *Credly) FetchUserPage(ctx context.Context, username string) ([]byte, error) {
	urlString := c.baseURL + "users/" + username + "/badges"

	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, parsedURL.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
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

// ExtractBadges extracts the Credly badges from the provided HTML body.
func ExtractBadges(htmlBody []byte) ([]Badge, error) {
	doc, err := html.Parse(strings.NewReader(string(htmlBody)))
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
