package main

import (
	"fmt"
)

type ftype func()

func aa() {
	fmt.Println("aa")
}

func bb() {
	fmt.Println("bb")
}

func do() {
	var f1 ftype

	f1 = aa
	defer f1()

	f1 = bb
	defer f1()
}

func main() {
	do()
}
