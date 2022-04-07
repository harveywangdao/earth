package main

import (
	"fmt"
)

func solveNQueens(n int) [][]string {
	res := [][]string{}
	used := make([]bool, n)
	path := make([]int, 0, n)
	_solveNQueens(n, path, used, &res)
	return res
}

func _solveNQueens(n int, path []int, used []bool, res *[][]string) {
	if len(path) == n {
		str := make([]byte, n)
		for i := 0; i < n; i++ {
			str[i] = '.'
		}

		line := make([]string, n)
		for i := 0; i < n; i++ {
			str[path[i]] = 'Q'
			line[i] = string(str)
			str[path[i]] = '.'
		}
		*res = append(*res, line)
		return
	}

	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}

		ok := true
		for j := 0; j < len(path); j++ {
			x := len(path) - j
			y := i - path[j]

			if x == y || x == (-y) {
				ok = false
				break
			}
		}

		if !ok {
			continue
		}

		path = append(path, i)
		used[i] = true
		_solveNQueens(n, path, used, res)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func printQueen(queen [][]string) {
	for i := 0; i < len(queen); i++ {
		for j := 0; j < len(queen[i]); j++ {
			fmt.Println(queen[i][j])
		}
		fmt.Println()
	}
	fmt.Println("len queen:", len(queen))
}

func main() {
	printQueen(solveNQueens(8))
}
