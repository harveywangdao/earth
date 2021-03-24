package main

import (
	"fmt"
)

type matrix [2][2]int

func mul(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	l, r := 0, m*n-1
	for l <= r {
		mid := (l + r) / 2
		i, j := mid/n, mid%n
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func findDiagonalOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])

	res := make([]int, 0, m*n)
	down := false

	for c := 0; c <= m+n-2; c++ {
		if down {
			var i, j int
			if c >= n {
				i = c - n + 1
			}
			j = c - i

			for ; i < m && j >= 0; i, j = i+1, j-1 {
				res = append(res, matrix[i][j])
			}
		} else {
			var i, j int
			if c >= m {
				j = c - m + 1
			}
			i = c - j

			for ; i >= 0 && j < n; i, j = i-1, j+1 {
				res = append(res, matrix[i][j])
			}
		}

		down = !down
	}
	return res
}

func main() {
	A := matrix{{1, 2}, {3, 4}}
	B := matrix{{5, 6}, {7, 8}}
	fmt.Println(mul(A, B))
}
