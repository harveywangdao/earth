package main

import (
	"fmt"
)

func do1() (n int) {
	n = 12
	return n
}

func do2() (n int) {
	n = 34
	return
}

func do3() (n int) {
	n = 56
	return 78
}

func main() {
	a := do1()
	b := do2()
	c := do3()
	fmt.Println(a, b, c)
}
