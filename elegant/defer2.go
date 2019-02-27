package main

import (
	"fmt"
)

func add() (sum int) {
	sum = 1
	defer func() {
		sum = 2
	}()

	return
}

func add2() int {
	sum := 1
	defer func() {
		sum = 2
	}()

	return sum + 4
}

func add3() (sum int) {
	sum = 1
	defer func() {
		sum = 3
	}()

	return 10
}

func add4() (sum int) {
	sum = 1
	defer func() {
		sum = 4
	}()

	return sum
}

func add5() (sum int) {
	sum = 1
	defer func() {
		sum = 5
	}()

	return sum + 70
}

func main() {
	fmt.Println(add())  //2
	fmt.Println(add2()) //5
	fmt.Println(add3()) //3
	fmt.Println(add4()) //4
	fmt.Println(add5()) //5
}
