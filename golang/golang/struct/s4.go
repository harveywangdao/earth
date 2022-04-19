package main

type People interface {
	Eat()
}

type Student struct {
}

func (s *Student) Eat() {
}

// 编译不过
func main() {
	s1 := Student{}
	var p People = s1
	p.Eat()
}
