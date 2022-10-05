package main

import (
	"fmt"
)

func sq(arr []int, start, end int) {
	if start >= end {
		return
	}

	base := arr[start]
	i := start + 1
	j := end
	for i < j {
		if arr[i] >= base {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		} else {
			i++
		}
	}

	if arr[i] < arr[start] {
		arr[start], arr[i] = arr[i], arr[start]
	}

	sq(arr, start, i-1)
	sq(arr, i, end)
}

func main() {
	arr := []int{12, 3, 5, 8, 3, 2}
	sq(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
