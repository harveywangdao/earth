package main

type Dog struct {
	Name string
}

func (d *Dog) GetName() string {
	return d.Name
}

func (d *Dog) Bark() {
}

type Cat struct {
	Age int
}

func (c Cat) GetAge() int {
	return c.Age
}

func (c Cat) Meow() {
}

func do1() {
	d1 := Dog{Name: "sss"}
	d2 := &Dog{Name: "fff"}
	d1.GetName()
	d2.GetName()

	c1 := Cat{Age: 12}
	c2 := &Cat{Age: 13}
	c1.GetAge()
	c2.GetAge()
}

func do2() {
	var d *Dog
	//d.GetName()
	d.Bark()

	//var c *Cat
	//c.GetAge()
	//c.Meow()
}

func main() {
	do2()
}
