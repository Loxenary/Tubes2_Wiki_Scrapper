package main

import (
	"fmt"
	"sync"
	"time"
)


var count int
var visited SafeMap


// Performing Iterative Depth search concurrently
func IDS(startURL, targetURL string, depthLimit int, counter *int) []string {
	
	var mutex sync.Mutex
	var path []string

	for depth := 2; depth <= depthLimit; depth++{
		
		// Create new map cacher
		visited = *NewSafeMap()
		fmt.Println("depth",depth)
		
		// Start Timer
		start := time.Now()
		
		
		path = DLS(startURL, targetURL, depth, &visited, &mutex, counter)
		fmt.Println("DLS Time",depth,":",time.Since(start))
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
func DLS(currentURL, targetURL string, depthLimit int, visited *SafeMap, mutex *sync.Mutex, counter *int) []string {

    if visited.Get(currentURL) {
        fmt.Println("flag visited")
        return nil
    }

    if depthLimit == 0 {
        return nil
    }

    if currentURL == targetURL {
        return []string{currentURL}
    }

    if depthLimit == 1 {
        return nil
    }
    visited.Set(currentURL, true)

	mutex.Lock()
	(*counter)++
	mutex.Unlock()

    links, found := getListofLinks1(targetURL, currentURL, *visited)

    if found {
		fmt.Println("found target from",currentURL)
        return []string{currentURL, targetURL}
    } else if !found && depthLimit == 2 {
        return nil
    }

    var wg sync.WaitGroup

    for _, link := range links {
    
        wg.Add(1)
        worker := func(link string)[]string {
            defer wg.Done()
            subPath := DLS(link, targetURL, depthLimit-1, visited, mutex, counter)
            if subPath != nil {
                writeFile("output.txt", append([]string{"output subpath :",currentURL}, subPath...))
				return append([]string{currentURL}, subPath...)
            }
			return nil
        }(link)

		if(worker != nil){
			fmt.Println("TESTESTEST")  
			return worker
		}
    }
	wg.Wait()
		
    return nil
}