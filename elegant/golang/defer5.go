package main

import (
	"fmt"
)

func es() {
	fmt.Println("000")

	defer fmt.Println("111")
	defer fmt.Println("222")
	defer fmt.Println("333")

	/*defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()*/

	panic("haha")

	fmt.Println("444")
}

func main() {
	es()
	fmt.Println("555")
}
