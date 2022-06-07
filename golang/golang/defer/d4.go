package main

import (
	"fmt"
)

func f1() (x int) {
	defer fmt.Println(x) //0
	x = 7
	return 9
}

func f2() (x int) {
	defer func() {
		fmt.Println(x) //9
	}()
	x = 7
	return 9
}

func f3() (x int) {
	x = 7
	defer fmt.Println(x) //7
	return 9
}

func f4() (x int) {
	defer func(n int) {
		fmt.Println(n, x) //0,9
	}(x)
	x = 7
	return 9
}

func do1() {
	fmt.Println("ret:", f1())
	fmt.Println("ret:", f2())
	fmt.Println("ret:", f3())
	fmt.Println("ret:", f4())
}

func main() {
	do1()
}
