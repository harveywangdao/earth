package main

import (
	"fmt"
)

func do1() {
	var ns []int
	var c int
	for i := 0; i < 10000; i++ {
		ns = append(ns, i)
		if c != cap(ns) {
			fmt.Println(len(ns), cap(ns))
			c = cap(ns)
		}
	}
}

func do2() {
	s := []int{1, 2}
	s = append(s, 4, 5, 6)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}

func do3() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5] // 2, 3, 4
	fmt.Println(len(s1), cap(s1))
	s2 := s1[2:6:7] //4, 5, 6, 7
	fmt.Println(s2)
	fmt.Println(len(s2), cap(s2))

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}

func do4() {
	a := make([]int, 4, 8)
	a[0] = 10
	a[1] = 11
	a[2] = 12
	a[3] = 13
	fmt.Println(a, len(a), cap(a))

	a1 := a[2:8]
	fmt.Println(a1, len(a1), cap(a1))
}

func do5() {
	a := make([]int, 8)
	a[0] = 10
	a[1] = 11
	a[2] = 12
	a[3] = 13
	a[4] = 14
	a[5] = 15
	a[6] = 16
	a[7] = 17
	fmt.Println(a, len(a), cap(a))

	a1 := a[2:4:8]
	fmt.Println(a1, len(a1), cap(a1))
}

func main() {
	do3()
}
