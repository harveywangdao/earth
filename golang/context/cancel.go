package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("ctx done")
		}
	}()

	fmt.Println("sleep 2")
	time.Sleep(2 * time.Second)
	cancel()

	fmt.Println("wait ctx done")
	wg.Wait()
	fmt.Println("app end")
}
