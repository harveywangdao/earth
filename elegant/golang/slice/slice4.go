package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func printslice(s *[]int) {
	p := (*reflect.SliceHeader)(unsafe.Pointer(s))
	fmt.Printf("ptr: %v, len: %v, cap: %v\n", p.Data, p.Len, p.Cap)
}

func printmem() {
	m := runtime.MemStats{}
	runtime.ReadMemStats(&m)
	fmt.Println("mem:", m.Alloc)
}

func f1() []int {
	printmem()
	s1 := make([]int, 1024*1024*100) // 100MB*8
	printmem()
	printslice(&s1)
	s2 := s1[len(s1)-1:]
	printslice(&s2)
	return s2
}

func do1() {
	s1 := f1()
	printslice(&s1)
	runtime.GC()
	printmem()
	s1[0] = 199
	printslice(&s1)
	printmem()
}

func f2() []int {
	s2 := []int{}
	s3 := []int{}
	printslice(&s2)
	printslice(&s3)
	return s2
}

func do2() {
	f2()
}

func main() {
	do2()
}
