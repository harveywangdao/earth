package main

import (
	"fmt"
	"sync"
	"time"
)

func do1() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("child goroutine")
		}
	}()

	select {}
}

func do2() {
	var wg sync.WaitGroup

	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()

		select {
		case c1 <- <-c2:
			println("c1")
		case <-c3:
			println("c3")
		}
	}()

	println("111")
	time.Sleep(time.Second * 3)
	c2 <- 123
	println("222")
	time.Sleep(time.Second * 3)
	close(c3)
	wg.Wait()
}

func do3(c1, c2 chan int) {
	select {
	case <-c1:
		println("1")
	case <-c2:
		println("2")
	}
}

func talk(msg string, sleep int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}()
	return ch
}

func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- <-input1:
			case ch <- <-input2:
			}
		}
	}()
	return ch
}

func do4() {
	ch := fanIn(talk("A", 10), talk("B", 1000))
	for i := 0; i < 10; i++ {
		fmt.Printf("%q\n", <-ch)
	}
}

func do5() {
	go func() {
		println("111")
		time.Sleep(time.Hour)
		println("222")
	}()
	println("333")
	time.Sleep(time.Hour)
	println("444")
}

func main() {
	do4()
}
