package main

import (
	"fmt"
)

func add() (sum int) {
	sum = 1
	defer func() {
		sum = 2
	}()

	return
}

func add2() int {
	sum := 1
	defer func() {
		sum = 2
	}()

	return sum + 4
}

func main() {
	fmt.Println(add())
	fmt.Println(add2())
}
