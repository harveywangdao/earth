package main

import (
	"fmt"
)

func do1() {
	var a uint64 = 1
	var b uint64 = 2 << 32
	var c uint64 = a + b

	d := uint32(c)
	fmt.Printf("0x%X\n", a)
	fmt.Printf("0x%X\n", b)
	fmt.Printf("0x%X\n", c)
	fmt.Printf("0x%X\n", d)
}

func main() {
	do1()
}
