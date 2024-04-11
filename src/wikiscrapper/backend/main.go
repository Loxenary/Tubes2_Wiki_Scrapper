package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "github.com/PuerkitoBio/goquery"
    "os"
    "time"
)

// type Root struct{
//     ID int
//     Link string
//     Parent int
// }

func main() {

    clearFile("links.txt")
    clearFile("output.txt")

    visited := make(map[string]bool)
    // URL of the web page to scrape
    url := "/wiki/Platter_(dishware)"
    target := "/wiki/Jewellery"
    start := time.Now()
    //path := BFS(url,target)
    path := DLS(url,target,4,visited)
    writeFile("output.txt",path)
    run_time := time.Since(start)
    
    fmt.Println("runtime :" , run_time)
}

func getListofLinks(url string,visited map[string]bool) []string{
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

    // Extract title
    //title := doc.Find("title").Text()
    //fmt.Println("Title:", url)

    // Extract links
    var links []string
    var linkCount int
    doc.Find("#mw-content-text").Each(func(i int, content *goquery.Selection) {
        // Extract links within the main content area
        content.Find("a").Each(func(i int, s *goquery.Selection) {
            // Get the link's href attribute
            link, exists := s.Attr("href")
            
            if exists && strings.HasPrefix(link, "/wiki/") && !ignoreLink(link) && !isin(link,links) && !visited[link] && !strings.ContainsAny(link,"#") {
                // Append the link to the slice
                links = append(links, link)
                linkCount++
                if linkCount >= 50 {
                    return // Exit the loop if 50 links are found
                }
            }
        })
        if linkCount >= 50 {
            return // Exit the loop if 50 links are found
        }
    })

    // fmt.Println("Links:")
    // for _, link := range links {
    //     fmt.Println(link)
    // }
    writeFile("links.txt", links)
    return links
}

func ignoreLink(link string) bool{
    ignoreList := []string{
        "/wiki/File:" ,
        "/wiki/Help:" ,
        "/wiki/Special:" ,
        "/wiki/Template:" ,
        "/wiki/Template_Talk:" ,
        "/wiki/Template_talk:" ,
        "/wiki/Wikipedia:" ,
        "/wiki/Category:",
        "/wiki/Portal:" ,
        "/wiki/User:" ,
        "/wiki/User_Talk:" ,
        "/wiki/Talk:",
    }

    for _, prefix := range ignoreList {
		if strings.HasPrefix(link, prefix) {
			return true
		}
	}
	return false
}


func writeFile(filename string, links []string) {
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        log.Fatal("Error opening file:", err)
    }
    defer file.Close()

    for _, link := range links {
        // Write each link to the file
        _, err := file.WriteString(link + "\n")
        if err != nil {
            log.Fatal("Error writing to file:", err)
        }
    }

    //fmt.Println("Links appended to", filename)
}

func clearFile(filename string) {
    // Open the file with write-only mode and truncate it (clear content)
    file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
    if err != nil {
        log.Fatal("Error clearing file:", err)
    }
    defer file.Close()

    // Truncate the file to clear its content
    err = file.Truncate(0)
    if err != nil {
        log.Fatal("Error truncating file:", err)
    }
}

func isin(link string,array []string) bool{
    for _, item := range array {
        if item == link {
            return true
        }
    }
    return false
}

func BFS(startURL, targetURL string) []string {
    visited := make(map[string]bool)
    queue := [][]string{{startURL}}
    for len(queue) > 0 {
        path := queue[0]
        queue = queue[1:]

        // Get the last visited URL in the path
        currentURL := path[len(path)-1]

        // Check if the target article is reached
        if currentURL == targetURL {
            return path
        }

        // Skip if the URL is already visited
        if visited[currentURL] {
            continue
        }
        visited[currentURL] = true

        // Get links from the current URL
        //println("flag")
        links := getListofLinks(currentURL,visited)

        // Add new paths to the queue
        for _, link := range links {
            newPath := append([]string{}, path...)
            newPath = append(newPath, link)
            queue = append(queue, newPath)
        }
    }
    return nil // Target article not found
}
  

// func IDS(startArticle, targetArticle string, depthLimit int) []string {
//     visited := make(map[string]bool)
//     for depth := 0; depth <= depthLimit; depth++ {
//         path := DLS(startArticle, targetArticle, depth)
//         if path != nil {
//             return path
//         }
//     }
//     return nil
// }

// Function to perform depth-limited search
func DLS(currentArticle, targetArticle string, depthLimit int,visited (map[string]bool)) []string {
    
    if depthLimit == 0 {
        return nil
    }

    if currentArticle == targetArticle {
        return []string{currentArticle}
    }

    if visited[currentArticle] {
        return nil
    }
    visited[currentArticle] = true

    links := getListofLinks(currentArticle,visited)
    for _, link := range links {
        path := DLS(link, targetArticle, depthLimit-1,visited)
        if path != nil {
            return append([]string{currentArticle}, path...)
        }
    }

    return nil
}


