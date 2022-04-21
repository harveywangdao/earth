package main

import (
	"sync"
	"time"
)

func do1() {
	mu := new(sync.Mutex)
	n := 0
	mu.Lock()
	n++
	mu.Unlock()
}

var (
	n = 0
)

func m1(mu *sync.Mutex) {
	for {
		mu.Lock()
		time.Sleep(time.Second * 2)
		n++
		mu.Unlock()
	}
}

func do2() {
	mu := new(sync.Mutex)
	for i := 0; i < 1000; i++ {
		go m1(mu)
	}
	select {}
}

func main() {
	do2()
}
