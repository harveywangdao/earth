package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/sasha-s/go-deadlock"
)

func do1() {
	var mu deadlock.Mutex
	mu.Lock()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Println("1111")
		mu.Lock()
		log.Println("2222")

		mu.Unlock()
	}()

	log.Println("aaaa")
	time.Sleep(time.Second * 40)
	mu.Unlock()
	log.Println("bbbb")

	wg.Wait()
}

func do2() {
	buf := make([]int, 10)
	buf2 := make([]int, 10, 20)
	buf3 := make([]int, 0, 20)

	log.Println(buf)
	log.Println(buf2)
	log.Println(buf3)
}

func do3() {
	go func() {
		time.Sleep(10 * time.Second)
	}()

	buf := make([]byte, 1024*16)
	for {
		n := runtime.Stack(buf, true)
		if n < len(buf) {
			buf = buf[:n]
			break
		}
		buf = make([]byte, 2*len(buf))
	}

	fmt.Print(string(buf))
}

func do4() {
	var mu1 deadlock.Mutex
	var mu2 deadlock.Mutex

	go func() {
		log.Println("3333")
		mu2.Lock()
		log.Println("4444")
		time.Sleep(1 * time.Second)
		mu1.Lock()
		log.Println("5555")
	}()

	mu1.Lock()
	log.Println("1111")

	time.Sleep(1 * time.Second)

	mu2.Lock()
	log.Println("2222")
}

func do5() {
	ff := func() {
		s := make([]uintptr, 50)
		s = s[:runtime.Callers(2+4, s)]
		fmt.Println(s)
	}

	ff()
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do4()
}
