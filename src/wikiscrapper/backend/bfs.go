package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)
func BFSWithoutPrioqueue(startURL, targetURL string, counter *int) []string{
	//Stash the web that has been visited
	visited := make(map[string]bool)
	//Stash the web that has been checked from another web
	webfind := make(map[string]bool)

	//Stash The url that just added
	var urlToFind = []string{startURL}

	//Stash the Path
	var pathQueue = [][]string{{startURL}}

	var result []string = nil

	var wg sync.WaitGroup
	var mutex sync.Mutex
	num_of_guorotine := 50

	stopchan := make(chan struct{})

	worker := func(workerID int, httpClient *http.Client){
		// Reduce Guorotine after shutdown
		defer wg.Done()
		for{
			mutex.Lock()
			//Early Terminate for 0 length
			if len(urlToFind) == 0{
				fmt.Println("Length is 0 ", "Worker : ", workerID)
				mutex.Unlock()
				return
			}
			(*counter)++
			//Get The Un-Checkable URL from urlToFind
			currentURL := urlToFind[0]
			urlToFind = urlToFind[1:]
			//Check the Last-Checked Path inside the Queue
			path := pathQueue[len(pathQueue)-1]

			fmt.Println("Worker: ", workerID, "Current URL: ", currentURL, "Length: ", len(path))

			if(currentURL == targetURL){
				(*counter)++
				//Signal to Stop Channel to stop other guorotines
				close(stopchan)
				result = path
				mutex.Unlock()
				return 
			}
			if visited[currentURL] {
				mutex.Unlock()
                continue
            }
			mutex.Unlock()
			
			mutex.Lock()
            visited[currentURL] = true
			webfind[currentURL] = true
            
			links, isFound := getListofLinksMult(targetURL,currentURL,webfind,httpClient)
			
			if(isFound){
				(*counter) += len(links)
                newPath := append(path, targetURL)
                result = newPath
				fmt.Println("FEFEFEFE")
				close(stopchan)
				mutex.Unlock()
                return
			}
			mutex.Unlock()
			appendedLink := []string{}
			for _, link := range links {
				mutex.Lock()
				if !webfind[link] {
                    newPath := append([]string{}, path...)
                    newPath = append(newPath, link)
                    pathQueue = append(pathQueue, newPath)
					urlToFind = append(urlToFind, link)
					appendedLink = append(appendedLink, link)
					webfind[link] = true
					(*counter)++
				}
				mutex.Unlock()
			}
			mutex.Lock()
			writeFile("links.txt", appendedLink)
			mutex.Unlock()
			select{
				default:
				case <- stopchan:
					return
			}
		}
	}

	Clients := make([]*http.Client,num_of_guorotine)
	wg.Add(num_of_guorotine)
	for i:=0 ; i < num_of_guorotine; i++{
		Clients[i] = &http.Client{
            Timeout: 10 * time.Second,
        }
		go worker(i+1,Clients[i])
	}
	wg.Wait()
	return result
}



func BFSWithChannel(startURL,targetURL string, counter *int) []string{
	stopChan := make(chan struct{})
	pathChan := make(chan []string)
	webFind := make(map[string]bool)
	fetched := make(map[string]bool)
	paths := make(map[string][]string)

	var httpClient = &http.Client{}

	queue := []string{startURL}

	var wg sync.WaitGroup

	for{
		if(len(queue) == 0){
			continue
		}
		currUrl := queue[0]
		queue = queue[1:]
		
		if(fetched[currUrl]){
			continue
		}

		wg.Add(1)
		go func(url string){
			defer wg.Done()
			defer close(stopChan)
			links, isFound := getListofLinksMult(targetURL,url,webFind,httpClient)
			if(isFound){
				(*counter) += len(links)
                pathChan <- append(paths[url],targetURL)
                return
			}

			fetched[url] = true

			for _, link := range links {
				if(!webFind[link]){
					queue = append(queue, link)
                    paths[link] = append(paths[url], link)
				}
            }
		}(currUrl)
		Checker()
		select{
		case <- stopChan:
			Checker()
			return <-pathChan
		default :
			Checker()
		// Continue the loop until <- stopchan is closed
		}
	}
}

func BFSWithPrioqueue(startURL, targetURL string, counter *int) []string {
	visited := make(map[string]bool)
	webFind := make(map[string]bool)
	var urlToFind Prioqueue
	urlToFind.Init(targetURL)
	urlToFind.Enqueue(startURL, 0)
	var pathQueue = make(map[string][]string)

	var result []string = nil
	// var linksgroup sync.WaitGroup
	var mutex sync.Mutex
	var webtex sync.Mutex
	// var isResultFound = false
	var wg sync.WaitGroup
	num_of_guorotine := 2

	stopchan := make(chan struct{})

	pathQueue[startURL] = []string{startURL}

	worker := func(workerID int, httpClient *http.Client) {
		defer wg.Done()
		for {
			urlToFind.Log("length")
			if urlToFind.Length() == 0 {
				if(len(visited) != 0){
					fmt.Println("Length is 0 ","Worker: ", workerID)
					return
				}
			}
			currentURL, priority, depth := urlToFind.Dequeue()

			if(depth == 99){
				continue
			}

			mutex.Lock()
			path := pathQueue[currentURL]
			mutex.Unlock()

			fmt.Println("Worker:", workerID, "URL TO FIND:", currentURL, "Priority:", priority, "Depth:", depth, "Length : ", urlToFind.Length())
			
			if currentURL == targetURL {
				result = path
				// isResultFound = true
				return
			}
			mutex.Lock()
			if visited[currentURL] {
				mutex.Unlock()
				continue
			}

			visited[currentURL] = true
			webFind[currentURL] = true
			mutex.Unlock()

			webtex.Lock()
			links, isFound := getListofLinksMult(targetURL, currentURL, webFind, httpClient)
			if isFound {
				(*counter) += len(links)
				newPath := append(path, targetURL)
				result = newPath
				webtex.Unlock()
				close(stopchan)
				return
			}
			
			webtex.Unlock()
			
			appendedLink := []string{}
			appendedDepth := []string{}
			for _, link := range links {
				
				mutex.Lock()
				if !webFind[link] {
					newDepth := strconv.Itoa(depth + 1)
					newPath := append([]string{}, path...)
					newPath = append(newPath, link)
					pathQueue[link] = newPath
					webFind[link] = true
					(*counter)++

					appendedLink = append(appendedLink, link)
					appendedDepth = append(appendedDepth, newDepth)
					urlToFind.Enqueue(link, depth + 1)
				}
				mutex.Unlock()
			}

			mutex.Lock()
			writeFilePrioque("links.txt", appendedLink, appendedDepth)
			mutex.Unlock()
			select{
				default:
				case <- stopchan:
					return
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