package main

import (
	"fmt"
)

type Dog struct {
	Name string
}

func (d *Dog) SetName(s string) {
	d.Name = s
}

type Cat struct {
	Age int
}

func (c Cat) SetAge(n int) {
	c.Age = n
}

func do1() {
	d1 := Dog{Name: "111"}
	d2 := &Dog{Name: "222"}
	d1.SetName("333")
	d2.SetName("444")
	fmt.Println(d1.Name)
	fmt.Println(d2.Name)

	c1 := Cat{Age: 11}
	c2 := &Cat{Age: 22}
	c1.SetAge(33)
	c2.SetAge(44)
	fmt.Println(c1.Age)
	fmt.Println(c2.Age)
}

func main() {
	do1()
}
