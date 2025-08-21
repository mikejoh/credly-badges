package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

// BadgeInfo will hold the extracted badge data
type BadgeInfo struct {
	Title    string
	ImageURL string
	// Add other fields you want to extract, e.g., Issuer, Date
}

// randomSleep pauses execution for a random duration to mimic human behavior
func randomSleep(min, max int) {
	duration := time.Duration(rand.Intn(max-min+1)+min) * time.Millisecond
	time.Sleep(duration)
}

func main() {
	// Create a new context
	opts := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3830.0 Safari/537.36"),
		chromedp.WindowSize(1920, 1080),
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Headless,
		chromedp.DisableGPU,
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	url := "https://www.credly.com/users/mikael-johansson-2" // It's often good to go directly to the badges page if available

	var badges []BadgeInfo
	log.Println("Navigating to:", url)

	// Selectors for the page elements
	allBadgesContainerSelector := ".skills-profile__badge-portfolio-section__view_portfolio_container"   // Selector for the main container holding all badges
	badgeContainerSelector := ".settings__skills-profile__edit-skills-profile__badge-card"               // Selector for individual badge containers
	titleSelector := ".settings__skills-profile__edit-skills-profile__badge-card__badge-title"           // Selector for title within a badge container
	imageSelector := ".settings__skills-profile__edit-skills-profile__badge-card__badge-image-container" // Selector for badge image container
	cookieAcceptButtonSelector := "#onetrust-accept-btn-handler"                                         // Common ID for a cookie accept button

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		// Try to click the cookie accept button if it exists. Use a short timeout so it doesn't block if not present.
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("Checking for cookie consent dialog...")
			cookieCtx, cookieCancel := context.WithTimeout(ctx, 5*time.Second)
			defer cookieCancel()
			clickErr := chromedp.Run(cookieCtx, chromedp.Click(cookieAcceptButtonSelector, chromedp.ByQuery))
			if clickErr != nil {
				log.Printf("Failed to click cookie accept button (might not be present): %v", clickErr)
			} else {
				log.Println("Cookie accept button clicked.")
				randomSleep(1000, 3000) // Small delay after clicking
			}
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			randomSleep(2000, 5000) // Random delay after initial navigation/cookie handling
			return nil
		}),
		// Wait for the main container of all badges to be visible
		chromedp.WaitVisible(allBadgesContainerSelector, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("Page loaded, attempting to extract badge nodes...")

			// First, let's find all badge containers
			var badgeContainers []*cdp.Node
			if err := chromedp.Nodes(badgeContainerSelector, &badgeContainers, chromedp.ByQueryAll).Do(ctx); err != nil {
				log.Printf("Could not get badge containers: %v", err)
				return err
			}
			log.Printf("Found %d potential badge elements.", len(badgeContainers))

			if len(badgeContainers) == 0 {
				log.Println("No badge elements found with the selector:", badgeContainerSelector)
				// Let's try to get a screenshot for debugging
				var buf []byte
				if err := chromedp.CaptureScreenshot(&buf).Do(ctx); err != nil {
					log.Printf("Failed to capture screenshot: %v", err)
				} else {
					if err := os.WriteFile("screenshot.png", buf, 0o644); err != nil {
						log.Printf("Failed to save screenshot: %v", err)
					} else {
						log.Println("Screenshot saved to screenshot.png")
					}
				}
				return nil
			}

			for _, container := range badgeContainers {
				var title, imgURL string

				// Extract title
				if err := chromedp.Text(titleSelector, &title, chromedp.ByQuery, chromedp.FromNode(container)).Do(ctx); err != nil {
					log.Printf("Could not get title for a badge: %v", err)
				}

				// Extract image URL
				if err := chromedp.AttributeValue(imageSelector+" img", "src", &imgURL, nil, chromedp.ByQuery, chromedp.FromNode(container)).Do(ctx); err != nil {
					log.Printf("Could not get image URL for a badge: %v", err)
				}

				if title != "" { // Only add if we at least got a title
					badges = append(badges, BadgeInfo{Title: title, ImageURL: imgURL})
					log.Printf("Extracted badge: Title='%s', ImageURL='%s'", title, imgURL)
				} else {
					log.Println("Skipping a badge due to missing title.")
				}
			}
			return nil
		}),
	)

	if err != nil {
		log.Fatalf("Failed to fetch badges: %v", err)
	}

	log.Printf("Successfully extracted %d badges:", len(badges))
	for i, badge := range badges {
		log.Printf("%d: %s - %s", i+1, badge.Title, badge.ImageURL)
	}
}
