package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	runtime.Gosched()

	go func() {
		i := int64(0)
		for {
			i++
			if i%50000000000 == 0 {
				fmt.Println("go---1", i)
			}
		}
	}()

	go func() {
		i := int64(0)
		for {
			i++
			fmt.Println("go---2", i)
		}
	}()

	fmt.Println(runtime.NumGoroutine(), runtime.NumCPU())
	var c chan int
	c <- 1
}
