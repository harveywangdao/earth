package main

func do1() {
	str1 := "juice"
	s1 := []byte(str1)
	_ = s1
}

func do2() {
	s1 := []byte{'j', 'u', 'i', 'c', 'e'}
	str1 := string(s1)
	_ = str1
}

func main() {
	do1()
}
