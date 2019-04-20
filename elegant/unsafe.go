package main

import (
	"fmt"
	"unsafe"
)

type ClassA struct {
	a int32
	b byte
	c int
	d bool
}

func main() {
	var bytea byte = 7
	var int8a int8 = 7
	var uint8a uint8 = 7
	var int16a int16 = 7
	var uint16a uint16 = 7
	var int32a int32 = 7
	var uint32a uint32 = 7
	var int64a int64 = 7
	var uint64a uint64 = 7
	var runea rune = 7
	var inta int = 7
	var uinta uint = 7
	var float32a float32 = 7
	var float64a float64 = 7
	var boola bool = true
	var complex64a complex64 = 1
	var complex128a complex128 = 1
	var stringa string = "dsdf"

	fmt.Println("bytea =", unsafe.Sizeof(bytea))
	fmt.Println("int8a =", unsafe.Sizeof(int8a))
	fmt.Println("uint8a =", unsafe.Sizeof(uint8a))
	fmt.Println("int16a =", unsafe.Sizeof(int16a))
	fmt.Println("uint16a =", unsafe.Sizeof(uint16a))
	fmt.Println("int32a =", unsafe.Sizeof(int32a))
	fmt.Println("uint32a =", unsafe.Sizeof(uint32a))
	fmt.Println("int64a =", unsafe.Sizeof(int64a))
	fmt.Println("uint64a =", unsafe.Sizeof(uint64a))
	fmt.Println("runea =", unsafe.Sizeof(runea))
	fmt.Println("inta =", unsafe.Sizeof(inta))
	fmt.Println("uinta =", unsafe.Sizeof(uinta))
	fmt.Println("float32a =", unsafe.Sizeof(float32a))
	fmt.Println("float64a =", unsafe.Sizeof(float64a))
	fmt.Println("boola =", unsafe.Sizeof(boola))
	fmt.Println("complex64a =", unsafe.Sizeof(complex64a))
	fmt.Println("complex128a =", unsafe.Sizeof(complex128a))
	fmt.Println("stringa =", unsafe.Sizeof(stringa))

	var v1 complex128
	v1 = 1 + 2i
	v2 := complex(3, 4)

	fmt.Println("v1 =", v1)
	fmt.Println("v2 =", v2)
	fmt.Println("v1+v2 =", v1+v2)
	fmt.Println("real(v1) =", real(v1))
	fmt.Println("imag(v1) =", imag(v1))

	var a int = 9
	p := &a
	fmt.Println("a =", a)
	fmt.Println("p =", p)
	fmt.Println("*p =", *p)

	p2 := unsafe.Pointer(p)
	fmt.Println("p2 =", p2)

	p3 := (*int)(p2)
	fmt.Println("*p3 =", *p3)

	p4 := uintptr(p2)
	fmt.Println("p4 =", p4)

	p4++
	fmt.Println("p4 =", p4)

	p5 := unsafe.Pointer(p4)
	fmt.Println("p5 =", p5)

	p6 := (*int)(p5)
	fmt.Println("p6 =", p6)
	fmt.Println("*p6 =", *p6)

	ca := ClassA{
		a: 1,
		b: 3,
		c: 45,
		d: false,
	}

	fmt.Printf("%+v\n", ca)
	fmt.Println("sizeof(ca) =", unsafe.Sizeof(ca))
	fmt.Println("sizeof(&ca) =", unsafe.Sizeof(&ca))
	fmt.Println("Alignof(ca) =", unsafe.Alignof(ca))
	fmt.Println("Alignof(&ca) =", unsafe.Alignof(&ca))
	fmt.Println("Alignof(ca.a) =", unsafe.Alignof(ca.a))
	fmt.Println("Alignof(ca.b) =", unsafe.Alignof(ca.b))
	fmt.Println("Alignof(ca.c) =", unsafe.Alignof(ca.c))
	fmt.Println("Alignof(ca.d) =", unsafe.Alignof(ca.d))
	fmt.Println("Offsetof(ca.a) =", unsafe.Offsetof(ca.a))
	fmt.Println("Offsetof(ca.b) =", unsafe.Offsetof(ca.b))
	fmt.Println("Offsetof(ca.c) =", unsafe.Offsetof(ca.c))
	fmt.Println("Offsetof(ca.d) =", unsafe.Offsetof(ca.d))

	ptr := unsafe.Pointer(&ca)
	ptr1 := (*int32)(ptr)
	*ptr1 = 9

	ptr2 := (*byte)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(ca.b)))
	*ptr2 = 9

	ptr3 := (*int)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(ca.c)))
	*ptr3 = 9

	ptr4 := (*byte)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(ca.d)))
	*ptr4 = 9

	fmt.Printf("%+v\n", ca)
}
