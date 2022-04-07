package main

type Life int

//go:generate stringer -type=Life
const (
	Human Life = iota
	Dog
	Cat
)
