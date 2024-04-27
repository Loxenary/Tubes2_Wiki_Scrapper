package main

import (
	"sync"
)

// A Safe Map to used in concurrent operations

type SafeMap struct {
    mu sync.Mutex
    visited map[string]bool
}

// Initiate a new SafeMap
func NewSafeMap() *SafeMap {
    return &SafeMap{visited: make(map[string]bool)}
}

// Set the value of a key in the map
func (sm *SafeMap) Set(key string, value bool) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    sm.visited[key] = value
}

// Get the value of a key in the map
func (sm *SafeMap) Get(key string) bool {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    return sm.visited[key]
}

type DepthSafeMap struct {
    mu sync.Mutex
    visited map[int]map[string]bool
}

// Initiate a new SafeMap
func NewDepthSafeMap() *DepthSafeMap {
    return &DepthSafeMap{visited: make(map[int]map[string]bool)}
}

// Set the value of a key in the map
func (sm *DepthSafeMap) Set(depth int,key string, value bool) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    if sm.visited[depth] == nil {
        sm.visited[depth] = make(map[string]bool)
    }
    sm.visited[depth][key] = value
}

// Get the value of a key in the map
func (sm *DepthSafeMap) Get(depth int,key string) bool {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    temp := sm.visited[depth]
    return temp[key]
}