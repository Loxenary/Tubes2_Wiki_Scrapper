package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
	"github.com/PuerkitoBio/goquery"
)

//Ini Struct buat nyimpen data input dari web
type Data struct{
	FROM string `json:"FROM"`
	TO string `json:"TO"`
	Algorithm string `json:"Algorithm"`
}

//Ini struct buat nyimpen data path yang dilalui
type Path struct{
	Item string `json:"item"`
}

//Ini struct buat nyimpen hasil output yang bakal di balikin ke web
type Response struct {
    Checkcount string `json:"checkcount"`
    NumPassed string `json:"numpassed"`
    Time   string `json:"time"`
    ListPath []Path `json:"listPath"`
}

var UrlData = make(chan Data, 1)
var OutputData = make(chan Response, 1)
func postDataHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Error(w,  "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//variable yang nyimpen data json awal
	var data Data
	//error handling + save data ke variable data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
    UrlData <- data
    fmt.Println("Data From: " + data.FROM)
    fmt.Println("Data To: " + data.TO)
    fmt.Println("Data algorithm: " + data.Algorithm)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode([]byte{})
}

func getDataHandler(w http.ResponseWriter, r *http.Request){
    if(r.Method != "GET"){
        http.Error(w,  "Method Not Allowed", http.StatusMethodNotAllowed)
		return
    }
    select {
    case response := <-OutputData:
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    default:
        w.WriteHeader(http.StatusNoContent)
    }
}

func PathConverter(str []string) [] Path{
    var paths []Path
    for i := len(str) - 1; i >= 0; i-- {
        paths = append(paths, Path{str[i]})
    }
    return paths
}

func processData(){
    for{

        data :=<- UrlData
        url := data.FROM
        target := data.TO
    
        // Process data...
        clearFile("links.txt")
        clearFile("output.txt")
        start := time.Now()
        var path []string
        if data.Algorithm == "BFS" {
            path = BFS(url, target)
        } else {
            visited := make(map[string]bool)
            path = DLS(url, target, 4, visited)
        }
        runtime := time.Since(start)
    
        // Construct response
        response := Response{
            Checkcount: "25",
            NumPassed:  fmt.Sprint(len(path)),
            Time:       fmt.Sprint(runtime),
            ListPath:   PathConverter(path),
        }
        
        fmt.Println()
        fmt.Println("Data Checkoutn: " + response.Checkcount)
        fmt.Println("Data Numpassed: " + response.NumPassed)
        fmt.Println("Data Time: " + response.Time)
        fmt.Println("Data ListPath: ")
        for i := 0; i < len(response.ListPath); i++ {
            fmt.Println(response.ListPath[i].Item)
        }
        OutputData <- response
    }
}
func main() {
    
    go processData()
    http.HandleFunc("/api/postData", postDataHandler)
    http.HandleFunc("/api/getData", getDataHandler)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
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

    // Extract links
    var links []string
    targetFound := false // Flag to indicate if the target URL has been found

    doc.Find("#mw-content-text").Each(func(i int, content *goquery.Selection) {
        // Extract links within the main content area
        content.Find("a").Each(func(i int, s *goquery.Selection) {
            // Get the link's href attribute
            link, exists := s.Attr("href")
            if exists && strings.HasPrefix(link, "/wiki/") && !ignoreLink(link) && !isin(link, links) && !visited[link] && !strings.ContainsAny(link, "#") {
                // Append the link to the slice
                links = append(links, link)
                if link == targeturl {
                    // If the link matches the target URL, set the flag
                    targetFound = true
                }
            }
        })
    })
    return links, targetFound
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

func BFSTest(startURL, targetURL string) []string {
    visited := make(map[string]bool)
    queue := [][]string{{startURL}}
    var wg sync.WaitGroup
    wg.Add(1)
    var resultPath []string
    for len(queue) > 0 {
        path := queue[0]
        queue = queue[1:]
        currentURL := path[len(path)-1]

        // Check if the target URL is reached
        if currentURL == targetURL {
            wg.Wait() // Wait for all goroutines to finish
            return path
        }
        if visited[currentURL] {
            continue
        }
        visited[currentURL] = true

        // Get links from the current URL concurrently
        go func(currentURL string) {
            defer wg.Done()
            links, found := getListofLinks(targetURL, currentURL, visited)
            if found {
                resultPath = path
            }
            var appendedData []string
            for _, link := range links {
                if !visited[link] {
                    newPath := append([]string{}, path...)
                    newPath = append(newPath, link)
                    queue = append(queue, newPath)
                    appendedData = append(appendedData, link)
                }
            }
            writeFile("links.txt", appendedData)
        }(currentURL)
    }
    wg.Wait() // Wait for all goroutines to finish
    return resultPath // Target article not found
}

func BFS(startURL, targetURL string) []string {
    visited := make(map[string]bool)
    queue := [][]string{{startURL}}
    for len(queue) > 0 {
        path := queue[0]
        queue = queue[1:]
        currentURL := path[len(path)-1]

        // Check if the target URL is reached
        if currentURL == targetURL {
            return path
        }
        if visited[currentURL] {
            continue
        }
        visited[currentURL] = true

        // Get links from the current URL
        //println("flag")
        links,found := getListofLinks(targetURL,currentURL,visited)
		if found {
            reversedPath := make([]string, len(path))
            for i := range path {
                reversedPath[i] = path[len(path)-1-i]
            }
			return append([]string{targetURL}, reversedPath...)
		}
        // Add new paths to the queue
        for _, link := range links {
            if !visited[link] {
                newPath := append([]string{}, path...)
                newPath = append(newPath, link)
                queue = append(queue, newPath)
            }
        }
    }
    return nil // Target article not found
}
  

func IDS(startURL, targetURL string, depthLimit int) []string {
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
    path := []string{currentURL}
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
		return append([]string{targetURL},path...)
		//fmt.Println("we")
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


