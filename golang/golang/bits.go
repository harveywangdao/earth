package main

import (
	"fmt"
	"math"
)

func do1() {
	var maxInt64 int64 = 0x7FFFFFFFFFFFFFFF
	fmt.Println("math.MaxInt64:", math.MaxInt64, "0x7FFFFFFFFFFFFFFF:", maxInt64)

	var minInt64 int64 = maxInt64 + 1
	fmt.Println("math.MinInt64", math.MinInt64, "minInt64:", minInt64)

	var maxUint64 uint64 = 0xFFFFFFFFFFFFFFFF
	fmt.Println("math.MaxUint64:", uint64(math.MaxUint64), "0xFFFFFFFFFFFFFFFF:", maxUint64)
}

func do2() {
	var n1 int64 = -1
	n2 := uint64(n1)
	fmt.Printf("n2 = 0x%X, n2 = %d\n", n2, n2)

	var n3 uint64 = 0xFFFFFFFFFFFFFFFF
	n4 := int64(n3)
	fmt.Printf("n4 = %d\n", n4)
}

func main() {
	do1()
	do2()
}
