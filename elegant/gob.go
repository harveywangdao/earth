package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type People struct {
	Name string
	Age  int
}

func main() {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	p := People{
		Name: "dfsd",
		Age:  14,
	}
	p2 := People{
		Name: "dfsd",
		Age:  15,
	}
	p3 := People{
		Name: "dfsd",
		Age:  16,
	}
	fmt.Println(buf.Bytes())
	enc.Encode(p)
	fmt.Println(buf.Bytes())
	enc.Encode(p2)
	fmt.Println(buf.Bytes())
	enc.Encode(p3)

	fmt.Println(buf.Bytes())

	p4 := People{}
	dec.Decode(&p4)
	fmt.Printf("%+v\n", p4)

	p5 := People{}
	dec.Decode(&p5)
	fmt.Printf("%+v\n", p5)

	p6 := People{}
	dec.Decode(&p6)
	fmt.Printf("%+v\n", p6)

	p7 := People{}
	dec.Decode(&p7)
	fmt.Printf("%+v\n", p7)
}
