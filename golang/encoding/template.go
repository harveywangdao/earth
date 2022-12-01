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

var funcs = template.FuncMap{"toInt64": toInt64}

func toInt64(s string) int64 {
	return 13
}

func do3() {
	sweaters := Inventory{"13", 16}
	tmpl, err := template.New("test").Funcs(funcs).Parse(`"{{toInt64 .Material}}"`)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

func do4() {
	args := map[string]interface{}{
		"TitleType": 2,
	}
	tmpl, err := template.New("test").Parse("{{if eq .TitleType 2}}@you in the comments{{else if eq .TitleType 3}}mention you in the documentation{{end}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, args)
	if err != nil {
		panic(err)
	}
}

func main() {
	do4()
}
