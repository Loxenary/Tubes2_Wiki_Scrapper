package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/PuerkitoBio/goquery"
)

func main() {
    // URL of the web page to scrape
    url := "https://example.com"

    // Send HTTP GET request
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

    // Extract title
    title := doc.Find("title").Text()
    fmt.Println("Title:", title)

    // Extract links
    doc.Find("a").Each(func(i int, s *goquery.Selection) {
        link, _ := s.Attr("href")
        fmt.Println("Link:", link)
    })
}
