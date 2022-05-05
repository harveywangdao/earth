package main

import (
	"fmt"
	"runtime"
)

func do1() {
	pc, file, line, ok := runtime.Caller(0)
	fn := runtime.FuncForPC(pc).Name()
	fmt.Println(fn, file, line, ok)
}

func main() {
	do1()
}
