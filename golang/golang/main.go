package main

import (
	"fmt"
)

func alignUp(n, a uintptr) uintptr {
	return (n + a - 1) &^ (a - 1)
}

func main() {
	fmt.Println(alignUp(31, 0))
}
