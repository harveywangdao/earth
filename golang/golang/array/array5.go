package main

func do1() {
	arr1 := [...]int{3, 4, 5}
	s1 := make([]int, 3)
	//copy(s1, arr1)
	//copy(arr1, s1)
	_, _ = arr1, s1
}

func do2() {
	arr1 := [...]int{3, 4, 5}
	s1 := arr1[:]
	s2 := arr1[0:]
	s3 := arr1[:len(arr1)]
	s4 := arr1[0:2:2]
	_, _, _, _ = s1, s2, s3, s4
}

func main() {
	do2()
}
