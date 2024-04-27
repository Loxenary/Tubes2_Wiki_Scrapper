package main

import (
	"fmt"
	"sync"
	"time"
)



// IDS
func IDS(startURL, targetURL string, depthLimit int, counter *int) []string {

	var mutex sync.Mutex
	var path []string // Variable penyimpanan path yang dicari

	if startURL == targetURL { // Apabila startURL == targetURL maka selesai
		return []string{startURL}
	} else if depthLimit <= 1 { // Error Handling , Jika depth == 1 berarti mengecek apakah start == target karena sudah dicek maka return nil)
		return nil
	}
	for depth := 2; depth <= depthLimit; depth++ { //mulai loop dari depth == 2 karena depth 1 sudah dicek dan depth < 1 -> nil
 
		// Membuat SafeMap baru pada setiap loop
		visited := *NewSafeMap()


		fmt.Println("depth", depth)

		// Start Timer untuk dls melihat waktu yang diperlukan setiap depth
		start := time.Now()
<<<<<<< HEAD

		path = DLS(startURL, targetURL, depth,&visited, &mutex, counter)
		fmt.Println("DLS Time", depth, ":", time.Since(start)) // Output runtime depth ke-depth
		fmt.Println(path) 
		if path != nil { //Apabila path != []
=======
		var wg sync.WaitGroup
		path = DLS(startURL, targetURL, depth,&visited, &mutex, counter, &wg)
		fmt.Println("DLS Time", depth, ":", time.Since(start))
		wg.Wait();
		
		fmt.Println(path)
		if path != nil {
>>>>>>> 0583096ff45f1c89931eb5b94d4c403194db6935
			return path
		}
		fmt.Println("Not Found")
	}

	fmt.Println("not found")
	return nil
}

<<<<<<< HEAD
// DLS concurrent
func DLS(currentURL, targetURL string, depthLimit int, visited *SafeMap, mutex *sync.Mutex, counter *int) []string {
=======
// Performing Depth limited search concurrently
func DLS(currentURL, targetURL string, depthLimit int, visited *SafeMap, mutex *sync.Mutex, counter *int, wg *sync.WaitGroup) []string {
>>>>>>> 0583096ff45f1c89931eb5b94d4c403194db6935

	if visited.Get(currentURL) { // Mengecek apakah sudah dikunjungi (nil jika sudah)
		fmt.Println("flag visited")
		return nil
	}

	if currentURL == targetURL {
		return []string{currentURL}
	}

<<<<<<< HEAD
	visited.Set(currentURL, true) // Mengset dikunjungi -> true
=======
	fmt.Println("Current URL: ",currentURL, " depth: ",depthLimit);

	visited.Set(currentURL, true)
>>>>>>> 0583096ff45f1c89931eb5b94d4c403194db6935
	mutex.Lock()
	(*counter)++ // Menambah jumlah article yang dikunjungi
	mutex.Unlock()

<<<<<<< HEAD
	links, found := getListofLinks1(targetURL, currentURL, *visited) // Mengambil URL dari currentURL
=======
	links, found := LinksProcessor(targetURL, currentURL)
>>>>>>> 0583096ff45f1c89931eb5b94d4c403194db6935

	if found {
		fmt.Println("found target from", currentURL)
		return []string{currentURL, targetURL} 
	} else if !found && depthLimit == 2 {  // Karena selanjutnya depth == depthLimit -1 == 1 , dan dalam links tidak ada target maka nil
		return nil
	}
<<<<<<< HEAD

	// variable wait group untuk DLS selanjutnya
	var wg sync.WaitGroup 

	for _, link := range links { //loop dari URL yang didapat, dimulai dari yang paling awal 
=======
	for _, link := range links {
>>>>>>> 0583096ff45f1c89931eb5b94d4c403194db6935

		wg.Add(1) 
		worker := func(link string) []string { //Menambah go routine
			defer wg.Done()
<<<<<<< HEAD
			subPath := DLS(link, targetURL, depthLimit-1, visited, mutex, counter) //Recursive
			if subPath != nil {
				return append([]string{currentURL}, subPath...) // Mengembalikan [currentURL,hasil DLS...]
=======
			subPath := DLS(link, targetURL, depthLimit-1, visited, mutex, counter, wg)
			if subPath != nil {
				return append([]string{currentURL}, subPath...)
>>>>>>> 0583096ff45f1c89931eb5b94d4c403194db6935
			}
			return nil //jika hasil DLS []
		}(link)

		if worker != nil {
			fmt.Println("Worker :", worker)
			return worker
		}
	}

	return nil
}