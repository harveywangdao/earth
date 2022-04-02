package main

import (
	"fmt"
)

//go:generate echo hello
func main() {
	var n Life = Human
	fmt.Println(n.String())
}
