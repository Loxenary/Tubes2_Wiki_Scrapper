package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
var OutputData = make(chan Response, 1) //Global variable for storing the output response

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

    //Signal data untuk diprocess di ProcessData method
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
    

        start := time.Now()
        var path []string
        counter := int(0)

        // Used to Inform the Frontend about current state of the program
        
        if data.Algorithm == "BFS" {
            path = BFSWithPrioqueue(url, target, &counter)
        } else {
            //visited := make(map[string]bool)
            path = IDS(url, target, 6, &counter)
        }
        runtime := time.Since(start)
        if (path == nil){
            fmt.Println("empty path")
        }else{
            fmt.Println(path)
        }

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

