package main

import (
	"fmt"
)

func solve1(board [][]byte) {
	nrow := len(board)
	if nrow < 3 {
		return
	}
	ncol := len(board[0])
	if ncol < 3 {
		return
	}

	// up
	for i := 0; i < ncol; i++ {
		if board[0][i] != 'O' {
			continue
		}
		_solve1(board, 0, i)
	}

	// down
	for i := 0; i < ncol; i++ {
		if board[nrow-1][i] != 'O' {
			continue
		}
		_solve1(board, nrow-1, i)
	}

	// left
	for i := 1; i < nrow-1; i++ {
		if board[i][0] != 'O' {
			continue
		}
		_solve1(board, i, 0)
	}

	// right
	for i := 1; i < nrow-1; i++ {
		if board[i][ncol-1] != 'O' {
			continue
		}
		_solve1(board, i, ncol-1)
	}

	for i := 0; i < nrow; i++ {
		for j := 0; j < ncol; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
}

func _solve1(board [][]byte, i, j int) {
	if board[i][j] != 'O' {
		return
	}
	board[i][j] = 'A'
	if j-1 >= 0 {
		_solve1(board, i, j-1)
	}
	if j+1 < len(board[i]) {
		_solve1(board, i, j+1)
	}
	if i-1 >= 0 {
		_solve1(board, i-1, j)
	}
	if i+1 < len(board) {
		_solve1(board, i+1, j)
	}
}

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(grid [][]byte) *UnionFind {
	u := &UnionFind{}
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'O' {
				u.parent = append(u.parent, i*n+j)
			} else {
				u.parent = append(u.parent, -1)
			}
			u.rank = append(u.rank, 0)
		}
	}
	u.parent = append(u.parent, n*m)
	u.rank = append(u.rank, 0)
	return u
}

func (u *UnionFind) find(i int) int {
	if u.parent[i] != i {
		u.parent[i] = u.find(u.parent[i])
	}
	return u.parent[i]
}

func (u *UnionFind) unite(x, y int) {
	rootx := u.find(x)
	rooty := u.find(y)
	if rootx != rooty {
		if u.rank[rootx] < u.rank[rooty] {
			rootx, rooty = rooty, rootx
		}
		u.parent[rooty] = rootx
		if u.rank[rootx] == u.rank[rooty] {
			u.rank[rootx] += 1
		}
	}
}

func (u *UnionFind) isunion(x, y int) bool {
	rootx := u.find(x)
	rooty := u.find(y)
	return rootx == rooty
}

func solve(board [][]byte) {
	nrow := len(board)
	if nrow < 3 {
		return
	}
	ncol := len(board[0])
	if ncol < 3 {
		return
	}

	uf := NewUnionFind(board)

	for i := 0; i < nrow; i++ {
		for j := 0; j < ncol; j++ {
			if board[i][j] == 'O' {
				if i == 0 || i == nrow-1 || j == 0 || j == ncol-1 {
					uf.unite(ncol*i+j, nrow*ncol)
				} else {
					if j-1 >= 0 && board[i][j-1] == 'O' {
						uf.unite(ncol*i+j, ncol*i+j-1)
					}
					if j+1 < len(board[i]) && board[i][j+1] == 'O' {
						uf.unite(ncol*i+j, ncol*i+j+1)
					}
					if i-1 >= 0 && board[i-1][j] == 'O' {
						uf.unite(ncol*i+j, ncol*(i-1)+j)
					}
					if i+1 < len(board) && board[i+1][j] == 'O' {
						uf.unite(ncol*i+j, ncol*(i+1)+j)
					}
				}
			}
		}
	}

	for i := 0; i < nrow; i++ {
		for j := 0; j < ncol; j++ {
			if board[i][j] == 'O' && uf.isunion(ncol*i+j, nrow*ncol) {
				continue
			}
			board[i][j] = 'X'
		}
	}
}

func main() {
	board := [][]byte{
		[]byte{'X', 'X', 'X', 'X'},
		[]byte{'X', 'O', 'O', 'X'},
		[]byte{'X', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'X'},
	}
	solve(board)
	fmt.Println(board)
}
