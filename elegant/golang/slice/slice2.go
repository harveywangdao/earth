package main

import (
	"fmt"
)

func do1() {
	fmt.Println("slice []int:")
	var s []int
	n := 0
	for i := 0; i < 10000; i++ {
		s = append(s, i)
		if n != cap(s) {
			if n == 0 {
				fmt.Println(cap(s))
			} else {
				fmt.Println(cap(s), float64(cap(s))/float64(n))
			}
			n = cap(s)
		}
	}
	fmt.Println()
}

func do2() {
	fmt.Println("slice []int32:")
	var s []int32
	n := 0
	for i := 0; i < 10000; i++ {
		s = append(s, int32(i))
		if n != cap(s) {
			if n == 0 {
				fmt.Println(cap(s))
			} else {
				fmt.Println(cap(s), float64(cap(s))/float64(n))
			}
			n = cap(s)
		}
	}
	fmt.Println()
}

func do3() {
	fmt.Println("slice []int16:")
	var s []int16
	n := 0
	for i := 0; i < 10000; i++ {
		s = append(s, int16(i))
		if n != cap(s) {
			if n == 0 {
				fmt.Println(cap(s))
			} else {
				fmt.Println(cap(s), float64(cap(s))/float64(n))
			}
			n = cap(s)
		}
	}
	fmt.Println()
}

func do4() {
	fmt.Println("slice []int8:")
	var s []int8
	n := 0
	for i := 0; i < 10000; i++ {
		s = append(s, int8(i))
		if n != cap(s) {
			if n == 0 {
				fmt.Println(cap(s))
			} else {
				fmt.Println(cap(s), float64(cap(s))/float64(n))
			}
			n = cap(s)
		}
	}
	fmt.Println()
}

func main() {
	do1()
	do2()
	do3()
	do4()
}
