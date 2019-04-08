package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)
	m1 := make(map[int]string, 10)
	m[2] = "fadf"
	m1[12] = "afasf"

	fmt.Println(len(m), m)
	fmt.Println(len(m1), m1)

	s1 := make([]int, 10)
	s2 := make([]int, 10, 20)

	fmt.Println(len(s1), cap(s1), s1)
	fmt.Println(len(s2), cap(s2), s2)

	c1 := make(chan int)
	c2 := make(chan int, 10)

	c2 <- 21

	fmt.Println(len(c1), cap(c1), c1)
	fmt.Println(len(c2), cap(c2), c2)

	mnew := *new(map[int]string)
	//mnew[4564] = "sdvsdv"
	fmt.Println(len(mnew), mnew)

	snew := *new([]int)
	//snew = append([]int(nil), 1)
	snew = append(snew, 1)
	fmt.Println(len(snew), cap(snew), snew)

	cnew := *new(chan int)
	//<-cnew
	//cnew <- 8
	fmt.Println(len(cnew), cap(cnew), cnew)
}
