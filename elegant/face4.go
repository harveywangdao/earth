package main

import (
	"fmt"
)

type Human interface {
	Speak(string) string
}

type People interface {
	Speak(string) string
	Eat()
}

type Stduent struct{}

func (s *Stduent) Speak(think string) string {
	fmt.Println(think)
	return think
}

func (s *Stduent) Eat() {
	fmt.Println("eat")
}

func main() {
	var pe People
	var hu Human
	pe = &Stduent{}
	hu = pe
	//pe = hu //error!
	_ = pe
	_ = hu

	hu.Speak("adcz")
	pe.Eat()
}
