package main

import (
	"os"
	"runtime/trace"
)

func send(c chan<- int) {
	n := 1
	for {
		c <- n
		n++
	}
}

func read(c <-chan int) {
	for {
		m := <-c
		if m >= 100 {
			return
		}
	}
}

func do1() {
	c := make(chan int, 1000)

	go send(c)
	read(c)
}

func main() {
	file, err := os.Create("app.trace")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = trace.Start(file)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	do1()
}
