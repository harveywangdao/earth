package main

import (
	"fmt"
)

func do1() {
	s1 := make([]int, 0, 8)
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	s1 = append(s1, 4)
	s1 = append(s1, 5)
	fmt.Println(s1)
	s2 := s1[3:3]
	s3 := s1[3:4]
	s4 := s1[3:5]
	s5 := s1[3:6]
	s6 := s1[3:7]
	s7 := s1[3:8]
	//s8 := s1[3:9]
	fmt.Println("s2:", len(s2), cap(s2))
	fmt.Println("s3:", len(s3), cap(s3))
	fmt.Println("s4:", len(s4), cap(s4))
	fmt.Println("s5:", len(s5), cap(s5))
	fmt.Println("s6:", len(s6), cap(s6))
	fmt.Println("s7:", len(s7), cap(s7))
	//fmt.Println("s8:", len(s8), cap(s8))
}

func main() {
	do1()
}
