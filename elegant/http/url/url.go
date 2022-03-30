package main

import (
	"fmt"
	"html"
	"net/url"
)

func do1() {
	u1 := "http://aas.ccc.com?aa=1&bb=2&cc=3"
	e1 := url.QueryEscape(u1)
	fmt.Println(e1)

	e2 := url.PathEscape(u1)
	fmt.Println(e2)

	e3, _ := url.PathUnescape(u1)
	fmt.Println(e3)

	e4, _ := url.QueryUnescape(u1)
	fmt.Println(e4)

	e5 := html.EscapeString(u1)
	fmt.Println(e5)
}

func main() {
	do1()
}
