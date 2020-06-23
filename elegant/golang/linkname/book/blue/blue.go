package blue

import (
	_ "unsafe"
)

//go:linkname getblue blue.getblue
func getblue(s string) string {
	return s + " blue"
}

//go:linkname getblue2 test/cache/book/green.GetGreen2
func getblue2(s string) string {
	return s + " blue2"
}
