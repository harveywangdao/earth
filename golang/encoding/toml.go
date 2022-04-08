package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type People struct {
	Name string
	Age  int
}

func do1() {
	p := People{}
	data := []byte(`
Name = "xiaoming"
Age = 12
Name = "xiaoming2"
    `)
	err := toml.Unmarshal(data, &p)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}

func main() {
	do1()
}
