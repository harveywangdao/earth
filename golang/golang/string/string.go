package main

import (
	"fmt"
	"strings"
)

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func do1() {
	p := &People{}
	p.String()
}

func do2() {
	ss := strings.Split("/dcs/asdfas/asdsa/", "/")
	fmt.Println(len(ss))
	fmt.Println(ss)
}

func main() {
	do2()
}
