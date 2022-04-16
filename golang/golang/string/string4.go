package main

import (
	"fmt"
)

func do1() {
	s1 := "a我是猪b"
	for k, v := range s1 {
		fmt.Printf("%T, %T, %v, 0x%X\n", k, v, k, v)
	}

	s2 := []byte(s1)
	for k, v := range s2 {
		fmt.Printf("%T, %T, %v, 0x%X\n", k, v, k, v)
	}

	s3 := []rune(s1)
	for k, v := range s3 {
		fmt.Printf("%T, %T, %v, 0x%X\n", k, v, k, v)
	}
}

func do2() {
	s1 := []byte{0xe6, 0x88, 0x91} //utf-8
	s2 := string(s1)
	fmt.Println(s2)

	s3 := []rune{0x6211} //unicode
	s4 := string(s3)
	fmt.Println(s4)
}

func main() {
	do2()
}
