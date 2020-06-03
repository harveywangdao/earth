package main

import (
	"fmt"
)

type People struct {
	Name string
	Age  int
}

func (p *People) SetName(n string) {
	p.Name = n
}

func (p People) SetAge(a int) {
	p.Age = a
}

func main() {
	p := People{
		Name: "xiaoming",
		Age:  12,
	}

	fmt.Println(p)

	p.SetAge(13)
	p.SetName("xiaohang")

	fmt.Println(p)

	p2 := &People{
		Name: "wang",
		Age:  20,
	}

	fmt.Println(p2)

	p2.SetAge(21)
	p2.SetName("zhang")

	fmt.Println(p2)
}
