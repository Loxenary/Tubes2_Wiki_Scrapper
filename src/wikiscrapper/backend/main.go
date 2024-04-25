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
        counter := int(0)
        if data.Algorithm == "BFS" {
            //path = BFSTest(url, target, &counter)
            path = BFS(url,target)
        } else {
            //visited := make(map[string]bool)
            path = IDS(url, target, 6)
        }
        runtime := time.Since(start)
        writeFile("output.txt",path)
        // Construct response
        response := Response{
            Checkcount: fmt.Sprint(counter),
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

func isin(link string, array []string) bool {
    // Calculate the size of each part
    partSize := len(array) / 4

    // Channel to receive results
    found := make(chan bool, 4)

    // Process each part of the array concurrently
    for i := 0; i < 4; i++ {
        start := i * partSize
        end := (i + 1) * partSize
        if i == 3 {
            end = len(array) // Ensure the last part includes the remaining elements
        }
        go func(arr []string) {
            for _, item := range arr {
                if item == link {
                    found <- true
                    return
                }
            }
            found <- false
        }(array[start:end])
    }

    // Collect results from the channels
    for i := 0; i < 4; i++ {
        if <-found {
            return true
        }
    }

    return false
}



func FindCorrectPath(currentUrl string,queue [][]string) []string{
    for _, path := range queue {
        if path[len(path)-1] == currentUrl {
            // If the last element of the path matches currentURL, return the path
            return path
        } 
        
    }
    return nil
}

func RemovePathFromQueue(queue [][]string, deleted []string) [][]string {
    // Find deleted index 
    
    for i, path := range queue {
        if(len(path) > 0 && len(deleted) > 0){
            if path[len(path)-1] == deleted[len(deleted)-1] {
                newQueue := [][]string{}
                for j, linkPath := range queue{
                    if(i != j){
                        newQueue = append(newQueue, linkPath)
                    }
                }
                return newQueue
            }
        }
    }
    return nil
}

func BFSTest(startURL, targetURL string, counter *int) []string {
    visited := make(map[string]bool)
    webFind := make(map[string]bool)
    var urlToFind Prioqueue
    urlToFind.Init(targetURL)
    urlToFind.Enqueue(startURL)
    var pathQueue = [][]string{{startURL}}
    for(urlToFind.Length()> 0){    
        currentUrl, priority := urlToFind.Dequeue()
        fmt.Println("URL TO FIND: ", currentUrl, "Priority : ",priority)
        path := FindCorrectPath(currentUrl, pathQueue)
        if(currentUrl == targetURL){
            return path
        }

        if(visited[currentUrl]){
            continue
        }
        visited[currentUrl] = true

        webFind[currentUrl] = true
        links, isFound := getListofLinks2(targetURL,currentUrl/*,webFind*/)
        
        if(isFound){
            (*counter) += len(links)
            path := append(path, targetURL)
            
            return path
        }
        appendedLink := []string{}
        for _, link := range links {
            if(!webFind[link]){
                (*counter)++
                newPath := append([]string{}, path...)
                newPath = append(newPath, link)
                pathQueue = append(pathQueue, newPath)
                urlToFind.Enqueue(link)
                appendedLink = append(appendedLink,link)
                webFind[link] = true
            }
        }
        writeFile("links.txt",appendedLink)
        pathQueue = RemovePathFromQueue(pathQueue, path)
    }
    return nil
}

func Checker(){
    fmt.Println("is this called??")
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
        links,found := getListofLinks2(targetURL,currentURL/*,visited*/)
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
        //fmt.Println(queue)
    }
    return nil // Target article not found
}

func BFScon(startURL, targetURL string) []string {
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
        links,found := getListofLinks2(targetURL,currentURL/*,visited*/)
		if found {
			return append(path,targetURL)
		}
        // Add new paths to the queue
        var wg sync.WaitGroup
        for _, link := range links {
            wg.Add(1)
            go func(link string, path []string) {
                defer wg.Done()
                if !visited[link] {
                    newPath := append([]string{}, path...)
                    newPath = append(newPath, link)
                    queue = append(queue, newPath)
                }
            }(link,path)
        }
        wg.Wait() 
        //fmt.Println(queue)
    }
    return nil // Target article not found
}
  

