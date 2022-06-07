package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func do1() {
	a := "asdas"
	b := "asdas"
	c := "asdas"

	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &c)

	s1 := (*reflect.StringHeader)(unsafe.Pointer(&a))
	s2 := (*reflect.StringHeader)(unsafe.Pointer(&b))
	s3 := (*reflect.StringHeader)(unsafe.Pointer(&c))

	fmt.Println(s1.Data)
	fmt.Println(s2.Data)
	fmt.Println(s3.Data)
}

func main() {
	do1()
}
