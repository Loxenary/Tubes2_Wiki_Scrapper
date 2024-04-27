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
	visited := *NewDepthSafeMap()
	for depth := 2; depth <= depthLimit; depth++ { //mulai loop dari depth == 2 karena depth 1 sudah dicek dan depth < 1 -> nil
 
		// Membuat SafeMap baru pada setiap loop
		


		fmt.Println("depth", depth)

		// Start Timer untuk dls melihat waktu yang diperlukan setiap depth
		start := time.Now()
		var wg sync.WaitGroup
		path = DLS(startURL, targetURL, depth,&visited, &mutex, counter, &wg)
		fmt.Println("DLS Time", depth, ":", time.Since(start))
		wg.Wait();
		
		fmt.Println(path)
		if path != nil {
			return path
		}
		fmt.Println("Not Found")
	}

	fmt.Println("not found")
	return nil
}

// Performing Depth limited search concurrently
func DLS(currentURL, targetURL string, depthLimit int, visited *DepthSafeMap, mutex *sync.Mutex, counter *int, wg *sync.WaitGroup) []string {

	if visited.Get(depthLimit,currentURL) { // Mengecek apakah sudah dikunjungi (nil jika sudah)
		fmt.Println("flag visited")
		return nil
	}

	if currentURL == targetURL {
		return []string{currentURL}
	}

	fmt.Println("Current URL: ",currentURL, " depth: ",depthLimit);

	visited.Set(depthLimit,currentURL, true)
	mutex.Lock()
	(*counter)++ // Menambah jumlah article yang dikunjungi
	mutex.Unlock()

	links, found := LinksProcessor(targetURL, currentURL)

	if found {
		fmt.Println("found target from", currentURL)
		return []string{currentURL, targetURL} 
	} else if !found && depthLimit == 2 {  // Karena selanjutnya depth == depthLimit -1 == 1 , dan dalam links tidak ada target maka nil
		return nil
	}
	for _, link := range links {

		wg.Add(1) 
		worker := func(link string) []string { //Menambah go routine
			defer wg.Done()
			subPath := DLS(link, targetURL, depthLimit-1, visited, mutex, counter, wg)
			if subPath != nil {
				return append([]string{currentURL}, subPath...)
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