package main

type People struct {
	Name string
	Age  *int
}

func do1() {
	s1 := "xiaoming"
	age := 13
	p1 := &People{
		Name: s1,
		Age:  &age,
	}
	_ = p1
}

func do2() {
	s1 := "xiaoming"
	age := 13
	p1 := &People{}
	p1.Name = s1
	p1.Age = &age
}

func main() {
	do1()
	do2()
}
