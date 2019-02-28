package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func say1(s string) {
	for i := 0; i < 200000000; i++ {
		fmt.Println(s)
	}
}

func say2(s string) {
	runtime.Gosched()
	fmt.Println(s)
	os.Exit(1)
}

func main() {
	go say1("hello")
	say2("world")
	time.Sleep(10 * time.Second)
}
