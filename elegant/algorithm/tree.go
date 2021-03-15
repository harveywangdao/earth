package main

import (
	"container/list"
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return _buildTree(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func _buildTree1(preorder []int, start1, end1 int, inorder []int, start2, end2 int) *TreeNode {
	root := &TreeNode{Val: preorder[start1]}
	if start1 == end1 {
		return root
	}

	i := start2
	for ; i <= end2; i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	if i == start2 {
		root.Right = _buildTree(preorder, start1+1, end1, inorder, i+1, end2)
	} else if i == end2 {
		root.Left = _buildTree(preorder, start1+1, end1, inorder, start2, i-1)
	} else {
		j := start1 + 1

		for x := start2; x <= i-1; x++ {
			for y := start1 + 1; y <= end1; y++ {
				if preorder[y] == inorder[x] {
					if y > j {
						j = y
					}
					break
				}
			}
		}

		root.Left = _buildTree(preorder, start1+1, j, inorder, start2, i-1)
		root.Right = _buildTree(preorder, j+1, end1, inorder, i+1, end2)
	}

	return root
}

func _buildTree(preorder []int, start1, end1 int, inorder []int, start2, end2 int) *TreeNode {
	root := &TreeNode{Val: preorder[start1]}
	if start1 == end1 {
		return root
	}

	i := start2
	for ; i <= end2; i++ {
		if inorder[i] == root.Val {
			break
		}
	}

	if i == start2 {
		root.Right = _buildTree(preorder, start1+1, end1, inorder, i+1, end2)
	} else if i == end2 {
		root.Left = _buildTree(preorder, start1+1, end1, inorder, start2, i-1)
	} else {
		j := start1 + (i - start2)
		root.Left = _buildTree(preorder, start1+1, j, inorder, start2, i-1)
		root.Right = _buildTree(preorder, j+1, end1, inorder, i+1, end2)
	}

	return root
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	} else if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	} else if root.Left == nil && root.Right == nil {
		return 1
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)
	if l < r {
		return l + 1
	}
	return r + 1
}

func isValidBST1(root *TreeNode) bool {
	var val int
	fisrt := true
	var _isValidBST func(node *TreeNode) bool
	_isValidBST = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if node.Left != nil && !_isValidBST(node.Left) {
			return false
		}

		if !fisrt {
			if val >= node.Val {
				return false
			}
		} else {
			fisrt = false
		}
		val = node.Val

		if node.Right != nil && !_isValidBST(node.Right) {
			return false
		}
		return true
	}

	return _isValidBST(root)
}

func isValidBST(root *TreeNode) bool {
	return _isValidBST(root, math.MinInt64, math.MaxInt64)
}

func _isValidBST(root *TreeNode, low, high int) bool {
	if root == nil {
		return true
	}

	if root.Val <= low || root.Val >= high {
		return false
	}

	return _isValidBST(root.Left, low, root.Val) && _isValidBST(root.Right, root.Val, high)
}

func isSymmetric1(root *TreeNode) bool {
	return _isSymmetric1(root.Left, root.Right)
}

func _isSymmetric1(left, right *TreeNode) bool {
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	if left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
	}

	if left == nil && left == nil {
		return true
	}

	if !_isSymmetric1(left.Left, right.Right) {
		return false
	}
	return _isSymmetric1(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	lstack := &MyStack{}
	rstack := &MyStack{}

	lstack.Push(root.Left)
	rstack.Push(root.Right)

	var left, right *TreeNode
	for {
		left = lstack.Pop().(*TreeNode)
		right = rstack.Pop().(*TreeNode)

		if left == nil && right != nil {
			return false
		}
		if left != nil && right == nil {
			return false
		}

		if left != nil && right != nil {
			if left.Val != right.Val {
				return false
			}
			lstack.Push(left.Right)
			rstack.Push(right.Left)

			lstack.Push(left.Left)
			rstack.Push(right.Right)
			continue
		}

		if lstack.IsEmpty() && rstack.IsEmpty() {
			break
		}
	}
	return true
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushFront(root)

	for queue.Len() > 0 {
		var row []int
		for sz := queue.Len(); sz > 0; sz-- {
			e := queue.Back()
			queue.Remove(e)

			node := e.Value.(*TreeNode)
			row = append(row, node.Val)

			if node.Left != nil {
				queue.PushFront(node.Left)
			}
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		res = append(res, row)
	}
	return res
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	mid := n / 2
	node := &TreeNode{Val: nums[mid]}
	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])
	return node
}

func inorderTraversal(root *TreeNode) []int {

}

func main() {
	root := &TreeNode{Val: 5}
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 4}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 6}
	root.Left = n1
	root.Right = n2
	n2.Left = n3
	n2.Right = n4
	fmt.Println(isValidBST(root))
}
