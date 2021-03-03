package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		line := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				line[j] = 1
			} else {
				line[j] = ret[i-1][j-1] + ret[i-1][j]
			}
		}
		ret[i] = line
	}
	return ret
}

func getRow(rowIndex int) []int {
	ret := make([]int, rowIndex+1)
	var last int

	for i := 0; i <= rowIndex; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				ret[j], last = 1, 1
			} else {
				ret[j], last = last+ret[j], ret[j]
			}
		}
	}

	return ret
}

func getRow2(rowIndex int) []int {
	ret := make([]int, rowIndex+1)
	ret[0] = 1
	for i := 1; i <= rowIndex; i++ {
		ret[i] = (rowIndex - i + 1) * ret[i-1] / i
	}

	return ret
}

func main() {
	fmt.Println(generate(5))
	return

	a := 0xd   //1101
	b := 0xb   //1011
	c := a ^ b //110
	d := c ^ a
	e := c ^ b
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", b)
	fmt.Printf("%b\n", c)
	fmt.Printf("%b\n", d)
	fmt.Printf("%b\n", e)
}
