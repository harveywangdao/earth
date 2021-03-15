package main

import (
	"fmt"
	"math"
	"strconv"
)

func stack1() {
	s := NewArrayStack(10)
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

type listNode struct {
	Val  interface{}
	Next *listNode
}

type MyStack struct {
	head *listNode
	sz   int
}

func (this *MyStack) Push(x interface{}) {
	n := &listNode{
		Val:  x,
		Next: this.head,
	}
	this.head = n
	this.sz++
}

func (this *MyStack) Pop() interface{} {
	if this.head == nil {
		return nil
	}
	x := this.head.Val
	this.head = this.head.Next
	this.sz--
	return x
}

func (this *MyStack) Peek() interface{} {
	if this.head == nil {
		return nil
	}
	return this.head.Val
}

func (this *MyStack) Size() int {
	return this.sz
}

func (this *MyStack) IsEmpty() bool {
	if this.head == nil {
		return true
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyListStack struct {
	head *ListNode
	sz   int
}

func (this *MyListStack) Push(x int) {
	n := &ListNode{
		Val:  x,
		Next: this.head,
	}
	this.head = n
	this.sz++
}

func (this *MyListStack) Pop() int {
	if this.head == nil {
		return -1
	}
	x := this.head.Val
	this.head = this.head.Next
	this.sz--
	return x
}

func (this *MyListStack) Peek() int {
	if this.head == nil {
		return -1
	}
	return this.head.Val
}

func (this *MyListStack) Len() int {
	return this.sz
}

func (this *MyListStack) IsEmpty() bool {
	if this.head == nil {
		return true
	}
	return false
}

type MyArrayStack struct {
	arr    []int
	total  int
	length int
}

func NewArrayStack(n int) *MyArrayStack {
	return &MyArrayStack{
		arr:   make([]int, n),
		total: n,
	}
}

func (m *MyArrayStack) Push(val int) bool {
	if m.length >= m.total {
		return false
	}
	m.arr[m.length] = val
	m.length++
	return true
}

func (m *MyArrayStack) Pop() int {
	if m.length <= 0 {
		return -1
	}
	m.length--
	return m.arr[m.length]
}

func (m *MyArrayStack) Peek() int {
	if m.length <= 0 {
		return -1
	}
	return m.arr[m.length-1]
}

func (m *MyArrayStack) Len() int {
	return m.length
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

func evalRPN(tokens []string) int {
	stack := &MyStack{}

	for i := 0; i < len(tokens); i++ {
		var num int
		switch tokens[i] {
		case "+":
			a := stack.Pop().(int)
			b := stack.Pop().(int)
			num = a + b
		case "-":
			a := stack.Pop().(int)
			b := stack.Pop().(int)
			num = b - a
		case "*":
			a := stack.Pop().(int)
			b := stack.Pop().(int)
			num = a * b
		case "/":
			a := stack.Pop().(int)
			b := stack.Pop().(int)
			num = b / a
		default:
			num, _ = strconv.Atoi(tokens[i])
		}

		stack.Push(num)
	}
	return stack.Pop().(int)
}

func main() {
	stack1()
}
