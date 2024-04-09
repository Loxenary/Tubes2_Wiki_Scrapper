package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Ini Struct buat nyimpen data input dari web
type Data struct{
	FROM string `json:"FROM"`
	TO string `json:"TO"`
	Algorithm string `json:"Algorithm"`
}

//Ini struct buat nyimpen data path yang dilalui
type Path struct{
	Index string `json:"index"`
	Item string `json:"item"`
}

//Ini struct buat nyimpen hasil output yang bakal di balikin ke web
type Response struct {
    Checkcount string `json:"checkcount"`
    NumPassed string `json:"numpassed"`
    Time   string `json:"time"`
    ListPath []Path `json:"listPath"`
}

//Ini dummyData buat path
var dummyPathList = []Path{
	{Index: "1", Item: "item1"},
	{Index: "2", Item: "item2"},
	{Index: "3", Item: "item3"},
	{Index: "4", Item: "item4"},
}

//Ini DummyData contoh hasil resultnya kaya gimana nanti
var dummyResult = Response{
	Checkcount: "1", 
	NumPassed: "2", 
	Time: "3.22", 
	ListPath: dummyPathList,
}

//Fungsi utama buat POST api dari web
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

	//variable yang nyimpen data From ama To
	//TODO: implement function yang implement bfs ama dsi
	dataAlgorithm := fmt.Sprintf(data.Algorithm)
	print("Algorithm : "+ dataAlgorithm + "\n")
	dataFrom := fmt.Sprintf(data.FROM)
	print("From: "+dataFrom +"\n")
	dataTo := fmt.Sprintf(data.TO)
	print("To: "+dataTo + "\n")
	
	//variable yang  nyimpen data hasil pencarian
	response := dummyResult

	//Setup json decoder
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}



func main(){
	http.HandleFunc("/api/postData", postDataHandler)
    http.ListenAndServe(":8080", nil)
}