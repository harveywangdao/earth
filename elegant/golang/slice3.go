package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func do1() {
	var s []int
	p := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println("nil slice:", p.Data, p.Len, p.Cap)
}

func do2() {
	s := []int{}
	p := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println("[]int{} slice", p.Data, p.Len, p.Cap)
}

func do3() {
	s := []int{1, 2}
	p := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println("[]int{1,2} slice", p.Data, p.Len, p.Cap)
}

func do4() {
	s := make([]int, 0)
	p := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println("make([]int, 0) slice", p.Data, p.Len, p.Cap)
}

func do5() {
	s := make([]int, 8)
	p := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println("make([]int, 8) slice", p.Data, p.Len, p.Cap)
}

func main() {
	do1()
	do2()
	do3()
	do4()
	do5()
}
