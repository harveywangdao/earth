package main

import (
	"fmt"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	stack := &MyStack{}
	pass := make(map[int]*Node)

	stack.Push(node)
	newNode := &Node{
		Val: node.Val,
	}
	pass[node.Val] = newNode

	for stack.Size() > 0 {
		old := stack.Pop().(*Node)
		nw := pass[old.Val]

		for _, nei := range old.Neighbors {
			nnei, ok := pass[nei.Val]
			if !ok {
				nnei = &Node{Val: nei.Val}
				pass[nei.Val] = nnei
				stack.Push(nei)
			}
			nw.Neighbors = append(nw.Neighbors, nnei)
		}
	}

	return newNode
}

func findTargetSumWays(nums []int, S int) int {
	count := 0
	_findTargetSumWays(nums, S, &count)
	return count
}

func _findTargetSumWays(nums []int, S int, count *int) {
	if len(nums) == 1 {
		if nums[0] == S {
			(*count)++
		}
		if -nums[0] == S {
			(*count)++
		}
		return
	}

	_findTargetSumWays(nums[1:], S-nums[0], count)
	_findTargetSumWays(nums[1:], S+nums[0], count)
}

func main() {

}
