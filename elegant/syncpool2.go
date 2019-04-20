package main

import (
	"fmt"
	"sync"
	"time"
)

var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 10)
		fmt.Println("make 1024")
		return b
	},
}

func main() {
	for i := 0; i < 10; i++ {
		go func(j int) {
			arr := bytePool.Get().([]byte)

			arr[j] = 28
			fmt.Println(arr)

			bytePool.Put(arr)
		}(i)
	}

	time.Sleep(2 * time.Second)
}
