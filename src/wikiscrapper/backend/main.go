package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
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


//Global Variable
var UrlData = make(chan Data, 1) //Global variable for storing post data
var OutputData = make(chan Response, 1) //Global variable for storing an algorithmn response

// API Post Request Handler
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
// API Get Request Handler
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


// Convert from []string into []Path
func PathConverter(str []string) [] Path{
    var paths []Path
    for i :=0; i < len(str); i++ {
        paths = append(paths, Path{str[i]})
    }
    return paths
}

// Process Data From Frontend to Backend and return The result back
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
<<<<<<< HEAD
        counter := 0
=======
        counter := int(0)

        // Used to Inform the Frontend about current state of the program
        
>>>>>>> 7a1cff5a99f5f8765452f9c09d5c93e6a1a2cd0f
        if data.Algorithm == "BFS" {
            path = BFSWithPrioqueue(url, target, &counter)
        } else {
            //visited := make(map[string]bool)
<<<<<<< HEAD
            //path = IDS(url, target, 6)
            path = IDS_col(url,target,6,&counter)
=======
            path = IDS(url, target, 6, &counter)
>>>>>>> 7a1cff5a99f5f8765452f9c09d5c93e6a1a2cd0f
        }
        runtime := time.Since(start)
        if (path == nil){
            fmt.Println("empty path")
        }else{
            fmt.Println(path)
        }
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
        if (response.ListPath == nil){
            fmt.Println("empty list")
        }
        for i := 0; i < len(response.ListPath); i++ {
            if (len(response.ListPath) != 0){
                fmt.Println(response.ListPath[i].Item)
            }else{
                fmt.Println("flag empty list")
            }
        }
        clearFile("links.txt")
        clearFile("output.txt")
        OutputData <- response
    }
}

//Main Function
func main() {
    
    go processData()
    http.HandleFunc("/api/postData", postDataHandler)
    http.HandleFunc("/api/getData", getDataHandler)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}


<<<<<<< HEAD
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

    if strings.ContainsAny(link,"%"){
        return true;
    }
    if strings.ContainsAny(link,"#"){
        return true;
    }
    for _, prefix := range ignoreList {
		if strings.HasPrefix(link, prefix) {
			return true
		}
	}
	return false
}


=======
//Write links into file named filename. Used as debug
>>>>>>> 7a1cff5a99f5f8765452f9c09d5c93e6a1a2cd0f
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

//Used to clear all the data inside the filename
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

// Used as Checker
func Checker(){
    fmt.Println("is this called??")
}
  

