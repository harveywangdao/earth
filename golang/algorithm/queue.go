package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func ring1() {
	r := ring.New(5)
	for i := 0; i < 5; i++ {
		r.Value = i
		r = r.Next()
	}
	r.Do(func(i interface{}) {
		fmt.Print(i, " ")
	})
	fmt.Println()

	r = r.Move(2)
	r.Do(func(i interface{}) {
		fmt.Print(i, " ")
	})
	fmt.Println()
}

func josephus1(n, m int) []int {
	r := ring.New(n)
	for i := 0; i < n; i++ {
		r.Value = i + 1
		r = r.Next()
	}

	var res []int
	r = r.Prev()
	for r != r.Next() {
		r = r.Move(m - 1)
		res = append(res, r.Unlink(1).Value.(int))
	}

	res = append(res, r.Value.(int))
	return res
}

func josephus(n, m int) []int {
	jo := make([]int, n)

	i, c := 0, 0
	var res []int
	for len(res) < n {
		if jo[i] == 0 {
			c++
			if c == m {
				jo[i] = -1
				c = 0
				res = append(res, i+1)
			}
		}

		i++
		if i == n {
			i = 0
		}
	}

	return res
}

func lastRemaining1(n int, m int) int {
	jo := make([]int, n)

	i, c := 0, 0
	var res, x int
	for x < n {
		if jo[i] == 0 {
			c++
			if c == m {
				jo[i] = -1
				c = 0
				res = i
				x++
			}
		}

		i++
		if i == n {
			i = 0
		}
	}

	return res
}

func _lastRemaining(n, m int) int {
	if n == 1 {
		return 0
	}
	x := _lastRemaining(n-1, m)
	return (m + x) % n
}

func lastRemaining2(n, m int) int {
	return _lastRemaining(n, m)
}

func lastRemaining(n, m int) int {
	x := 0
	for i := 2; i <= n; i++ {
		x = (m + x) % i
	}
	return x
}

func printg(grid [][]byte) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%c ", grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func numIslands1(grid [][]byte) int {
	c := 0
	ch := make(chan Coordinate, 100)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				continue
			}

			ch <- Coordinate{i, j}
			_numIslands1(grid, ch)
			c++
		}
	}
	return c
}

type Coordinate struct {
	i, j int
}

func _numIslands1(grid [][]byte, ch chan Coordinate) {
	var co Coordinate
	select {
	case co = <-ch:
	default:
		return
	}

	i, j := co.i, co.j
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		_numIslands1(grid, ch)
		return
	}
	if grid[i][j] == '0' {
		_numIslands1(grid, ch)
		return
	}

	if i+1 < len(grid) {
		ch <- Coordinate{i + 1, j}
	}
	if j+1 < len(grid[i]) {
		ch <- Coordinate{i, j + 1}
	}
	if i-1 >= 0 {
		ch <- Coordinate{i - 1, j}
	}
	if j-1 >= 0 {
		ch <- Coordinate{i, j - 1}
	}

	grid[i][j] = '0'
	_numIslands1(grid, ch)
}

func numIslands2(grid [][]byte) int {
	c := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				continue
			}
			_numIslands2(grid, i, j)
			c++
		}
	}
	return c
}

func _numIslands2(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	if j-1 >= 0 {
		_numIslands2(grid, i, j-1)
	}
	if j+1 < len(grid[i]) {
		_numIslands2(grid, i, j+1)
	}
	if i-1 >= 0 {
		_numIslands2(grid, i-1, j)
	}
	if i+1 < len(grid) {
		_numIslands2(grid, i+1, j)
	}
}

func numIslands3(grid [][]byte) int {
	n := len(grid)
	count := 0
	for x := 0; x < n; x++ {
		m := len(grid[x])
		for y := 0; y < m; y++ {
			if grid[x][y] == '0' {
				continue
			}
			count++

			queue := list.New()
			pos := m*x + y
			queue.PushFront(pos)

			for queue.Len() > 0 {
				e := queue.Back()
				queue.Remove(e)
				pos := e.Value.(int)
				i := pos / m
				j := pos % m

				if grid[i][j] == '0' {
					continue
				}

				grid[i][j] = '0'
				if j-1 >= 0 && grid[i][j-1] == '1' {
					co := m*i + j - 1
					queue.PushFront(co)
				}
				if j+1 < m && grid[i][j+1] == '1' {
					co := m*i + j + 1
					queue.PushFront(co)
				}
				if i-1 >= 0 && grid[i-1][j] == '1' {
					co := m*(i-1) + j
					queue.PushFront(co)
				}
				if i+1 < n && grid[i+1][j] == '1' {
					co := m*(i+1) + j
					queue.PushFront(co)
				}
			}
		}
	}
	return count
}

type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func NewUnionFind(grid [][]byte) *UnionFind {
	u := &UnionFind{}
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				u.parent = append(u.parent, i*n+j)
				u.count++
			} else {
				u.parent = append(u.parent, -1)
			}
			u.rank = append(u.rank, 0)
		}
	}
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
		u.count--
	}
}

func (u *UnionFind) getCount() int {
	fmt.Println(u.parent)
	return u.count
}

func numIslands(grid [][]byte) int {
	nr := len(grid)
	nc := len(grid[0])

	uf := NewUnionFind(grid)

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				grid[r][c] = '0'
				if r-1 >= 0 && grid[r-1][c] == '1' {
					uf.unite(r*nc+c, (r-1)*nc+c)
				}
				if r+1 < nr && grid[r+1][c] == '1' {
					uf.unite(r*nc+c, (r+1)*nc+c)
				}
				if c-1 >= 0 && grid[r][c-1] == '1' {
					uf.unite(r*nc+c, r*nc+c-1)
				}
				if c+1 < nc && grid[r][c+1] == '1' {
					uf.unite(r*nc+c, r*nc+c+1)
				}
			}
		}
	}
	return uf.getCount()
}

type MyArrayQueue struct {
	arr    []int
	arrlen int
	start  int
	length int
}

func NewMyArrayQueue(k int) MyArrayQueue {
	return MyArrayQueue{
		arr:    make([]int, k),
		arrlen: k,
		start:  0,
		length: 0,
	}
}

func (this *MyArrayQueue) EnQueue(value int) bool {
	if this.length >= this.arrlen {
		return false
	}

	this.arr[(this.start+this.length)%this.arrlen] = value
	this.length++
	return true
}

func (this *MyArrayQueue) DeQueue() bool {
	if this.length <= 0 {
		return false
	}
	this.start++
	this.length--
	if this.start >= this.arrlen {
		this.start = 0
	}
	return true
}

func (this *MyArrayQueue) Front() int {
	if this.length <= 0 {
		return -1
	}
	return this.arr[this.start]
}

func (this *MyArrayQueue) Rear() int {
	if this.length <= 0 {
		return -1
	}

	return this.arr[(this.start+this.length-1)%this.arrlen]
}

func (this *MyArrayQueue) IsEmpty() bool {
	if this.length <= 0 {
		return true
	}
	return false
}

func (this *MyArrayQueue) IsFull() bool {
	if this.length >= this.arrlen {
		return true
	}
	return false
}

type MyQueue struct {
	li *list.List
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		li: list.New(),
	}
}

func (this *MyQueue) Push(v interface{}) {
	this.li.PushFront(v)
}

func (this *MyQueue) Pop() interface{} {
	e := this.li.Back()
	this.li.Remove(e)
	return e.Value
}

func (this *MyQueue) Peek() interface{} {
	e := this.li.Back()
	return e.Value
}

func (this *MyQueue) Size() int {
	return this.li.Len()
}

func (this *MyQueue) IsEmpty() bool {
	return this.li.Len() == 0
}

type MyQueueByStack struct {
	s1 *MyListStack
	s2 *MyListStack
}

func NewMyQueueByStack() MyQueueByStack {
	return MyQueueByStack{
		s1: &MyListStack{},
		s2: &MyListStack{},
	}
}

func (this *MyQueueByStack) Push(x int) {
	this.s1.Push(x)
}

func (this *MyQueueByStack) Pop() int {
	if !this.s2.IsEmpty() {
		return this.s2.Pop()
	}

	for this.s1.Size() > 1 {
		this.s2.Push(this.s1.Pop())
	}
	return this.s1.Pop()
}

func (this *MyQueueByStack) Peek() int {
	if !this.s2.IsEmpty() {
		return this.s2.Peek()
	}

	for this.s1.Size() > 0 {
		this.s2.Push(this.s1.Pop())
	}

	return this.s2.Peek()
}

func (this *MyQueueByStack) Empty() bool {
	return this.s1.IsEmpty() && this.s2.IsEmpty()
}

type MyStackByQueue struct {
	q1 *MyListQueue
	q2 *MyListQueue
}

func NewMyStackByQueue() *MyStackByQueue {
	return &MyStackByQueue{
		q1: NewMyListQueue(),
		q2: NewMyListQueue(),
	}
}

func (this *MyStackByQueue) Push(x int) {
	if this.q1.IsEmpty() {
		this.q2.Push(x)
	} else {
		this.q1.Push(x)
	}
}

func (this *MyStackByQueue) Pop() int {
	if this.q1.IsEmpty() && this.q2.IsEmpty() {
		return -1
	}

	if !this.q1.IsEmpty() {
		for this.q1.Size() > 1 {
			v := this.q1.Pop()
			this.q2.Push(v)
		}
		return this.q1.Pop().(int)
	} else {
		for this.q2.Size() > 1 {
			v := this.q2.Pop()
			this.q1.Push(v)
		}
		return this.q2.Pop().(int)
	}
}

func (this *MyStackByQueue) Top() int {
	if this.q1.IsEmpty() && this.q2.IsEmpty() {
		return -1
	}

	if !this.q1.IsEmpty() {
		for this.q1.Size() > 1 {
			v := this.q1.Pop()
			this.q2.Push(v)
		}
		v := this.q1.Pop()
		this.q2.Push(v)

		return v.(int)
	} else {
		for this.q2.Size() > 1 {
			v := this.q2.Pop()
			this.q1.Push(v)
		}
		v := this.q2.Pop()
		this.q1.Push(v)
		return v.(int)
	}
}

func (this *MyStackByQueue) Empty() bool {
	return this.q1.IsEmpty() && this.q2.IsEmpty()
}

type MyStackByQueue2 struct {
	q1 *MyListQueue
	q2 *MyListQueue
}

func NewMyStackByQueue2() *MyStackByQueue2 {
	return &MyStackByQueue2{
		q1: NewMyListQueue(),
		q2: NewMyListQueue(),
	}
}

func (this *MyStackByQueue2) Push(x int) {
	this.q2.Push(x)

	for this.q1.Size() > 0 {
		v := this.q1.Pop()
		this.q2.Push(v)
	}
	this.q1, this.q2 = this.q2, this.q1
}

func (this *MyStackByQueue2) Pop() int {
	return this.q1.Pop().(int)
}

func (this *MyStackByQueue2) Top() int {
	return this.q1.Peek().(int)
}

func (this *MyStackByQueue2) Empty() bool {
	return this.q1.IsEmpty()
}

type MyStackByQueue3 struct {
	q *MyListQueue
}

func NewMyStackByQueue3() *MyStackByQueue3 {
	return &MyStackByQueue3{
		q: NewMyListQueue(),
	}
}

func (this *MyStackByQueue3) Push(x int) {

	this.q.Push(x)

	for sz := 0; sz < this.q.Size()-1; sz++ {
		v := this.q.Pop()
		this.q.Push(v)
	}
}

func (this *MyStackByQueue3) Pop() int {
	return this.q.Pop().(int)
}

func (this *MyStackByQueue3) Top() int {
	return this.q.Peek().(int)
}

func (this *MyStackByQueue3) Empty() bool {
	return this.q.IsEmpty()
}

func main() {
	/*grid := [][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}*/
	/*grid := [][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}*/
	/*grid := [][]byte{
		[]byte{'1', '1', '1'},
		[]byte{'0', '1', '0'},
		[]byte{'1', '1', '1'},
	}*/
	/*grid := [][]byte{
		[]byte{'1', '1', '1', '1', '1', '0', '1', '1', '1', '1'},
		[]byte{'1', '0', '1', '0', '1', '1', '1', '1', '1', '1'},
		[]byte{'0', '1', '1', '1', '0', '1', '1', '1', '1', '1'},
		[]byte{'1', '1', '0', '1', '1', '0', '0', '0', '0', '1'},
		[]byte{'1', '0', '1', '0', '1', '0', '0', '1', '0', '1'},
		[]byte{'1', '0', '0', '1', '1', '1', '0', '1', '0', '0'},
		[]byte{'0', '0', '1', '0', '0', '1', '1', '1', '1', '0'},
		[]byte{'1', '0', '1', '1', '1', '0', '0', '1', '1', '1'},
		[]byte{'1', '1', '1', '1', '1', '1', '1', '1', '0', '1'},
		[]byte{'1', '0', '1', '1', '1', '1', '1', '1', '1', '0'},
	}*/
	//fmt.Println(numIslands(grid))
	fmt.Println(lastRemaining(6, 3))
}
