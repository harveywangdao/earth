package main

import (
	"fmt"
	"sync/atomic"
)

func do1() {
	var n int32 = 1
	swapped := atomic.CompareAndSwapInt32(&n, 2, 3)
	fmt.Println(n, swapped)
	swapped = atomic.CompareAndSwapInt32(&n, 1, 3)
	fmt.Println(n, swapped)
}

func main() {
	do1()
}
