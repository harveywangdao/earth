package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func cond() {
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)

	cond.Signal()
	cond.Broadcast()

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			fmt.Println(i, "cond locking")
			cond.L.Lock()

			fmt.Println(i, "cond waiting")
			cond.Wait()
			fmt.Println(i, "cond wait done")

			cond.L.Unlock()
			fmt.Println(i, "cond unlock")
		}(i)
	}

	time.Sleep(2 * time.Second)
	cond.Signal()
	time.Sleep(2 * time.Second)
	cond.Broadcast()

	wg.Wait()
}

func multicore() {
	var wg sync.WaitGroup

	fmt.Println(runtime.NumCPU())

	//runtime.GOMAXPROCS(4)
	//runtime.GOMAXPROCS(1)

	fmt.Println(time.Now().Unix())

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				fmt.Println(i)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println(time.Now().Unix())
}

func rwlock() {
	//1.加了写锁，不能再加写锁  ok
	//2.加了写锁，不能再加读锁  ok
	//3.加了读锁，不能再加写锁
	//4.加了读锁，能再加读锁    ok

	var rw sync.RWMutex
	var wg sync.WaitGroup

	//fmt.Println(runtime.NumCPU())
	//runtime.GOMAXPROCS(4)

	num := 0

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				rw.Lock()
				fmt.Println("W", i, "Lock")

				time.Sleep(1 * time.Second) //1.2.

				num++

				if num >= 5 {
					fmt.Println("W", i, "Unlock")
					rw.Unlock()
					break
				}

				fmt.Println("W", i, "Unlock")
				rw.Unlock()
			}
		}(i)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				rw.RLock()
				fmt.Println("R", i, "RLock")

				time.Sleep(1 * time.Second) //3.4.

				fmt.Println("R", i, "num =", num)
				if num >= 5 {
					fmt.Println("R", i, "RUnlock over")
					rw.RUnlock()
					break
				}
				fmt.Println("R", i, "RUnlock")
				rw.RUnlock()
			}
		}(i)
	}

	wg.Wait()
}

func syncmap() {
	var sm sync.Map

	sm.Store(1, "aaa")

	v, ok := sm.Load(1)
	fmt.Println("v =", v, "ok =", ok)

	v, ok = sm.Load(2)
	fmt.Println("v =", v, "ok =", ok)

	v, ok = sm.LoadOrStore(1, "AAA")
	fmt.Println("v =", v, "ok =", ok)

	v, ok = sm.LoadOrStore(2, "bbb")
	fmt.Println("v =", v, "ok =", ok)

	sm.Store(3, "ccc")
	sm.Store(4, "ddd")
	sm.Store(5, "eee")

	sm.Range(func(key, value interface{}) bool {
		fmt.Println("key =", key, "value =", value)
		return true
	})

	sm.Delete(1)

	fmt.Println("after delete:")

	sm.Range(func(key, value interface{}) bool {
		fmt.Println("key =", key, "value =", value)
		return true
	})
}

func main() {
	cond()
}
