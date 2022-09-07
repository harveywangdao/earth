package main

import (
	"fmt"
	"unsafe"
)

type A1 struct {
	ZeroArr [0]int
}

type A2 struct {
	Age     int
	ZeroArr [0]int
}

type A3 struct {
	ZeroArr [0]int
	Age     int
}

func do1() {
	fmt.Printf("A1 Sizeof: %d, Alignof: %d\n", unsafe.Sizeof(A1{}), unsafe.Alignof(A1{}))
	fmt.Printf("A2 Sizeof: %d, Alignof: %d\n", unsafe.Sizeof(A2{}), unsafe.Alignof(A2{}))
	fmt.Printf("A3 Sizeof: %d, Alignof: %d\n", unsafe.Sizeof(A3{}), unsafe.Alignof(A3{}))
}

type ZeroArray [0]int
type ZeroStruct struct{}

func do2() {
	z1 := &ZeroArray{}
	z2 := &ZeroStruct{}
	fmt.Printf("%p, %p\n", z1, z2)
}

func main() {
	do2()
}
