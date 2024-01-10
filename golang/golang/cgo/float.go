package main

/*
#include <stdio.h>
#include <stdlib.h>

void print_float(float *arr, int n)
{
  for (int i=0; i<n; i++)
  {
    printf("%f\n", arr[i]);
  }
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func CFloats(f []float32) unsafe.Pointer {
	n := len(f)
	p := C.malloc(C.size_t(n * 4))
	sliceHeader := struct {
		p   unsafe.Pointer
		len int
		cap int
	}{p, n, n}
	s := *(*[]float32)(unsafe.Pointer(&sliceHeader))
	copy(s, f)
	return p
}

func GoFloats(p unsafe.Pointer, n int) []float32 {
	if n <= 0 {
		return make([]float32, 0)
	}
	sliceHeader := struct {
		p   unsafe.Pointer
		len int
		cap int
	}{p, n, n}
	s := *(*[]float32)(unsafe.Pointer(&sliceHeader))
	f := make([]float32, n)
	copy(f, s)
	return f
}

// export CGO_ENABLED=1
func main() {
	f1 := []float32{1.2, 1.3, 1.8}
	p1 := (*C.float)(CFloats(f1))
	C.print_float(p1, 3)

	f2 := GoFloats(unsafe.Pointer(p1), 3)
	fmt.Println(f2)

	str := "hello"
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.puts(cstr)
}
