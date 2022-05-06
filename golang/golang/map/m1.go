package main

func do1() {
	m := make(map[string]int, 10)
	m["2"] = 3
	a := m["2"]
	_ = a
	delete(m, "2")
}

func do2() {
	m1 := make(map[int]int, 64)
	m2 := make(map[int32]int, 64)
	m3 := make(map[int64]int, 64)
	m4 := make(map[float64]int, 64)

	m5 := make(map[bool]int, 64)

	m6 := make(map[interface{}]int, 64)

	m7 := make(map[string]int, 64)
	//m8 := make(map[[]int]int, 64)
	m8 := make(map[int][]int, 64)
	m9 := make(map[[3]int]int, 64)

	m10 := make(map[chan int]int, 64)

	//m11 := make(map[map[int]int]int, 64)
	m11 := make(map[int]map[int]int, 64)

	m12 := make(map[*int]int, 64)

	_ = m1
	_ = m2
	_ = m3
	_ = m4
	_ = m5
	_ = m6
	_ = m7
	_ = m8
	_ = m9
	_ = m10
	_ = m11
	_ = m12
}

func do3() {
	m1 := make(map[int]int)
	m2 := make(map[int]int, 64)
	m1[12] = 13
	m2[22] = 33
}

func do4() {
	m1 := make(map[float64]int, 64)
	m1[22.0] = 33
}

func main() {
	do1()
}
