package main

import (
	"fmt"
)

type People struct {
	Name string
	Age  int
}

func (p People) SetName1(s string) {
	p.Name = s
}

func (p *People) SetName2(s string) {
	p.Name = s
}

func do1() {
	p1 := People{Name: "xiaoming"}
	p1.SetName1("1111")
	fmt.Println(p1.Name)
	p1.SetName2("2222")
	fmt.Println(p1.Name)

	p2 := &People{Name: "xiaohong"}
	p2.SetName1("3333")
	fmt.Println(p2.Name)
	p2.SetName2("4444")
	fmt.Println(p2.Name)
}

func main() {
	do1()
}
