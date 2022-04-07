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

	a := []int{2}
	a[33] = 9
}

func do2() {
	var a interface{}
	a = 3
	b, ok := a.(string)
	fmt.Println(b, ok)
}

func main() {
	do2()
}
