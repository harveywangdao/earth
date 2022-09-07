package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type EmptyStructArray [32]struct{}

func do1() {
	fmt.Printf("%d\n", unsafe.Sizeof(EmptyStructArray{}))
}

func ps(s []struct{}) {
	p := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println(p.Data, p.Len, p.Cap)
}

func do2() {
	s := make([]struct{}, 64)
	_ = s
}

func do3() {
	s := make([]int, 64)
	_ = s
}

func main() {
	do2()
}
