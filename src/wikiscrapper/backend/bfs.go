package main

import (
	"fmt"
	"net/http"
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
	webFind := make(map[string]bool)
	var urlToFind Prioqueue
	urlToFind.Init(targetURL)
	urlToFind.Enqueue(startURL, 0)
	var pathQueue = make(map[string][]string)

	var result []string = nil
	// var linksgroup sync.WaitGroup
	var mutex sync.Mutex
	// var isResultFound = false
	var wg sync.WaitGroup
	num_of_guorotine := 2

	stopChan := make(chan struct{})

	pathQueue[startURL] = []string{startURL}

	worker := func(workerID int, httpClient *http.Client) {
		defer wg.Done()
		for {
			select {
			case <-stopChan:
				return
			default:
				mutex.Lock()
				if urlToFind.Length() == 0 {
					fmt.Println("Length is 0", "Worker: ", workerID)
					mutex.Unlock()
					continue
				}
				mutex.Unlock()

				currentURL, priority, depth := urlToFind.Dequeue()

				if depth == 99 {
					continue
				}

				mutex.Lock()
				path := pathQueue[currentURL]
				mutex.Unlock()

				fmt.Println("Worker:", workerID, "URL TO FIND:", currentURL, "Priority:", priority, "Depth:", depth, "Length : ", urlToFind.Length())

				mutex.Lock()
				webFind[currentURL] = true
				mutex.Unlock()

				//Process 2 dia bakal nunggu channel url
				mutex.Lock()
				links, isFound := getListofLinksMult(targetURL, currentURL, webFind, httpClient)
				mutex.Unlock()


				if isFound {
					mutex.Lock()
					*counter += len(links)
					newPath := append(path, targetURL)
					result = newPath
					mutex.Unlock()
					close(stopChan)
					return
				}
				appendedItem := []Item{}
				for _, link := range links {
					mutex.Lock()
					if !webFind[link] {
						mutex.Unlock()
						newPath := append([]string{}, path...)
						newPath = append(newPath, link)

						mutex.Lock()
						pathQueue[link] = newPath
						webFind[link] = true
						*counter++

						item := Item{
							key:      link,
							depth:    depth + 1,
							priority: urlToFind.priorityDecision(link),
						}
						appendedItem = append(appendedItem, item)
						mutex.Unlock()
						urlToFind.Enqueue(link, depth+1)
					}
				}

				mutex.Lock()
				writeFilePrioque("links.txt", appendedItem, currentURL)
				mutex.Unlock()
				//Inform sebuah 
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

// func BFSWithColly(startURL, targetURL string, counter *int) []string {
// 	c := colly.NewCollector(
// 		colly.AllowedDomains("en.wikipedia.org"),
// 		colly.Async(true),
// 		colly.CacheDir("links.txt"),
// 	)

// 	// Limit the crawler's speed to avoid hitting Wikipedia servers too hard
// 	c.Limit(&colly.LimitRule{
// 		DomainGlob:  "*.wikipedia.org*",
// 		RandomDelay: 1 * time.Second,
// 	})

// 	var path []string
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex

// 	queue := []string{startURL}
// 	visited := make(map[string]bool)
// 	webfind := make(map[string]bool)
// 	parent := make(map[string]string)

// 	for len(queue) > 0 {
// 		url := queue[0]
// 		queue = queue[1:]

// 		wg.Add(1)
// 		go func(url string) {
// 			defer wg.Done()

// 			c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 				link := e.Attr("href")
// 				if strings.HasPrefix(link, "/wiki/") {
// 					if()
// 					mu.Lock()
// 					*counter++
// 					mu.Unlock()

// 					childURL := e.Request.AbsoluteURL(link)

// 					mu.Lock()
// 					parent[childURL] = url
// 					mu.Unlock()

// 					if !visited[childURL] {
// 						visited[childURL] = true
// 						queue = append(queue, childURL)
// 					}

// 					mu.Lock()
// 					path = append(path, childURL)
// 					mu.Unlock()

// 					if childURL == targetURL {
// 						fmt.Println("Target URL found!")
// 						return
// 					}
// 				}
// 			})

// 			c.Visit(url)
// 		}(url)
// 	}

// 	wg.Wait()

// 	return getPath(startURL, targetURL, parent)
// }
// func getPath(startURL, targetURL string, parent map[string]string) []string {
// 	path := make([]string, 0)
// 	curr := targetURL

// 	for curr != startURL {
// 		path = append([]string{curr}, path...)
// 		p, ok := parent[curr]
// 		if !ok || p == "" {
// 			return nil
// 		}
// 		curr = p
// 	}

// 	path = append([]string{startURL}, path...)
// 	return path
// }