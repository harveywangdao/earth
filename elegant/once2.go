package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func onces() {
	fmt.Println("onces")
}

func onced() {
	fmt.Println("onced")
}

func main() {
	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("count:", v, "---", i)
	}

	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onced)
			fmt.Println("213")
		}()
	}

	time.Sleep(time.Second * 2)
}
