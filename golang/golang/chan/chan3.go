package main

func do1() {
	c1 := make(chan int)
	c1 <- 13
	n1 := <-c1
	n2, ok := <-c1
	_ = n1
	_ = n2
	_ = ok
}

func do2() {
	c1 := make(chan int, 16)
	c1 <- 17
}

func do3() {
	var c1 chan int
	//c1 <- 1
	//<-c1
	close(c1)
}

func main() {
	do1()
}
