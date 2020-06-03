package main

import (
	"fmt"
)

type ClassA struct {
	Pig string
}

type InterfaceA interface {
	PrintName()
}

func (c ClassA) PrintName() {
	fmt.Println(c.Pig)
}

func main() {
	var ca InterfaceA
	ca = &ClassA{
		Pig: "xiaozhu",
	}

	ca.PrintName()
}
