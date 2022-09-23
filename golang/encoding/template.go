package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func do1() {
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

func do2() {
	sweaters := Inventory{"wool", 16}
	tmpl, err := template.New("test").Parse("{{if eq .Count 17}}T1{{else if eq .Count 16}}T0{{end}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

func main() {
	do2()
}
