package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx1, cancel1 := context.WithTimeout(context.Background(), 10*time.Second)
	ctx2, cancel2 := context.WithTimeout(ctx1, 5*time.Second)

	var wg sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan int)

	fmt.Println("start", time.Now())

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx1.Done():
			fmt.Println("ctx1 done", time.Now())
		case <-ch1:
			fmt.Println("ch1 done", time.Now())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx2.Done():
			fmt.Println("ctx2 done", time.Now())
		case <-ch2:
			fmt.Println("ch2 done", time.Now())
		}
	}()

	time.AfterFunc(2*time.Second, func() {
		cancel2()
	})

	wg.Wait()
	cancel1()
}
