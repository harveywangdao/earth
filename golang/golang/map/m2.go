package main

import (
	"fmt"
)

func do1() {
	var m map[int]int
	if m == nil {
		fmt.Println("m==nil")
	}
	a := m[3]     // mapaccess1_fast64
	b, ok := m[4] // mapaccess2_fast64
	fmt.Println(a, b, ok)

	m[2] = 3 // mapassign_fast64
}

func main() {
	do1()
}
