package main

import (
	"fmt"
)

func do1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	panic("12345")
}

func main() {
	do1()
}
