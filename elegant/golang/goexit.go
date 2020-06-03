package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func task(r *rand.Rand) {
	defer fmt.Println("Hello defer")
	for {
		fmt.Println("Hello")
		aa := r.Int()
		if aa%19 == 0 {
			fmt.Println("Hello Exit1", aa)
			runtime.Goexit()
			fmt.Println("Hello Exit2")
		}
	}
}

func main() {
	defer func() {
		fmt.Println("game over")
	}()

	r := rand.New(rand.NewSource(time.Now().Unix()))
	go task(r)

	time.Sleep(2 * time.Second)
}
