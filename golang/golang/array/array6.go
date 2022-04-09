package main

func do1() {
	arr1 := [...]int{3, 4, 5}
	p1 := &arr1
	x := len(arr1)
	y := len(p1)
	_, _ = x, y
}

func do2() {
	n := 45
	_ = n
}

func main() {
	do1()
}
