package main

import (
	"fmt"
)

func do1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("AAAAAAAAAAA", r)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("BBBBBBBBBBB", r)
			var b *int
			*b = 2
		}
	}()

	var p *int
	*p = 2
}

func main() {
	do1()
}
