package main

import (
	"bytes"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants"
)

func getGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func do1() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b := make([]byte, 22)
			runtime.Stack(b, false)
			fmt.Println(string(b))
		}()
	}
	wg.Wait()

	runtime.GC()
	time.Sleep(1 * time.Second)
	fmt.Println("\n")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b := make([]byte, 22)
			runtime.Stack(b, false)
			fmt.Println(string(b))
		}()
	}
	wg.Wait()

	runtime.GC()
	time.Sleep(1 * time.Second)
	fmt.Println("\n")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			log.Println(getGoroutineID())
		}()
	}
	wg.Wait()
}

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func do2() {
	defer ants.Release()
	runTimes := 1000

	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}

	wg.Wait()
	log.Printf("running goroutines: %d\n", ants.Running())

	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()

	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	log.Printf("running goroutines: %d\n", p.Running())
	log.Printf("finish all tasks, result is %d\n", sum)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do2()
}
