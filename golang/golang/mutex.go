package main

import (
	"sync"
)

func do1() {
	var mu sync.Mutex
	n := 0
	mu.Lock()
	n++
	mu.Unlock()
}

func main() {
	do1()
}
