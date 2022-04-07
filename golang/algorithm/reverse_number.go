package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	fmt.Println("origin:", x)
	y := x
	var arr []int
	for {
		z := x % 10
		x = x / 10
		arr = append(arr, z)

		if x == 0 {
			break
		}
	}

	fmt.Println(arr)

	max := []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 7}
	min := []int{-2, -1, -4, -7, -4, -8, -3, -6, -4, -8}

	if y > 0 {
		if len(arr) == len(max) {
			for i := 0; i < len(arr); i++ {
				if arr[i] > max[i] {
					return 0
				} else if arr[i] < max[i] {
					break
				} else {
					continue
				}
			}
		}
	} else if y < 0 {
		if len(arr) == len(min) {
			for i := 0; i < len(arr); i++ {
				if arr[i] < min[i] {
					return 0
				} else if arr[i] > min[i] {
					break
				} else {
					continue
				}
			}
		}
	}

	var s int
	for _, v := range arr {
		s = s*10 + v
	}

	return s
}

func main() {
	//a := 0x7FFFFFFF // 2147483647
	//b := 0x80000000 // -2147483648
	//fmt.Println(int32(a))
	//fmt.Println(int32(b))

	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(math.MaxInt32))
	fmt.Println(reverse(math.MinInt32))

	fmt.Println(reverse(-2147483412))
}
