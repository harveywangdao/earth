package main

type Apple struct {
	Weight int
}

type People struct {
	Age   int
	Fruit *Apple
}

func newApple() *Apple {
	a1 := &Apple{
		Weight: 23,
	}
	return a1
}

func newPeople1() *People {
	p1 := &People{
		Age: 23,
	}
	return p1
}

func newPeople2() *People {
	a1 := newApple()
	p1 := &People{
		Age: 23,
	}
	p1.Fruit = a1
	return p1
}

func f1(p1 *People) {
	p1.Fruit = nil
}

func f2(p1 *People, a1 *Apple) {
	p1.Fruit = a1
}

// 堆-删除写屏障 p1是堆对象,a1是堆对象
func do1() {
	p1 := newPeople2()
	f1(p1)
}

// 堆-插入写屏障 p1是堆对象,a1是堆对象
func do2() {
	p1 := newPeople1()
	a1 := newApple()
	f2(p1, a1)
}

// 栈-删除写屏障 p1是栈对象,a1是堆对象
func do3() {
	a1 := newApple()
	p1 := &People{
		Age:   12,
		Fruit: a1,
	}
	f1(p1)
}

// 栈-插入写屏障 p1是栈对象,a1是堆对象
func do4() {
	a1 := newApple()
	p1 := &People{
		Age: 12,
	}
	f2(p1, a1)
}

func main() {
	do3()
}
