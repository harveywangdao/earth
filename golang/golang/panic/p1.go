package main

import (
	"fmt"
)

func do1() {
	panic("12345")
}

func do2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()
	go do1()
	select {}
}

func main() {
	do2()
}
