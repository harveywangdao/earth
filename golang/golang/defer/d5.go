package main

import (
	"fmt"
)

func f1(x, y int) {
	println(x, y)
}

func f2(x, z int) (y int) {
	return x + z + 22
}

func do1() int {
	a := 11
	b := 13

	defer f1(a, b)
	defer f2(a, b)
	fmt.Println()
	a = 14
	return a
}

func do2() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
		println("111")
		panic("ddd")
	}()

	defer func() {
		/*if e := recover(); e != nil {
			fmt.Println(e)
		}*/
		println("222")
		panic("aaa")
		println("333")
	}()

	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
		println("444")
		panic("bbb")
		println("555")
	}()

	defer func() {
		println("666")
	}()

	panic("ccc")
}

func main() {
	do2()
}
