package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strings"
	"sync"
	"time"
)

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
		links, found := getListofLinks1(targetURL, currentURL, visited)
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

func IDS(startURL, targetURL string, depthLimit int) []string {

	end = false
	for depth := 0; depth <= depthLimit; depth++{
		//visited := make(map[string]bool)
		fmt.Println("depth",depth)
		start := time.Now()
		//path := DLS(startURL, targetURL, depth,visited)
		path := DLS1(startURL, targetURL, depth)
		fmt.Println("waktu DLS",depth,":",time.Since(start))
		if path != nil {
			writeFile("output.txt",append([]string{"Output IDS :"},path...))
			return path
		}
		
	}

	fmt.Println("not found")
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

var end bool = false

func DLS1(currentURL, targetURL string, depthLimit int) []string {
	//fmt.Println("DLS",currentURL,"Depth :",depthLimit)
	if depthLimit == 0 {
		return nil
	}

	if currentURL == targetURL {
		return []string{currentURL}
	}

	if depthLimit == 1{
		return nil
	}
	links, found := getListofLinks2(targetURL, currentURL)
	if found {
		return []string{currentURL,targetURL}
	}else if !found && depthLimit  == 2{
		return nil
	}

	var wg sync.WaitGroup
	resultChan := make(chan []string)
	sem := make(chan struct{}, 20)

	for _, link := range links {
		if end{
			fmt.Println("end")
			break
		}
		sem <- struct{}{}
		wg.Add(1)
		go func(link string) {
			defer func() {
                <-sem // release semaphore
                wg.Done()
            }()
			subPath := DLS1(link, targetURL, depthLimit-1)
			if subPath != nil {
				writeFile("output.txt",append([]string{currentURL}, subPath...))
				resultChan <- append([]string{currentURL}, subPath...)
				end = true
			}
		}(link)
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
