package green

import (
	_ "test/cache/book/blue"
	_ "unsafe"
)

//go:linkname getBlue blue.getblue
func getBlue(s string) string

func Getgreen(s string) string {
	return getBlue(s + " green")
}

func GetGreen2(s string) string
