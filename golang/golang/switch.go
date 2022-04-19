package main

import (
	"fmt"
)

func do1(c int) {
	switch c {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}
}

func do2(c interface{}) {
	switch c.(type) {
	case int8:
		fmt.Println("int8")
	case int16:
		fmt.Println("int16")
	case int32:
		fmt.Println("int32")
	default:
		fmt.Println("default")
	}
}

func do3(a, b int) {
	switch {
	case a > b:
		fmt.Println("a > b")
	case a < b:
		fmt.Println("a < b")
	case a == b:
		fmt.Println("a == b")
	default:
		fmt.Println("default")
	}
}

func do4(a, b int) {
	switch {
	case a+1 > b:
		fmt.Println("a+1 > b")
	case a > b:
		fmt.Println("a > b")
	case a-1 > b:
		fmt.Println("a-1 > b")
	default:
		fmt.Println("default")
	}
}

func do5(c int) {
	switch c {
	case 1:
		println("1")
	case 2:
		println("2")
	case 3:
		println("3")
	default:
		println("default")
	}
}

func do6() {
	c1 := make(chan bool)
	c2 := make(chan bool)

	go func() {
		switch {
		case <-c1:
			println("1")
		case <-c2:
			println("2")
		default:
			println("default")
		}
	}()

	c2 <- true
	select {}
}

func do7(a, b int) {
	select {
	case a > b:
		fmt.Println("a > b")
	case a < b:
		fmt.Println("a < b")
	case a == b:
		fmt.Println("a == b")
	default:
		fmt.Println("default")
	}
}

func main() {
	do6()
}
