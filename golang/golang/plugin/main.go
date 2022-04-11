package main

import (
	"fmt"
	"plugin"
)

func do1() {
	p, err := plugin.Open("add.so")
	if err != nil {
		panic(err)
	}

	s, err := p.Lookup("Add")
	if err != nil {
		panic(err)
	}

	addFunc, ok := s.(func(x, y int) int)
	if !ok {
		panic("func err")
	}
	r := addFunc(1, 2)
	fmt.Println(r)
}

func main() {
	do1()
}
