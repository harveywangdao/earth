package main

import (
	"fmt"
	"strings"
)

func do1() {
	s1 := "123"
	b1 := []byte{'a', 'b'}
	b1 = append(b1, 'c')
	b1 = append(b1, s1...)
	b1 = append(b1, b1...)
	fmt.Println(b1)
}

func do2() {
	b := strings.Builder{}
	b.Grow(36)
	b.WriteString("123")
}

func do3() {
	s := "\x00\x00\x00\x00"
	fmt.Println(s)
	fmt.Println(len(s))
}

func main() {
	do3()
}
