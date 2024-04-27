package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)


//an Implementation of BFS using prioqueue
func BFSWithPrioqueue(startURL, targetURL string, counter *int) []string {

	// Caching Webs to find
	webFind := make(map[string]bool)

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
	num_of_guorotine := 10

	// Signal to stop all goroutines
	stopchan := make(chan struct{})

	pathQueue[startURL] = []string{startURL}

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
			webFind[currentURL] = true
			mutex.Unlock()

			mutex.Lock()
			links, isFound := getListofLinksMult(targetURL, currentURL, webFind, httpClient)
			mutex.Unlock()

			mutex.Lock()
			if isFound {
				select{
				default:
					newPath := append(path, targetURL)
					result = newPath
					mutex.Unlock()
					Checker()
					close(stopchan)
					return
				case <- stopchan:
					mutex.Unlock()
					return
				}
				
			}

			mutex.Unlock()
			for _, link := range links {
				mutex.Lock()
				if !webFind[link] {
					newPath := append([]string{}, path...)
					newPath = append(newPath, link)
					pathQueue[link] = newPath
					webFind[link] = true
					urlToFind.Enqueue(link, depth + 1)
				}
				mutex.Unlock()
			}
		}
	}
	clients := make([]*http.Client, num_of_guorotine)

	for i := 0; i < num_of_guorotine; i++ {
		clients[i] = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	wg.Add(num_of_guorotine)
	for i := 0; i < num_of_guorotine;i++{
		go worker(i,clients[i])
	}
	wg.Wait()
	return result
}