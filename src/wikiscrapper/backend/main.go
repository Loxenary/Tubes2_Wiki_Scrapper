package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"github.com/PuerkitoBio/goquery"
)

// type Root struct{
//     ID int
//     Link string
//     Parent int
// }

func main() {

    clearFile("links.txt")
    clearFile("output.txt")

    // URL of the web page to scrape
    url := "/wiki/Sergey_Bubka"
    target := "/wiki/Twin"
    start := time.Now()
    //path := BFS(url,target)
	path := BFS_con(url,target)
    //path := IDS(url,target,6)
    writeFile("output.txt",path)
    run_time := time.Since(start)

    // visited := make(map[string]bool)
    // start2 := time.Now()
    // getListofLinks2(target,url,visited)
    // run_time2 := time.Since(start2)
    
    fmt.Println("runtime :" , run_time)
}

// func getListofLinks2(targeturl, url string, visited map[string]bool) ([]string, bool) {
//     url = "https://en.wikipedia.org" + url
//     response, err := http.Get(url)
//     if err != nil {
//         log.Fatal("Error fetching URL:", err)
//     }
//     defer response.Body.Close()

//     // Parse HTML
//     doc, err := goquery.NewDocumentFromReader(response.Body)
//     if err != nil {
//         log.Fatal("Error parsing HTML:", err)
//     }

//     // Extract links
//     var links []string
//     targetFound := false // Flag to indicate if the target URL has been found

//     doc.Find("#mw-content-text").Each(func(i int, content *goquery.Selection) {
//         // Extract links within the main content area
//         content.Find("a").Each(func(i int, s *goquery.Selection) {
//             // Get the link's href attribute
//             link, exists := s.Attr("href")

//             if exists && strings.HasPrefix(link, "/wiki/") && !ignoreLink(link) && !isin(link, links) && !visited[link] && !strings.ContainsAny(link, "#") {
//                 // Append the link to the slice
//                 links = append(links, link)
//                 if link == targeturl {
//                     // If the link matches the target URL, set the flag
//                     targetFound = true
//                 }
//             }
//         })
//     })

//     writeFile("links.txt", links)
//     return links, targetFound
// }

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

        // Check if the target URL is reached
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
        links,found := getListofLinks(targetURL,currentURL,visited)
		if found {
			return append(path, targetURL)
			//fmt.Println("we")
		}
        // Add new paths to the queue
        for _, link := range links {
            newPath := append([]string{}, path...)
            newPath = append(newPath, link)
            queue = append(queue, newPath)
        }
    }
    return nil // Target URL not found
}
func BFS_con(startURL, targetURL string) []string {
	visited := make(map[string]bool)
	queue := [][]string{{startURL}}
	//result := make(chan []string)

	for len(queue) > 0 {
		select {
		case <-time.After(time.Second * 10):
			fmt.Println("Timeout reached")
			return nil
		default:
			path := queue[0]
			queue = queue[1:]

			// Get the last visited URL in the path
			currentURL := path[len(path)-1]

			// Check if the target URL is reached
			if currentURL == targetURL {
				return path
			}

			// Skip if the URL is already visited
			if visited[currentURL] {
				continue
			}
			visited[currentURL] = true

			// Use a wait group to synchronize goroutines
			var wg sync.WaitGroup

			// Get links from the current URL
			links, found := getListofLinks(targetURL, currentURL, visited)

			if found {
				return append(path,targetURL)
			}

			// Add new paths to the queue
			for _, link := range links {
				wg.Add(1)
				go func(link string) {
					defer wg.Done()
					newPath := append([]string{}, path...)
					newPath = append(newPath, link)
					queue = append(queue, newPath)
				}(link)
			}

			// Wait for all goroutines to finish before processing the next level
			wg.Wait()
		}
	}

	return nil // Target URL not found
}

func IDS(startURL, targetURL string, depthLimit int) []string {
    var wg sync.WaitGroup
    resultChan := make(chan []string)

    for depth := 0; depth <= depthLimit; depth++ {
        wg.Add(1)
        go func(depth int) {
            defer wg.Done()
            visited := make(map[string]bool)
            path := DLS(startURL, targetURL, depth, visited)
            if path != nil {
                resultChan <- path
            }
        }(depth)
    }

    go func() {
        wg.Wait()
        close(resultChan)
    }()

    for path := range resultChan {
        if path != nil {
            return path
        }
    }

    fmt.Println("not found")
    return nil
}
func IDS2(startURL, targetURL string, depthLimit int) []string {
    for depth := 0; depth <= depthLimit; depth++ {
		visited := make(map[string]bool)
		fmt.Println("depth" , depth)
        path := DLS(startURL, targetURL, depth,visited)
        if path != nil {
            return path
        }
    }
	fmt.Println("not found")
    return nil
}

func IDS3(startURL, targetURL string, depthLimit int) []string {
    i := 0
    found := false
    while (i < depthLimit && !found){
        path := DLS(startURL,targtargetURL,i,visited)
        i++
    }
    for depth := 0; depth <= depthLimit; depth++ {
		visited := make(map[string]bool)
		fmt.Println("depth" , depth)
        path := DLS(startURL, targetURL, depth,visited)
        if path != nil {
            return path
        }
    }
	fmt.Println("not found")
    return nil
}

// Function to perform depth-limited search
func DLS(currentURL, targetURL string, depthLimit int,visited (map[string]bool)) []string {
    if depthLimit == 0 {
        return nil
    }

    if currentURL == targetURL {
        return []string{currentURL}
    }

    if visited[currentURL] {
        return nil
    }
    visited[currentURL] = true

    links,found := getListofLinks(targetURL,currentURL,visited)
	//var i = 0
	if found {
		return []string{targetURL}
		//fmt.Println("we")
	}else if (!found && depthLimit-1 == 0){
        return nil
    }
    
    for _, link := range links {
		//fmt.Println("dls loop",i)
		//i++
        path := DLS(link, targetURL, depthLimit-1,visited)
        if path != nil {
            return append([]string{currentURL}, path...)
        }
    }

    return nil
}

type linkExtractionResult struct {
    Links       []string
    TargetFound bool
}

// Helper function to extract links from a part of the document
func extractLinks(doc *goquery.Document, visited map[string]bool, targeturl string, links *[]string, start, end int) bool {
    var targetFound bool
    doc.Selection.Slice(start, end).Find("#mw-content-text").Each(func(i int, content *goquery.Selection) {
        // Extract links within the main content area
        content.Find("a").Each(func(i int, s *goquery.Selection) {
            // Get the link's href attribute
            link, exists := s.Attr("href")

            if exists && strings.HasPrefix(link, "/wiki/") && !ignoreLink(link) && !isin(link, *links) && !visited[link] && !strings.ContainsAny(link, "#") {
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

func getListofLinks(targeturl, url string, visited map[string]bool) ([]string, bool) {
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
    resultChan := make(chan linkExtractionResult, 4) // Two parts: top and bottom

    partSize := doc.Length() / 4

    // Process each part of the document concurrently
    for i := 0; i < 4; i++ {
        start := i * partSize
        end := (i + 1) * partSize
        if i == 3 {
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

    writeFile("links.txt", allLinks)
    return allLinks, targetFound
}

