package main

import (
	"container/list"
	"fmt"
)

func openLock1(deadends []string, target string) int {
	ts := []byte(target)
	steps := 0
	for i := 0; i < 4; i++ {
		n := int(ts[i] - '0')
		if n > 5 {
			n = 10 - n
		}
		steps += n
	}
	return steps
}

func openLock(deadends []string, target string) int {
	dm := make(map[string]int)
	for i := 0; i < len(deadends); i++ {
		dm[deadends[i]] = 1
	}
	if dm["0000"] == 1 {
		return -1
	}
	if target == "0000" {
		return 0
	}

	old := make(map[string]int)

	queue := list.New()
	old["0000"] = 1
	queue.PushFront("0000")

	step := 0
	for queue.Len() > 0 {
		for sz := queue.Len(); sz > 0; sz-- {
			e := queue.Back()
			queue.Remove(e)
			cur := []byte(e.Value.(string))

			for i := 0; i < 4; i++ {
				cur[i] = upone(cur[i])
				str := string(cur)
				if str == target {
					return step + 1
				}
				if dm[str] == 0 && old[str] == 0 {
					old[str] = 1
					queue.PushFront(str)
				}
				cur[i] = downone(cur[i])

				cur[i] = downone(cur[i])
				str = string(cur)
				if str == target {
					return step + 1
				}
				if dm[str] == 0 && old[str] == 0 {
					old[str] = 1
					queue.PushFront(str)
				}
				cur[i] = upone(cur[i])
			}
		}
		step++
	}
	return -1
}

func upone(c byte) byte {
	c++
	if c > '9' {
		return '0'
	}
	return c
}

func downone(c byte) byte {
	c--
	if c < '0' {
		return '9'
	}
	return c
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	n := len(image)
	if n == 0 {
		return image
	}
	m := len(image[0])

	old := image[sr][sc]
	if old == newColor {
		return image
	}

	image[sr][sc] = newColor
	queue := NewMyListQueue()
	queue.Push(sr*m + sc)

	for queue.Size() > 0 {
		x := queue.Pop().(int)
		r, c := x/m, x%m
		if r-1 >= 0 && image[r-1][c] == old {
			image[r-1][c] = newColor
			queue.Push((r-1)*m + c)
		}
		if r+1 < n && image[r+1][c] == old {
			image[r+1][c] = newColor
			queue.Push((r+1)*m + c)
		}
		if c-1 >= 0 && image[r][c-1] == old {
			image[r][c-1] = newColor
			queue.Push(r*m + c - 1)
		}
		if c+1 < m && image[r][c+1] == old {
			image[r][c+1] = newColor
			queue.Push(r*m + c + 1)
		}
	}

	return image
}

func updateMatrix1(matrix [][]int) [][]int {
	n := len(matrix)
	if n == 0 {
		return matrix
	}
	m := len(matrix[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] != 0 {
				matrix[i][j] = _updateMatrix1(matrix, i, j)
			}
		}
	}
	return matrix
}

func _updateMatrix1(matrix [][]int, r, c int) int {
	n := len(matrix)
	m := len(matrix[0])
	mod := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}

	queue := NewMyListQueue()
	queue.Push(r*m + c)
	visited := make(map[int]int)
	visited[r*m+c] = 1

	depth := 0
	for {
		depth++
		for i := queue.Size(); i > 0; i-- {
			x := queue.Pop().(int)
			r, c = x/m, x%m

			for j := 0; j < 4; j++ {
				a, b := r+mod[j][0], c+mod[j][1]
				if a >= 0 && b >= 0 && a < n && b < m {
					if matrix[a][b] == 0 {
						return depth
					}
					y := a*m + b
					if visited[y] == 0 {
						queue.Push(y)
						visited[y] = 1
					}
				}
			}
		}
	}
	return 0
}

func updateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	if n == 0 {
		return matrix
	}
	m := len(matrix[0])
	mod := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}

	queue := NewMyListQueue()
	visited := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				queue.Push(i*m + j)
				visited[i*m+j] = 1
			}
		}
	}

	depth := 0
	for queue.Size() > 0 {
		depth++
		for i := queue.Size(); i > 0; i-- {
			x := queue.Pop().(int)
			r, c := x/m, x%m

			for j := 0; j < 4; j++ {
				a, b := r+mod[j][0], c+mod[j][1]
				y := a*m + b

				if a >= 0 && b >= 0 && a < n && b < m && visited[y] == 0 {
					matrix[a][b] = depth
					queue.Push(y)
					visited[y] = 1
				}
			}
		}
	}

	return matrix
}

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	queue := NewMyListQueue()
	visited := make(map[int]int)
	visited[0] = 1
	queue.Push(0)

	for queue.Size() > 0 {
		x := queue.Pop().(int)

		for i := 0; i < len(rooms[x]); i++ {
			if visited[rooms[x][i]] == 0 {
				queue.Push(rooms[x][i])
				visited[rooms[x][i]] = 1
			}
		}
	}

	return len(visited) == n
}

func main() {
	fmt.Println(openLock(nil, "0202"))
}
