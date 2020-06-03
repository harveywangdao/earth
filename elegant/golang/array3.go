package main

import (
	"fmt"
)

//var arr = []int{6, 4, -3, 5, -2, -1, 0, 1, -9}
//var arr = []int{-6, -4, -3, -5, 2, 1, 1, 9, 7}
var arr = []int{6, 4, 0, -3, 5, -2, -1, 0, 1, 0, -9, 0}

func main() {
	p := 0
	n := len(arr) - 1
	z := 0

	for {
		if p+z == n {
			break
		}

		if arr[p] >= 0 {
			p++
		} else if arr[p] < 0 {
			temp := arr[n]
			arr[n] = arr[p]
			arr[p] = temp
			n--
		}
	}

	fmt.Println(arr)
}
