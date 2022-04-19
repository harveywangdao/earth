package main

import (
	"sync"
)

func do1() {
	mu := new(sync.Mutex)
	n := 0
	mu.Lock()
	n++
	mu.Unlock()
}

func do2() {
	mu := new(sync.Mutex)
	mu.Lock()
	mu.Lock()
}

func main() {
	do2()
}
