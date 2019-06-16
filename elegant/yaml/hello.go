package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type DDD struct {
	E int
	F string
}

type AAA struct {
	A int
	B string
	C []string
}

type BBB struct {
	C   int
	D   string
	Ddd DDD
}

type CCC struct {
	Aaa AAA
	Bbb []BBB
}

func marshal() {
	c := CCC{}
	c.Aaa.A = 34
	c.Aaa.B = "jsdf"
	c.Aaa.C = []string{"sdf", "fsdf", "sdfsd"}

	for i := 0; i < 5; i++ {
		b := BBB{}
		b.C = i + 34
		b.D = "xdcz"
		b.Ddd.E = 10 + i
		b.Ddd.F = "sdf"
		c.Bbb = append(c.Bbb, b)
	}

	data, err := yaml.Marshal(&c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("c =", c)
	fmt.Println(string(data))

	ioutil.WriteFile("my.yaml", data, 0666)
}

func unmarshal() {
	data, err := ioutil.ReadFile("my.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	c := CCC{}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("c =", c)
}

func main() {
	marshal()
	unmarshal()
}
