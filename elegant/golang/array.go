package main

import (
	"fmt"
)

func main() {
	s1 := []int{2, 4, 564, 756, 578, 67, 8}
	s2 := s1
	s3 := s1[:6]
	s4 := make([]int, 0, len(s1))
	s4 = append(s4, s1...)

	s2[0] = 777
	s3[0] = 888
	s4[0] = 999

	fmt.Printf("s1: %p\n", &s1)
	fmt.Printf("s2: %p\n", &s2)
	fmt.Printf("s3: %p\n", &s3)
	fmt.Printf("s4: %p\n", &s4)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	for i, n := range s1 {
		fmt.Println(i, n)
		if i < 6 {
			s1[i+1] = 233
		}
	}

	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)
	fmt.Println("s4:", s4)

	a := [3]int{1, 2, 3}
	for i, v := range a { //i,v从a复制的对象里提取出
		if i == 0 {
			a[1], a[2] = 200, 300
			fmt.Println(a) //输出[1 200 300]
		}
		a[i] = v + 100 //v是复制对象里的元素[1, 2, 3]
	}
	fmt.Println(a) //输出[101, 102, 103]

	b := []int{1, 2, 3}
	for i, v := range b { //i,v从a复制的对象里提取出
		if i == 0 {
			b[1], b[2] = 200, 300
			fmt.Println(b) //输出[1 200 300]
		}
		b[i] = v + 100 //v是复制对象里的元素[1, 2, 3]
	}
	fmt.Println(b) //输出[101, 300, 400]

	c := [3]int{1, 2, 3}
	d := c
	e := c[:]
	e = append(e, 4)

	c[0] = 888
	d[0] = 999

	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
}
