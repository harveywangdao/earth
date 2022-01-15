package main

import (
	"fmt"
	"math"
)

func do1() {
	m := map[int]int{}
	for i := 0; i < 5; i++ {
		m[i] = i
	}

	n := 5
	for k, v := range m {
		fmt.Println(k, v)
		m[n] = n
		n++
	}

	fmt.Println(m)
}

func do2() {
	m := map[int]int{}
	m[1] = 1
	// /fmt.Println(&m[1])
}

func do3() {
	m := make(map[float64]int)
	m[1.4] = 1
	m[2.4] = 2
	m[math.NaN()] = 3
	m[math.NaN()] = 3

	for k, v := range m {
		fmt.Printf("[%v, %d] ", k, v)
	}

	fmt.Printf("\nk: %v, v: %d\n", math.NaN(), m[math.NaN()])
	fmt.Printf("k: %v, v: %d\n", 2.400000000001, m[2.400000000001])
	fmt.Printf("k: %v, v: %d\n", 2.4000000000000000000000001, m[2.4000000000000000000000001])

	fmt.Println(math.NaN() == math.NaN())
}

func do4() {
	m := make(map[float64]int)
	m[2.4] = 2

	fmt.Println(math.Float64bits(2.4))
	fmt.Println(math.Float64bits(2.400000000001))
	fmt.Println(math.Float64bits(2.4000000000000000000000001))
}

func do5() {
	type S struct {
		ID int
	}
	s1 := S{ID: 1}
	s2 := S{ID: 1}

	var h = map[*S]int{}
	h[&s1] = 1
	fmt.Println(h[&s1])
	fmt.Println(h[&s2])
	fmt.Println(s1 == s2)
}

func modify1(x map[int]int) {
	x = make(map[int]int)
	x[2] = 2
}

func modify2(x map[int]int) {
	x[2] = 2
}

func do6() {
	m := make(map[int]int)
	m[1] = 1
	fmt.Println(m)
	modify1(m)
	fmt.Println(m)
	modify2(m)
	fmt.Println(m)
}

func modify3(x []int) {
	x = make([]int, 2)
	x[0] = 10
	x[1] = 20
}

func modify4(x []int) {
	x[0] = 100
	x[1] = 200
}

func modify5(x []int) {
	x[0] = 1000
	x[1] = 2000
	x = append(x, 3000)
}

func do7() {
	m := make([]int, 2)
	m[0] = 1
	m[1] = 2
	fmt.Println(m)
	modify3(m)
	fmt.Println(m)
	modify4(m)
	fmt.Println(m)
	modify5(m)
	fmt.Println(m)
}

func main() {
	do7()
}
