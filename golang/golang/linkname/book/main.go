package main

import (
	"log"
	"test/cache/book/green"
)

func do1() {
	log.Println(green.Getgreen("dd"))
	log.Println(green.GetGreen2("vv"))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
