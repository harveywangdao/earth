package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "123456789哈"

	fmt.Println(len(s))
	fmt.Println(len([]byte(s)))

	fmt.Println(len([]rune(s)))
	fmt.Println(bytes.Count([]byte(s), nil) - 1)
	fmt.Println(strings.Count(s, "") - 1)
	fmt.Println(utf8.RuneCountInString(s))

	var b bytes.Buffer

	for i := 0; i < 300; i++ {
		b.WriteByte(byte(i%10) + '0')

		//fmt.Println(b.String())
		fmt.Println(b.Cap())
		fmt.Println(b.Len())
	}

	var a []int

	for i := 0; i < 300; i++ {
		a = append(a, i)

		//fmt.Println(a)
		fmt.Println(cap(a))
		fmt.Println(len(a))
	}
}
