package main

<<<<<<< HEAD
import "sync"

type SafeMap struct {
	sync.Mutex
	visited map[string]bool
}

func NewSafeMap() *SafeMap {
	return &SafeMap{visited: make(map[string]bool)}
}

func (sm *SafeMap) Set(key string, value bool) {
	sm.Lock()
	defer sm.Unlock()
	sm.visited[key] = value
}

func (sm *SafeMap) Get(key string) bool {
	sm.Lock()
	defer sm.Unlock()
	return sm.visited[key]
=======
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
>>>>>>> 7a1cff5a99f5f8765452f9c09d5c93e6a1a2cd0f
}