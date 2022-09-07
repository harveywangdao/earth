package main

import (
	"fmt"
	"unsafe"
)

func do1() {
	fn1 := func(n int) int {
		return n + 2
	}
	fmt.Printf("Sizeof %d\n", unsafe.Sizeof(fn1))
}

func do2() {
	age := 10
	fn1 := func() int {
		return age + 3
	}
	age = 8
	fn1()
}

func addfn(add1 int) func(int) int {
	add2 := 5
	fn1 := func(add3 int) int {
		add4 := 4
		return add1 + add2 + add3 + add4
	}
	add2 = 10
	return fn1
}

func do3() {
	m1 := addfn(6)
	m2 := addfn(7)

	//fmt.Printf("%T, %T\n", m1, m2)

	a1 := m1(11)
	a2 := m2(23)

	a3 := a1 + a2
	_ = a3
}

func main() {
	do3()
}
