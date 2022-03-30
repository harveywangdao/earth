package main

import (
	"fmt"
	"runtime"
)

func printmem() {
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	fmt.Println("mem:", m.Alloc)
}

func f1(arr [1024 * 1024 * 10]int) {
	printmem()
	arr[0] = 222
	printmem()
}

func do1() {
	printmem()
	arr := [1024 * 1024 * 10]int{}
	printmem()
	arr[0] = 2
	f1(arr)
	fmt.Println(arr[0])
	printmem()
}

func do2() {
	arr := [4]int{1, 2, 3, 4}
	for k, v := range arr {
		if k == 0 {
			arr[0] = 10
			arr[1] = 11
			arr[2] = 12
			arr[3] = 13
		}
		fmt.Println(v)
	}
	fmt.Println(arr)
}

func do3() {
	printmem()
	arr := [1024 * 1024 * 100]int{}
	printmem()
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			printmem()
		}
		arr[i]++
	}
}

func do4() {
	printmem()
	arr := [1024 * 1024 * 50]int{}
	printmem()
	for k, v := range arr {
		if k == 0 {
			printmem()
		}
		arr[k] = v + 1
	}
	arr[0] = 1
}

func main() {
	do1()
}
