package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {
	fmt.Println("ddd", stu)
}

func live() People {
	var stu *Student
	return stu
}

func main() {
	var i1 interface{}
	var s1 *Student
	var p1 People

	p1 = s1

	fmt.Printf("1%v\n", s1)
	fmt.Printf("2%v\n", i1)
	fmt.Printf("3%v\n", p1)

	fmt.Printf("4%v\n", i1 == nil)
	fmt.Printf("5%v\n", s1 == nil)
	fmt.Printf("6%v\n", p1 == nil)

	fmt.Printf("7%v\n", s1 == i1)
	fmt.Printf("8%v\n", p1 == i1)
	fmt.Printf("9%v\n", s1 == p1)

	l := live()

	fmt.Printf("%v\n", l)
	fmt.Printf("%v\n", l == nil)

	if l == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}

	s1.Show()
}
