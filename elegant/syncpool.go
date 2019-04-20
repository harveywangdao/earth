package main

import (
	"fmt"
	"sync"
	"time"
)

var bytePool = sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		fmt.Println("make 1024")
		return &b
	},
}

func main() {
	//defer
	//debug.SetGCPercent(debug.SetGCPercent(-1))
	a := time.Now().UnixNano()
	fmt.Println(a)

	for i := 0; i < 10000; i++ {
		obj := make([]byte, 1024)
		//fmt.Println("without pool")
		_ = obj
	}

	b := time.Now().UnixNano()
	fmt.Println(b)

	for j := 0; j < 10000; j++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		//fmt.Println("with pool")
		bytePool.Put(obj)
	}

	c := time.Now().UnixNano()
	fmt.Println(c)

	fmt.Println("without pool", b-a, "ns")
	fmt.Println("with pool", c-b, "ns")
}
