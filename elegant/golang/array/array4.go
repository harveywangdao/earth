package main

import (
	"fmt"
	"runtime"
)

func printmem() {
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	fmt.Println("mem:", m.Alloc/1024/1024, "MB")
}

// range会拷贝一份数组,修改原数组不会影响副本
func do1() {
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

// 大数组在堆上分配
func do2() {
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

// range会拷贝一份数组,修改原数组不会影响副本,实际情况是只有一份内存
func do3() {
	printmem()
	arr := [1024 * 1024 * 50]int{}
	printmem()
	for k, v := range arr {
		if k == 0 {
			printmem()
		}
		arr[k] = v + 1
	}
}

func main() {
	do3()
}
