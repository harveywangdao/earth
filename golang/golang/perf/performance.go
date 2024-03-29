package main

import (
	"context"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

//go:noinline
func ReadArray1() {
	arr := [1024 * 1024 * 10]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
}

//go:noinline
func ReadArray2() {
	arr := [1024 * 1024 * 10]int{}
	for k, v := range arr {
		arr[k] = k + v
	}
}

func do1() {
	go func() {
		for {
			ReadArray1()
		}
	}()
	go func() {
		for {
			ReadArray2()
		}
	}()
	select {}
}

func do2() {
	n := 1
	var lc sync.Mutex
	go func() {
		for {
			lc.Lock()
			n++
			lc.Unlock()
		}
	}()

	go func() {
		for {
			lc.Lock()
			n++
			lc.Unlock()
		}
	}()

	select {}
}

func do3(ctx context.Context) {

}

func main() {
	//fmt.Printf("%s", 12)
	go func() {
		http.ListenAndServe(":6060", nil)
	}()
	go do1()
	do2()
	do3(context.Background())
}
