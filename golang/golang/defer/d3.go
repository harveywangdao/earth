package main

import (
	"fmt"
)

func do1() (a int) {
	a = 11

	defer func() {
		a = 22
	}()

	defer func() {
		a = 33
	}()

	return a
}

func main() {
	a := do1()
	fmt.Println(a)
}
