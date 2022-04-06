package main

import (
	"fmt"
)

func do1() {
	fmt.Println("000")

	defer fmt.Println("111")
	defer fmt.Println("222")
	defer fmt.Println("333")

	/*defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()*/

	panic("haha")

	fmt.Println("444")
}

func do2() {
	a := 1

	defer func() {
		f1(a)
	}()

	defer func(x int) {
		f1(x)
	}(a)

	a++
}

func f1(a int) {
	fmt.Println(a)
}

func main() {
	do2()
}
