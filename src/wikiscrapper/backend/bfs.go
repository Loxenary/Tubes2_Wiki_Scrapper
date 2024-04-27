package main

import (
	"fmt"
	"net/http"
	"sync"
)


//an Implementation of BFS using prioqueue
func BFSWithPrioqueue(startURL, targetURL string, counter *int) []string {

	// Caching Webs to find
	webFind := *NewSafeMap()

	// Prioqueue to determine the node to execute first
	var urlToFind Prioqueue
	// Initialize Prioqueue
	urlToFind.Init(targetURL)
	urlToFind.Enqueue(startURL, 0)

	// Map to save the path
	var pathQueue = make(map[string][]string)

	// Save the result
	var result []string = nil
	var mutex sync.Mutex
	
	var wg sync.WaitGroup

	// The amount of goroutines that works
	num_of_guorotine := 100

	// Signal to stop all goroutines
	stopchan := make(chan struct{})

	pathQueue[startURL] = []string{startURL}
	webFind.Set(startURL, true); //Caching the first element into the map

	// Goroutine function
	worker := func(workerID int, httpClient *http.Client) {
		defer wg.Done()
		for {
			select{
			default:
			case <- stopchan:
				return
			}
			currentURL, priority, depth := urlToFind.Dequeue()
			
			// If urlToFind cannot be Dequeued
			if(depth == 99){
				continue
			}
			mutex.Lock()
			path := pathQueue[currentURL]
			mutex.Unlock()

			fmt.Println("Worker:", workerID, "URL TO FIND:", currentURL, "Priority:", priority, "Depth:", depth, "Length : ", urlToFind.Length())

			if currentURL == targetURL {
				result = path
				return
			}
			mutex.Lock()
			(*counter)++
			mutex.Unlock()

			links, isFound := HttpLinksProcessor(targetURL, currentURL,httpClient)


			mutex.Lock()
			if isFound {
				select{
				default:
					newPath := append(path, targetURL)
					result = newPath
					mutex.Unlock()
					close(stopchan)
					return
				case <- stopchan:
					mutex.Unlock()
					return
				}
				
			}

			mutex.Unlock()
			for _, link := range links {
				
				if !webFind.Get(link) {
					newPath := append([]string{}, path...)
					newPath = append(newPath, link)
					mutex.Lock()
					pathQueue[link] = newPath
					mutex.Unlock()

					webFind.Set(link,true)
					urlToFind.Enqueue(link, depth + 1)
				}
				
			}
		}
	}
	clients := make([]*http.Client, num_of_guorotine)

	for i := 0; i < num_of_guorotine; i++ {
		clients[i] = &http.Client{
			CheckRedirect: http.DefaultClient.CheckRedirect,
		}
	}
	wg.Add(num_of_guorotine)
	for i := 0; i < num_of_guorotine;i++{
		go worker(i,clients[i])
	}
	wg.Wait()
	return result
}