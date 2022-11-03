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

func do2() {
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

// 字节切片被转成base64
func do3() {
	p := &People{
		Name: "xiaoming",
		Age:  12,
	}
	data, _ := json.Marshal(p)

	m := make(map[string]interface{})
	m["data1"] = string(data)
	m["data2"] = data

	data2, _ := json.Marshal(m)
	fmt.Println(string(data2))

	m2 := make(map[string]interface{})
	json.Unmarshal(data2, &m2)
	fmt.Println(m2)

	type Human struct {
		Data1 string `json:"data1"`
		Data2 []byte `json:"data2"`
	}
	h1 := Human{}
	json.Unmarshal(data2, &h1)
	fmt.Println(h1.Data1, string(h1.Data2))

	type Human2 struct {
		Data1 string `json:"data1"`
		Data2 string `json:"data2"`
	}
	h2 := Human2{}
	json.Unmarshal(data2, &h2)
	fmt.Println(h2.Data1, h2.Data2)
}

func main() {
	do3()
}
