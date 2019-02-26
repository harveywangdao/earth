package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(1, err)
		}

		fmt.Println("打印前")
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(2, err)
		}

		fmt.Println("打印中")
	}()

	defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")
}
