package main

import (
	"sync"
	"unicode/utf8"
)

type Target struct{
	name string
}

type Item struct {
    key      string
    priority int
}

type PriorityQueue struct {
    items []*Item
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
    pq.pq.items = make([]*Item, 0)
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

func (pq *Prioqueue) Enqueue(key string) {
    pq.Lock()
    defer pq.Unlock()
    if pq.pq.items == nil {
        pq.pq.items = make([]*Item, 0)
    }

    removeWiki := key[5:]
    priority := pq.priorityDecision(removeWiki)
    item := &Item{key, priority}
    if(pq.treshold > 0){
        pq.treshold--
        pq.pq.items = append(pq.pq.items, item)
        pq.bubbleUp(len(pq.pq.items) - 1)
    }else{
        if(item.priority < 15){
            pq.pq.items = append(pq.pq.items, item)
            pq.bubbleUp(len(pq.pq.items) - 1)
        }
    }
    
}

func (pq *Prioqueue) bubbleUp(index int) {
    for index > 0 {
        parentIndex := (index - 1) / 2
        if pq.pq.items[index].priority < pq.pq.items[parentIndex].priority {
            pq.pq.items[index], pq.pq.items[parentIndex] = pq.pq.items[parentIndex], pq.pq.items[index]
            index = parentIndex
        } else {
            break
        }
    }
}

func (pq *Prioqueue) Dequeue() (string,int) {
    pq.Lock()
    defer pq.Unlock()
    if len(pq.pq.items) == 0 {
        return "",0
    }
    root := pq.pq.items[0]
    lastIndex := len(pq.pq.items) - 1
    pq.pq.items[0] = pq.pq.items[lastIndex]
    pq.pq.items = pq.pq.items[:lastIndex]
    pq.heapifyDown(0)
    return root.key,root.priority
}

func (pq *Prioqueue) heapifyDown(index int) {
    lastIndex := len(pq.pq.items) - 1
    for {
        leftChildIndex := 2*index + 1
        rightChildIndex := 2*index + 2
        swapIndex := index

        if leftChildIndex <= lastIndex && pq.pq.items[leftChildIndex].priority < pq.pq.items[swapIndex].priority {
            swapIndex = leftChildIndex
        }
        if rightChildIndex <= lastIndex && pq.pq.items[rightChildIndex].priority < pq.pq.items[swapIndex].priority {
            swapIndex = rightChildIndex
        }

        if swapIndex != index {
            pq.pq.items[index], pq.pq.items[swapIndex] = pq.pq.items[swapIndex], pq.pq.items[index]
            index = swapIndex
        } else {
            break
        }
    }
}
