package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var int32num int32 = 9
	fmt.Println("int32num =", int32num)

	atomic.AddInt32(&int32num, 1)
	fmt.Println("AddInt32 int32num =", int32num)

	swap := atomic.CompareAndSwapInt32(&int32num, 10, 11)
	fmt.Println("CompareAndSwapInt32 int32num =", int32num, "swap =", swap)

	v := atomic.LoadInt32(&int32num)
	fmt.Println("LoadInt32 v =", v)

	atomic.StoreInt32(&int32num, 12)
	fmt.Println("StoreInt32 int32num =", int32num)

	v = atomic.SwapInt32(&int32num, 13)
	fmt.Println("SwapInt32 int32num =", int32num, "old v =", v)
}
