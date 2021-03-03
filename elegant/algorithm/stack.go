package main

import (
	"fmt"
	"math"
)

type MyStack struct {
	arr    []int
	total  int
	length int
}

func NewStack(n int) *MyStack {
	return &MyStack{
		arr:   make([]int, n),
		total: n,
	}
}

func (m *MyStack) Push(val int) bool {
	if m.length >= m.total {
		return false
	}
	m.arr[m.length] = val
	m.length++
	return true
}

func (m *MyStack) Pop() (int, bool) {
	if m.length <= 0 {
		return 0, false
	}
	m.length--
	return m.arr[m.length], true
}

func (m *MyStack) Len() int {
	return m.length
}

func stack1() {
	s := NewStack(10)
	for i := 0; i < 11; i++ {
		ok := s.Push(i)
		if !ok {
			fmt.Println(i, "push fail")
		}
	}

	for i := 0; i < 11; i++ {
		v, ok := s.Pop()
		fmt.Println(v, ok)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinStack struct {
	head    *ListNode
	minhead *ListNode
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	n := &ListNode{
		Val:  x,
		Next: this.head,
	}
	this.head = n

	y := x
	if this.minhead != nil && y > this.minhead.Val {
		y = this.minhead.Val
	}
	n2 := &ListNode{
		Val:  y,
		Next: this.minhead,
	}
	this.minhead = n2
}

func (this *MinStack) Pop() {
	if this.head == nil {
		return
	}
	this.head = this.head.Next
	this.minhead = this.minhead.Next
}

func (this *MinStack) Top() int {
	if this.head == nil {
		return -1
	}
	return this.head.Val
}

func (this *MinStack) GetMin() int {
	if this.head == nil {
		return -1
	}
	return this.minhead.Val
}

func (this *MinStack) getMin() int {
	if this.head == nil {
		return math.MaxInt64
	}
	min := this.head.Val
	p := this.head.Next
	for p != nil {
		if p.Val < min {
			min = p.Val
		}
		p = p.Next
	}
	return min
}

func main() {
	stack1()
}
