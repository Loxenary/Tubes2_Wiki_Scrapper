package main

import (
	"strings"
)

// A Link Used as a channel 
type linkExtractionResult struct {
	Links       []string
	TargetFound bool
}



func ignoreLink(link string) bool{
    ignoreList := []string{
        "/wiki/File:" ,
        "/wiki/Help:" ,
        "/wiki/Special:" ,
        "/wiki/Template:" ,
        "/wiki/Template_Talk:" ,
        "/wiki/Template_talk:" ,
        "/wiki/Wikipedia:" ,
        "/wiki/Category:",
        "/wiki/Portal:" ,
        "/wiki/User:" ,
        "/wiki/User_Talk:" ,
        "/wiki/Talk:",
    }

    if(strings.Contains(link,"%")){
        return true
    }

    for _, prefix := range ignoreList {
		if strings.HasPrefix(link, prefix){
			return true
		}
	}
	return false
}

func isin(link string, array []string) bool {
    // Calculate the size of each part
    partSize := len(array) / 4

    // Channel to receive results
    found := make(chan bool, 4)

    // Process each part of the array concurrently
    for i := 0; i < 4; i++ {
        start := i * partSize
        end := (i + 1) * partSize
        if i == 3 {
            end = len(array) // Ensure the last part includes the remaining elements
        }
        go func(arr []string) {
            for _, item := range arr {
                if item == link {
                    found <- true
                    return
                }
            }
            found <- false
        }(array[start:end])
    }

    // Collect results from the channels
    for i := 0; i < 4; i++ {
        if <-found {
            return true
        }
    }

    return false
}
