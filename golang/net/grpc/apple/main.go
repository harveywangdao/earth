package main

import (
	proto2 "github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/proto"
	"log"
)

func do1() {
	h := &HelloRequest{
		Name: "xiaoming",
	}
	log.Println([]byte(h.Name))

	data, err := proto.Marshal(h)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)

	data2, err := proto2.Marshal(h)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data2)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	do1()
}
