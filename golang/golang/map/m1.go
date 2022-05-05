package main

func do1() {
	m := make(map[string]int, 10)
	m["2"] = 3
	a := m["2"]
	_ = a
	delete(m, "2")
}

func main() {
	do1()
}
