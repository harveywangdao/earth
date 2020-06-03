package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type People struct {
	Name string
	Age  int
}

func add(a, b int) int {
	fmt.Println(a, b)
	return a + b
}

func do1() {
	defer fmt.Println(add(100, 200))

	defer func() {
		fmt.Println(add(10, 20))
	}()

	fmt.Println(add(1, 2))
}

func main() {
	do1()
	return
	a := People{
		Name: "xiaoming",
		Age:  12,
	}

	data, _ := json.Marshal(a)
	fmt.Println("data1 =", string(data))

	data, _ = json.Marshal(&a)
	fmt.Println("data2 =", string(data))

	n := 1
	m := 1

	for n = n + 1; n > m; {
		fmt.Println("n:", n)
		fmt.Println("m:", m)
		time.Sleep(1 * time.Second)
	}
}
