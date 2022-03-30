package main

import (
	"fmt"
	"runtime"
)

func printmem(tag string) {
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	fmt.Println(tag, "mem:", m.Alloc/1024/1024, "MB")
}

func do1() {
	printmem("1")
	arr := [1024 * 1024 * 10]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	printmem("2")

	arr2 := arr
	printmem("3")
	arr[0] = 100
	arr2[0] = 100
	printmem("4")
}

func main() {
	do1()
}
