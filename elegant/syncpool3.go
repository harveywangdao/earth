package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]int, 10)
		fmt.Println("make 10")
		return b
	},
}

func main() {
	arr := bytePool.Get().([]int)

	for i := 0; i < 10; i++ {
		arr[i] = i
	}

	fmt.Println(arr)
	bytePool.Put(arr)

	arr2 := bytePool.Get().([]int)
	fmt.Println(arr2)
	bytePool.Put(arr2)

	runtime.GC()

	arr3 := bytePool.Get().([]int)
	fmt.Println(arr3)
	bytePool.Put(arr3)

	time.Sleep(2 * time.Second)
}
