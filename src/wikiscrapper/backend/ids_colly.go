package main

import (
	//"bufio"
	"fmt"
	"regexp"
	"strings"
	"sync"

	//"os"
	//"strings"
	//"sync"
	"time"

	"github.com/gocolly/colly"
)

// var end bool = false
// var count int
// var visited SafeMap
//var stopchan = make(chan struct{})

func IDS_col(startURL, targetURL string, depthLimit int, counter *int) []string {
    

    
	var path_col []string
    url := "https://en.wikipedia.org" + startURL
    //var wg sync.WaitGroup

    done := make(chan struct{})
	//sem := make(chan struct{}, 5)

    // Map to store visited URLs for each depth level
    visitedURLs := make(map[string]bool)

	for depth := 1; depth <= depthLimit; depth++ {
		var wg sync.WaitGroup
        // Initialize visited URLs map for the current depth level
        visitedURLs = make(map[string]bool)

        c := colly.NewCollector(
			// Set appropriate options like allowed domains, user agent, etc.
			colly.AllowedDomains("en.wikipedia.org"),
	
			colly.MaxDepth(depth),
			colly.DisallowedURLFilters(
				regexp.MustCompile(`.*/wiki/Main_Page`),
			),
			colly.URLFilters(
				regexp.MustCompile(`/wiki/.*`),
			),
			
			colly.AllowURLRevisit(),
			colly.Async(true),
		)
		
		// Define behavior for visited pages
		c.Limit(&colly.LimitRule{
			Parallelism: 5,
			Delay:       1 * time.Millisecond,
		})
		// Set error handler
		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Error fetching URL:", r.Request.URL, "-", err)
		})
		
        fmt.Println("depth :", depth)
		fmt.Println("maxdepth :",c.MaxDepth)

        *counter = 0
        data := make(map[int]string)
        data[0] = startURL
        count := 1
        c.OnHTML("#mw-content-text", func(e *colly.HTMLElement) {
            e.ForEach("a[href]",func(_ int,el *colly.HTMLElement){
				// Extract links and process them
				link := el.Attr("href")
				//link := e.Request.URL.Path
				//fmt.Println("info link : ")
	
				if strings.HasPrefix(link, "/wiki/") && !ignoreLink(link) && !visitedURLs[link] {
					*counter++
					
					data[count] = link
					// Check if the link has been visited at the current depth level
					visitedURLs[link] = true // Mark the link as visited at the current depth
					if link == targetURL {
						select {
						case <-done:
							// If already closed, don't do anything
							el.Request.Abort()
							return
						default:
							close(done) // Close the done channel to signal completion
							for idx := 0; idx < depth; idx++ {
								path_col = append(path_col, data[count-idx])
							}
							path_col = append(path_col, data[0])
						}
						return
	
					} else {
						//sem <- struct{}{}
						c.MaxDepth--
						if (c.MaxDepth >= 1){
							count++
							fmt.Println("max depth :",c.MaxDepth)
							//DLS_col(link,depth-1,c)
							c.Visit("https://en.wikipedia.org"+link)
						}
						c.MaxDepth++
					}
					
				}
			})
            
        })
		
	
		// Add number of expected Go routines to wait for
		wg.Add(1)
        c.Visit(url)
        c.Wait()
		wg.Wait()
        select {
        case <-done:
            return path_col
        default:
        }

        // Clear visited URLs map for the next depth level iteration
        //delete(visitedURLs,true)
    }
    return nil
    
}

func DLS_col(currentURL string,depthLimit int, c *colly.Collector){
	fmt.Println(currentURL,"depth:",depthLimit)
	fmt.Println("max depth in dls:",c.MaxDepth)
	if depthLimit == 0 {
		return
	}
	c.Visit("https://en.wikipedia.org"+currentURL)
	
}
