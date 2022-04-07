package main

import (
	"fmt"
	"github.com/mitchellh/hashstructure"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Extra  map[string]string
}

func main() {
	input := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"emails": []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mitchellh",
		},
	}

	fmt.Println(input)

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", result)
	hash()
}

func hash() {
	p1 := &Person{
		Name: "xiaoming",
		Age:  12,
	}

	hash, err := hashstructure.Hash(p1, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", hash)

	p2 := &Person{
		Name: "xiaoming",
		Age:  11,
	}

	hash, err = hashstructure.Hash(p2, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d\n", hash)
}
