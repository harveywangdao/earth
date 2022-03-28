package main

import (
	"fmt"
)

func f1(arr [100]int) {
	arr[0] = 222
}

func do1() {
	arr := [100]int{}
	arr[0] = 2
	f1(arr)
	fmt.Println(arr[0])
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
}

func main() {
	do2()
}
