package main

import (
	"fmt"
)

func do1() {
	defer func() {
		fmt.Println("111")
	}()

	defer func() {
		fmt.Println("222")
	}()

	defer func() {
		fmt.Println("333")
	}()
}

func f2() int {
	a := 1

	defer func() {
		a = 3
	}()

	return a
}

func do2() {
	b := f2()
	fmt.Println(b)
}

func f3() (a int) {
	a = 1

	defer func() {
		a = 3
	}()

	return a
}

func do3() {
	b := f3()
	fmt.Println(b)
}

func f4() (a int) {
	a = 1

	defer func() {
		a = 3
	}()

	return
}

func do4() {
	b := f4()
	fmt.Println(b)
}

func main() {
	do4()
}
