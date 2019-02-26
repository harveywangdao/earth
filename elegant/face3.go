package main

import (
	"fmt"
)

type student struct {
	Name string
}

func f(v interface{}) {
	switch msg := v.(type) {
	case *student, student:
		//msg.Name
		fmt.Println(msg)
	}
}

func main() {
	stu := &student{
		Name: "sdfgsd",
	}

	f(stu)
}
