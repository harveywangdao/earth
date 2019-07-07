package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")

	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("child goroutine")
		}
	}()

	select {}

	fmt.Println("end")
}
