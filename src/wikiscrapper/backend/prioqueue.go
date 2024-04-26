package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"unicode/utf8"
)

type Target struct{
	name string
}

type Item struct {
    key      string
    priority int
    depth int
}

type PriorityQueue struct {
    items []Item
}

type Prioqueue struct {
    treshold int // Used to preserve some Enqueue before declining priorities
    target Target
    pq     PriorityQueue
    sync.Mutex
}

func (pq *Prioqueue) Length() int {
    return len(pq.pq.items)
}

func (pq *Prioqueue) ConstructTarget(target string) {
    var t Target
    t.name = target
    pq.target = t
}

func (pq *Prioqueue) Init(target string) {
    pq.Lock()
    defer pq.Unlock()
    pq.pq.items = make([]Item, 0)
    pq.ConstructTarget(target)
    pq.treshold = 10 // Limit of Enqueue without priorities

}


func StringCompare(s1, s2 string) int {
    minlengthTreshold := 32
    if len(s1) == 0 {
        return utf8.RuneCountInString(s2)
    }
    if len(s2) == 0 {
        return utf8.RuneCountInString(s1)
    }
    if s1 == s2 {
        return 0
    }
    if len(s1) > len(s2) {
        s1, s2 = s2, s1
    }

    lS1 := len(s1)
    lS2 := len(s2)
    var x []int
    if lS1+1 > minlengthTreshold {
        x = make([]int, lS1+1)
    } else {
        x = make([]int, minlengthTreshold)
        x = x[:lS1+1]
    }
    for i := 1; i < len(x); i++ {
        x[i] = int(i)
    }
    for i := 1; i <= lS2; i++ {
        prev := int(i)
        for j := 1; j <= lS1; j++ {
            current := x[j-1] // match
            if s2[i-1] != s1[j-1] {
                current = min(min(x[j-1]+1, prev+1), x[j]+1)
            }
            x[j-1] = prev
            prev = current
        }
        x[lS1] = prev
    }
    return int(x[lS1])
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func (pq *Prioqueue) priorityDecision(key string) int {
    return StringCompare(key, pq.target.name)
}

func(pq *Prioqueue) ReSortList(item Item){
    id := 0
    for i:= 0; i < pq.Length();i++{
        temp := pq.pq.items[i]
        if(item.depth > temp.depth){
            id = i + 1
            continue
        }

        if(item.depth == temp.depth){
            if item.priority < temp.priority {
                id = i
                break
            } else {
                id = i + 1 // Move to the next position
                continue
            }
        }else if(item.depth < temp.depth){
            id = i
            break
        }
    }

    if(pq.Length() > 0){
        pq.pq.items = append(pq.pq.items[:id], append([]Item{item}, pq.pq.items[id:]...)...)
    }else{
        pq.pq.items = append([]Item{}, item)
    }
}
func writeFilePrioque(filename string, data []Item, parent string) {
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        log.Fatal("Error opening file:", err)
    }
    defer file.Close()

    for _, link := range data {
        // Write each link to the file
        _, err := file.WriteString("Link: " +link.key + " Depth : "+ strconv.Itoa(link.depth) + " Priority: "+ strconv.Itoa(link.priority)+ " Parent : " + parent +  "\n")
        if err != nil {
            log.Fatal("Error writing to file:", err)
        }
    }

    //fmt.Println("Links appended to", filename)
}
func (pq *Prioqueue) Enqueue(key string, depth int) {
    pq.Lock()
    defer pq.Unlock()
    if pq.pq.items == nil {
        pq.pq.items = make([]Item, 0)
    }

    removeWiki := key[5:]
    priority := pq.priorityDecision(removeWiki)

    if(priority == 99){
        return
    }
    item := Item{key, priority, depth}

    
    if(pq.Length() > 5000 ){
        if(priority < 20){
            pq.ReSortList(item);
        }else{
            return;
        }
    }else {
        pq.ReSortList(item);
    }
}

//Dequeue The Prio List 
//Return The Link, Priority, and Depth of the most prioritized Item
func (pq *Prioqueue) Dequeue() (string,int,int) {
    pq.Lock()
    defer pq.Unlock()
    if len(pq.pq.items) == 0 {
        return "",99,99
    }
    root := pq.pq.items[0]
    pq.pq.items = pq.pq.items[1:]
    return root.key,root.priority,root.depth
}
/*
Display Data of Prioqueue
Status : Full (Display Full Data)
Status : ListOnly (Display Only The List)
Status: Length (Dsipaly Only the Length)
*/
func (pq *Prioqueue) Log(status string){
    
    if(status == "full"){
        fmt.Println("=====PRIOQUEUE DATA====")
        fmt.Println("Data: ")
        fmt.Println(pq.pq.items)
        fmt.Println("Length: ")
        fmt.Println(len(pq.pq.items))
        fmt.Println()

    }else if(status == "ListOnly"){
        fmt.Println("=====PRIOQUEUE DATA====")
        fmt.Println("Data: ")
        fmt.Println(pq.pq.items)
    }else{
        fmt.Println("=====PRIOQUEUE DATA====")
        fmt.Println("Length: ")
        fmt.Println(len(pq.pq.items))
    }
}
