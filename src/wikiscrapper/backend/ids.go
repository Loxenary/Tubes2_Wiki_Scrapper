package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strings"
	"sync"
	"time"
)


type SafeMap struct {
    mu sync.Mutex
    visited map[string]bool
}



func NewSafeMap() *SafeMap {
    return &SafeMap{visited: make(map[string]bool)}
}

func (sm *SafeMap) Set(key string, value bool) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    sm.visited[key] = value
}

func (sm *SafeMap) Get(key string) bool {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    return sm.visited[key]
}



var end bool = false
var count int
var visited SafeMap

func IDS(startURL, targetURL string, depthLimit int) []string {
	
	end = false
	for depth := 0; depth <= depthLimit; depth++{
		visited = *NewSafeMap()
		count = 0
		if end {break}
		//visited := make(map[string]bool)
		fmt.Println("depth",depth)
		start := time.Now()
		//path := DLS(startURL, targetURL, depth,visited)
		path := DLS1(startURL, targetURL, depth,visited)
		fmt.Println("waktu DLS",depth,":",time.Since(start))
		fmt.Println(path)
		if path != nil {
			fmt.Println("Ketemu")
			fmt.Println("melalui",count,"artikel")
			return path
		}
		fmt.Println("Tidak Ketemu")
	}

	fmt.Println("not found")
	return nil
}

func DLS1(currentURL, targetURL string, depthLimit int, visited SafeMap) []string {
    // Make a copy of the visited map
    localVisited := SafeMap{visited: make(map[string]bool)}
    for key, value := range visited.visited {
        localVisited.visited[key] = value
    }

    //fmt.Println("DLS", currentURL, "Depth :", depthLimit)

    if localVisited.Get(currentURL) {
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
    localVisited.Set(currentURL, true)
    count++
    links, found := getListofLinks1(targetURL, currentURL, localVisited)
    if found {
		fmt.Println("found target from",currentURL)
        return []string{currentURL, targetURL}
    } else if !found && depthLimit == 2 {
        return nil
    }

    var wg sync.WaitGroup
    resultChan := make(chan []string)
    sem := make(chan struct{}, 30)

    for _, link := range links {
        if end {
            fmt.Println("end")
            break
        }
        sem <- struct{}{}
        wg.Add(1)
        go func(link string) {
            defer func() {
                <-sem // release semaphore
				//fmt.Println("sem left:",len(sem))
                wg.Done()
            }()
            subPath := DLS1(link, targetURL, depthLimit-1, localVisited)
            if subPath != nil {
                writeFile("output.txt", append([]string{"output subpath :",currentURL}, subPath...))
				fmt.Println("subPath from ",currentURL,":",subPath)
				fmt.Println("end :",end)
				end = true
                resultChan <- append([]string{currentURL}, subPath...)
				
                
            }
        }(link)
    }

    go func() {
        wg.Wait()
		//fmt.Println("closing resultchan")
        close(resultChan)
    }()
    select {
    case path := <-resultChan:
        if path != nil {
			//fmt.Println("flag return")
            return path
        }
    default:
		fmt.Println("no path recieved from",currentURL)
        // If no path received, continue
    }

    return nil
}



func IDScon(startURL, targetURL string, depthLimit int) []string {
	var wg sync.WaitGroup
	resultChan := make(chan []string)

	for depth := 0; depth <= depthLimit; depth++ {
		visited := make(map[string]bool)
		wg.Add(1)
		go func(depth int) {
			defer wg.Done()
			path := DLS(startURL, targetURL, depth,visited)
			if path != nil {
				resultChan <- path
			}
		}(depth)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for path := range resultChan {
		if path != nil {
			return path
		}
	}

	fmt.Println("not found")
	return nil
}

func IDS1(startURL, targetURL string, depthLimit int) []string {
	var wg sync.WaitGroup
	var resultMutex sync.Mutex
	resultByDepth := make(map[int][]string)
	depthsWithResult := make(map[int]bool)

	for depth := 0; depth <= depthLimit; depth++ {
		fmt.Println("Depth", depth)
		wg.Add(1)
		go func(depth int) {
			defer wg.Done()
			visited := make(map[string]bool)
			path := DLS(startURL, targetURL, depth, visited)
			resultMutex.Lock()
			if path != nil {
				resultByDepth[depth] = path
				depthsWithResult[depth] = true
			}
			resultMutex.Unlock()
		}(depth)
	}

	wg.Wait()

	// Find the deepest depth with a result
	deepestResultDepth := -1
	for depth := depthLimit; depth >= 0; depth-- {
		if depthsWithResult[depth] {
			deepestResultDepth = depth
			break
		}
	}

	if deepestResultDepth == -1 {
		fmt.Println("not found")
		return nil
	}

	// Return the path of the deepest depth with a result
	return resultByDepth[deepestResultDepth]
}

// Function to perform depth-limited search
func DLS(currentURL, targetURL string, depthLimit int, visited map[string]bool) []string {
	fmt.Println("DLS",currentURL,"Depth :",depthLimit)
	if depthLimit == 0 {
		return nil
	}

	if currentURL == targetURL {
		return []string{currentURL}
	}

	if visited[currentURL] {
		return nil
	}
	visited[currentURL] = true

	if depthLimit > 1{
		links, found := getListofLinks2(targetURL, currentURL)
		if found{
			return []string{currentURL,targetURL}
			//fmt.Println("we")
		}else if !found && depthLimit == 2{
			return nil
		}
		for _, link := range links {
			//fmt.Println("dls loop",i)
			//i++
			path := DLS(link, targetURL, depthLimit-1, visited)
			if path != nil {
				return append([]string{currentURL}, path...)
			}
		}
	}
	return nil
}

// func ids_path(startURL,targetURL string,depthLimit int)[]string{
// 	paths := [][]string{{}}
// 	path := []string{startURL}
	

// }

// func dls_path(targetURL string,path []string,depthLimit int)[]string{
// 	currentURL := path[len(path)-1]
// 	if depthLimit == 0{return nil}
// 	if currentURL == targetURL {return append(path,targetURL)}
// 	if depthLimit == 1{return nil}
// 	links,found := getListofLinks2(targetURL,currentURL)
// 	if found {
// 		return append(path,targetURL)
// 	}else if !found && depthLimit == 2{
// 		return nil
// 	}
// 	for _,link := range links{
// 		subPath := dls_path(targetURL,append(path,link),depthLimit-1)

// 	}
// }

// func save_ids_path(paths [][]string){
// 	for _, path := range paths {
//         writeFile("ids_cache.txt",path)
//     }
// }

// func load_ids_path()([][]string,error){
// 	filename := "ids_cache.txt"
// 	// Open the file
//     file, err := os.Open(filename)
//     if err != nil {
//         return nil, err
//     }
//     defer file.Close()

//     var arrayOfArrays [][]string

//     // Create a scanner to read the file line by line
//     scanner := bufio.NewScanner(file)
//     for scanner.Scan() {
//         line := scanner.Text()
//         // Split the line into an array of strings
//         array := strings.Split(line, ",")
//         arrayOfArrays = append(arrayOfArrays, array)
//     }

//     // Check for any errors during scanning
//     if err := scanner.Err(); err != nil {
//         return nil, err
//     }

//     return arrayOfArrays, nil
// }