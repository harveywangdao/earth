package main

import (
	"sync"
)

/*
GOMAXPROCS=4 GODEBUG=schedtrace=1000,scheddetail=1 go run debug.go
*/
func do1() {
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var counter int
			for i := 0; i < 1e10; i++ {
				counter++
			}
		}()
	}
	wg.Wait()
}

func main() {
	do1()
}
