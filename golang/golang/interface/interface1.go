package main

type People struct {
	Name string
}

func do1() {
	var p1 *People
	println(p1 == nil) //true

	var i1 interface{}
	i1 = p1
	println(i1 == nil) //false
}

func do2() {
	var p1 People
	var i1 interface{}
	i1 = p1
	_ = i1
}

func do3() {
	s1 := "345"
	var i1 interface{}
	i1 = s1

	s2, ok := i1.(string)
	println(s2, ok)
}

func main() {
	do1()
}
