package main

import (
	"fmt"
)

const (
	AA = iota
	BB
	CC = "cc"
	DD = "dd"
	EE
	FF
	GG = iota
	HH
)

type People struct {
	Name string
}

func f1(arr [3]int) {
	arr[1] = 99
	fmt.Println(arr)
}

func f2(arr []int) {
	arr[1] = 99
	fmt.Println(arr)
}

func main() {
	p1 := People{"xiaoming"}

	p2 := p1

	p2.Name = "xiaohong"

	fmt.Println(p1, p2)

	fmt.Println(AA, BB, CC, DD, EE, FF, GG, HH)

	b := [3]int{1, 2, 3}
	f1(b)
	fmt.Println(b)

	c := []int{1, 2, 3}
	f2(c)
	fmt.Println(c)
}
