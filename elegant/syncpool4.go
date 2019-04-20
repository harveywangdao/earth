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

func sub() {
	var count int = 100
	newfun := func() interface{} {
		count++
		return count
	}

	pool := sync.Pool{New: newfun}

	v1 := pool.Get()
	fmt.Printf("v1 :%v\n", v1)

	pool.Put(9)
	pool.Put(10)
	pool.Put(11)
	pool.Put(12)

	v2 := pool.Get()
	fmt.Printf("v2 :%v\n", v2)

	runtime.GC()

	v3 := pool.Get()
	fmt.Printf("v3 :%v\n", v3)

	pool.New = nil

	v4 := pool.Get()
	fmt.Printf("v4 :%v\n", v4)
}

func main() {
	arr := make([]int, 10)

	for i := 0; i < 10; i++ {
		arr[i] = i
	}

	fmt.Println(arr)
	bytePool.Put(arr)

	runtime.GC()
	arr[7] = 45
	fmt.Println(arr)

	arr2 := bytePool.Get().([]int)
	arr2[3] = 47
	fmt.Println(arr2)
	bytePool.Put(arr2)

	runtime.GC()

	arr3 := bytePool.Get().([]int)
	fmt.Println(arr3)
	bytePool.Put(arr3)

	sub()
	time.Sleep(2 * time.Second)
}
