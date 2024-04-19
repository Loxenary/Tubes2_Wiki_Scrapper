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

var UrlData = make(chan Data)
var OutputData = make(chan Response)

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
	//Setup json decoder
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Ok")
}

func getDataHandler(w http.ResponseWriter, r *http.Request){
    if(r.Method != "GET"){
        http.Error(w,  "Method Not Allowed", http.StatusMethodNotAllowed)
		return
    }
    select {
	case response := <-OutputData:
		// Data available, encode and send the response
		w.Header().Set("Content-Type", "application/json")		
		json.NewEncoder(w).Encode(response)
	default:
		// No data available, send an empty response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{})
	}
}

func PathConverter(str []string) [] Path{
    var paths []Path
    for _, s := range str{
        path := Path{Item: s}
        paths = append(paths, path)
    }
    return paths
}


func main() {
    http.HandleFunc("/api/postData", postDataHandler)
    http.HandleFunc("/api/getData", getDataHandler)
    go func() {
        if err := http.ListenAndServe(":8080", nil); err != nil {
            log.Fatal(err)
        }
    }()
    for{
        select{
        case data:= <- UrlData:
                fmt.Printf("TSETETETETE");
                url := data.FROM;
                target := data.TO;
                clearFile("links.txt")
                clearFile("output.txt")
                start := time.Now()
                var path []string
                if(data.Algorithm == "BFS"){
                    path = BFS(url,target)
                }else{
                    visited:= make(map[string]bool)
                    path = DLS(url,target,4,visited)
                }
                runtime := time.Since(start)
                response := Response{
                    Checkcount: "25",
                    NumPassed: fmt.Sprint(len(path)),
                    Time:   fmt.Sprint(runtime),
                    ListPath: PathConverter(path),
                }
                writeFile("output.txt",path)
                OutputData <- response;
        }
    }
}

func getListofLinksNew(url string, visited map[string]bool, target string) [] string{
    url = "https://en.wikipedia.org" + url

    response, err := http.Get(url)
    if err != nil {
        log.Fatal("Error fetching URL:", err)
    }
    defer response.Body.Close()

    doc, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error parsing HTML:", err)
    }

    var links []string
    doc.Find("#mw-content-text").Each(func(i int, content *goquery.Selection) {
        // Extract links within the main content area
        content.Find("a").Each(func(i int, s *goquery.Selection) {
            // Get the link's href attribute
            link, exists := s.Attr("href")
            
            if exists && strings.HasPrefix(link, "/wiki/") && !ignoreLink(link) && !isin(link,links) && !visited[link] && !strings.ContainsAny(link,"#") {
                // Append the link to the slice
                links = append(links, link)
            }
        })
        
    })
    writeFile("links.txt", links)
    return links;
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
    doc.Find("#mw-content-text").Each(func(i int, content *goquery.Selection) {
        // Extract links within the main content area
        content.Find("a").Each(func(i int, s *goquery.Selection) {
            // Get the link's href attribute
            link, exists := s.Attr("href")
            
            if exists && strings.HasPrefix(link, "/wiki/") && !ignoreLink(link) && !isin(link,links) && !visited[link] && !strings.ContainsAny(link,"#") {
                // Append the link to the slice
                links = append(links, link)
            }
        })
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
    var wg sync.WaitGroup
    for len(queue) > 0 {
        path := queue[0]
        queue = queue[1:]
        currentURL := path[len(path)-1]
        if currentURL == targetURL {
            return path
        }
        if visited[currentURL] {
            continue
        }
        visited[currentURL] = true
        links := make(chan []string)
        go func(url string) {
            defer wg.Done()
            links <- getListofLinks(url, visited)
        }(currentURL)
        select {
        case <-time.After(time.Second * 5):
            // Timeout handling
            fmt.Println("Timeout occurred for", currentURL)
            continue
        case newLinks := <-links:
            // Add new paths to the queue
            for _, link := range newLinks {
                if !visited[link] {
                    newPath := append([]string{}, path...)
                    newPath = append(newPath, link)
                    queue = append(queue, newPath)
                }
            }
        }
        wg.Add(1)
    }
    wg.Wait()
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


