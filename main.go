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

	"golang.org/x/net/html"
)

const BaseURL = "https://www.credly.com/"

var (
	username string
)

type Badge struct {
	ImageSrc string
	Alt      string
}

func main() {
	flag.StringVar(&username, "username", "", "Credly username")
	flag.Parse()

	if username == "" {
		username = os.Getenv("CREDLY_USERNAME")
		if username == "" {
			log.Fatal("Username is not provided. Please provide it as a command-line argument or set the CREDLY_USERNAME environment variable.")
		}
	}

	urlString := BaseURL + "users/" + username + "/badges"
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	req, err := http.NewRequest(http.MethodGet, parsedURL.String(), http.NoBody)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
							fmt.Printf("img tag: %+v\n", c) // print the img tag
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

  fmt.Println(badges)
}
