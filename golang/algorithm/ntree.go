package main

import (
	"fmt"
)

type Node struct {
	Val      int
	Children []*Node
}

func preorder1(root *Node) []int {
	var res []int
	_preorder1(root, &res)
	return res
}

func _preorder1(root *Node, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	for _, node := range root.Children {
		_preorder1(node, res)
	}
}

func preorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := &MyStack{}
	stack.Push(root)

	for stack.Size() > 0 {
		n := stack.Pop().(*Node)
		res = append(res, n.Val)

		for i := len(n.Children) - 1; i >= 0; i-- {
			stack.Push(n.Children[i])
		}
	}

	return res
}

func postorder1(root *Node) []int {
	var res []int
	_postorder1(root, &res)
	return res
}

func _postorder1(root *Node, res *[]int) {
	if root == nil {
		return
	}
	for _, n := range root.Children {
		_postorder1(n, res)
	}
	*res = append(*res, root.Val)
}

func postorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := &MyStack{}
	var last *Node

	for root != nil || stack.Size() > 0 {
		for root != nil {
			stack.Push(root)
			if len(root.Children) > 0 {
				root = root.Children[0]
			} else {
				root = nil
			}
		}

		root = stack.Pop().(*Node)
		i := 0
		for ; i < len(root.Children); i++ {
			if root.Children[i] == last {
				i++
				break
			}
		}

		if i >= len(root.Children) {
			res = append(res, root.Val)
			last = root
			root = nil
		} else {
			stack.Push(root)
			root = root.Children[i]
		}
	}

	return res
}

func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := NewMyQueue()
	queue.Push(root)

	for queue.Size() > 0 {
		var row []int
		n := queue.Size()
		for i := 0; i < n; i++ {
			node := queue.Pop().(*Node)
			row = append(row, node.Val)
			for _, child := range node.Children {
				queue.Push(child)
			}
		}
		res = append(res, row)
	}
	return res
}

func main() {

}
