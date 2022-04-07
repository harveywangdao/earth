package main

import (
	"fmt"
	"sync"
	"time"
)

func do1() {
	now := time.Now()
	next := now.Add(0)
	timer := time.NewTimer(next.Sub(now))

	fmt.Println("1")
	<-timer.C
	fmt.Println("2")
}

func do2() {
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), 17, 14, 0, 0, now.Location())
	if now.After(next) {
		next = next.Add(time.Hour * 24)
	}

	timer := time.NewTimer(next.Sub(now))
	<-timer.C
}

//优雅地关闭timer
func do3() {
	timer := time.NewTimer(1 * time.Second)
	stopC := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case t := <-timer.C:
				fmt.Println("time.C:", t)
				timer.Reset(1 * time.Second)
			case <-stopC:
				fmt.Println("stopC")
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)
	timer.Stop()
	stopC <- true
	fmt.Println("stop timer")

	wg.Wait()
}

// timer结束后还能接收timer.C吗? 可以
func do4() {
	timer := time.NewTimer(1 * time.Second)
	stopC := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case t := <-timer.C:
				fmt.Println("time.C:", t)
			case <-stopC:
				fmt.Println("stopC")
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)
	timer.Stop()
	stopC <- true
	fmt.Println("stop timer")

	wg.Wait()
}

func main() {
	do4()
}
