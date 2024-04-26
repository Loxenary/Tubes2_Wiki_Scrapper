package main

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func getListofLinks(targeturl, url string, visited SafeMap) ([]string, bool) {
    url = "https://en.wikipedia.org" + url

    response, err := http.Get(url)
    if err != nil {
        log.Fatal("Error fetching URL:", err)
    }
    defer response.Body.Close()

    // Parse HTML
    doc, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error parsing HTML:", err)
    }

    // Extract links
    var links []string
    targetFound := false // Flag to indicate if the target URL has been found

    doc.Find("#mw-content-text").Each(func(i int, content *goquery.Selection) {
        // Extract links within the main content area
        content.Find("a").Each(func(i int, s *goquery.Selection) {
            // Get the link's href attribute
            link, exists := s.Attr("href")
            if exists && strings.HasPrefix(link, "/wiki/") && !ignoreLink(link) && !isin(link, links) && !visited.Get(link) && !strings.ContainsAny(link, "#") {
                // Append the link to the slice
                links = append(links, link)
                if link == targeturl {
                    // If the link matches the target URL, set the flag
                    targetFound = true
                    //return
                }
            }
        })
    })

    //writeFile("links.txt", links)
    return links, targetFound
}


type linkExtractionResult struct {
	Links       []string
	TargetFound bool
}

// Helper function to extract links from a part of the document
func extractLinks(doc *goquery.Document, visited SafeMap, targeturl string, links *[]string, start, end int) bool {
	var targetFound bool
	doc.Selection.Slice(start, end).Find("#mw-content-text").Each(func(i int, content *goquery.Selection) {
		// Extract links within the main content area
		content.Find("a").Each(func(i int, s *goquery.Selection) {
			// Get the link's href attribute
			link, exists := s.Attr("href")

			if exists && strings.HasPrefix(link, "/wiki/") && !ignoreLink(link) && !isin(link, *links) && !visited.Get(link) && !strings.ContainsAny(link, "#") {
				// Append the link to the slice
				*links = append(*links, link)
				if link == targeturl {
					// If the link matches the target URL, set the flag
					targetFound = true
				}
			}
		})
	})
	return targetFound
}

func getListofLinks1(targeturl, url string, visited SafeMap) ([]string, bool) {
	url = "https://en.wikipedia.org" + url
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching URL:", err)
	}
	defer response.Body.Close()

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}

	// Extract links concurrently
	var wg sync.WaitGroup
	n := 4
	resultChan := make(chan linkExtractionResult, n) // Two parts: top and bottom

	partSize := doc.Length() / n

	// Process each part of the document concurrently
	for i := 0; i < n; i++ {
		start := i * partSize
		end := (i + 1) * partSize
		if i == n-1 {
			end = doc.Length() // Last part may be larger if doc.Length() is not divisible by 4
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			var links []string
			targetFound := extractLinks(doc, visited, targeturl, &links, start, end)
			resultChan <- linkExtractionResult{Links: links, TargetFound: targetFound}
		}(start, end)
	}

	// Wait for all parts to finish and close the result channel
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results
	var allLinks []string
	var targetFound bool
	for result := range resultChan {
		allLinks = append(allLinks, result.Links...)
		if result.TargetFound {
			targetFound = true
		}
	}

	//writeFile("links.txt", allLinks)
	return allLinks, targetFound
}

func getListofLinks2(targeturl, url string) ([]string, bool) {
    url = "https://en.wikipedia.org" + url

    response, err := http.Get(url)
    if err != nil {
        log.Fatal("Error fetching URL:", err)
    }
    defer response.Body.Close()

    // Parse HTML
    doc, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error parsing HTML:", err)
    }

    // Extract links
    var links []string
    targetFound := false // Flag to indicate if the target URL has been found

    doc.Find("#mw-content-text").Each(func(i int, content *goquery.Selection) {
        // Extract links within the main content area
        content.Find("a").Each(func(i int, s *goquery.Selection) {
            // Get the link's href attribute
            link, exists := s.Attr("href")
            if exists && strings.HasPrefix(link, "/wiki/") && !ignoreLink(link) && !isin(link, links) && !strings.ContainsAny(link, "#") {
                // Append the link to the slice
                links = append(links, link)
                if link == targeturl {
                    // If the link matches the target URL, set the flag
                    targetFound = true
                    //return
                }
            }
        })
    })
    writeFile("links.txt",links)
    return links, targetFound
}

