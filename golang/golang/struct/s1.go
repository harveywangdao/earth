package main

type People struct {
	Name string
	Age  int
}

type Teacher struct {
	People
	School string
}

// 结构体嵌套不能相互赋值
func do1() {
	//var p *People
	//p = &Teacher{}

	//var t *Teacher
	//t = &People{}
}

func main() {
	do1()
}
