package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	var rwlock sync.RWMutex

	n := 0

	go func() {
		for {
			time.Sleep(3 * time.Second)

			fmt.Println("Lock-start")
			rwlock.Lock()
			fmt.Println("Lock-end")

			n++
			time.Sleep(7 * time.Second)

			fmt.Println("Unlock-start")
			rwlock.Unlock()
			fmt.Println("Unlock-end")

			break
		}
	}()

	go func() {
		for {
			time.Sleep(1 * time.Second)

			fmt.Println("Read1 RLock-start")
			rwlock.RLock()
			fmt.Println("Read1 RLock-end")

			fmt.Println("Read1--", n)

			time.Sleep(7 * time.Second)

			fmt.Println("Read1 RUnlock-start")
			rwlock.RUnlock()
			fmt.Println("Read1 RUnlock-end")

			break
		}
	}()

	count := 2
	for i := count; i < count+100; i++ {
		go func(i int) {
			for {
				time.Sleep(2 * time.Second)

				fmt.Println("Read" + strconv.Itoa(i) + " RLock-start")
				rwlock.RLock()
				fmt.Println("Read" + strconv.Itoa(i) + " RLock-end")

				fmt.Println("Read"+strconv.Itoa(i)+"--", n)

				time.Sleep(7 * time.Second)

				fmt.Println("Read" + strconv.Itoa(i) + " RUnlock-start")
				rwlock.RUnlock()
				fmt.Println("Read" + strconv.Itoa(i) + " RUnlock-end")

				break
			}
		}(i)
	}

	count += 10
	go func() {
		time.Sleep(4 * time.Second)

		for i := count; i < count+2; i++ {
			fmt.Println("Read" + strconv.Itoa(i) + " RLock-start")
			rwlock.RLock()
			fmt.Println("Read" + strconv.Itoa(i) + " RLock-end")

			fmt.Println("Read"+strconv.Itoa(i)+"--", n)

			//time.Sleep(5 * time.Second)

			fmt.Println("Read" + strconv.Itoa(i) + " RUnlock-start")
			rwlock.RUnlock()
			fmt.Println("Read" + strconv.Itoa(i) + " RUnlock-end")
		}
	}()

	var c chan bool
	c <- false
}
