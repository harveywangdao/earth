package main

func getf1() func() {
	n := 1 //n escapes to heap
	return func() {
		n++
		println(n)
	}
}

func do1() {
	f1 := getf1()
	f1()
	f1()
}

func main() {
	do1()
}
