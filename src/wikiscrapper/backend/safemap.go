package main

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
}