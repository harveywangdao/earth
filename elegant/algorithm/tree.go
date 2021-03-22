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

func inorderTraversal3(root *TreeNode) []int {
	var res []int
	st := &MyStack{}
	cur := root
	for cur != nil {
		if cur.Right != nil {
			st.Push(cur.Right)
		}

		if cur.Left != nil {
			st.Push(&TreeNode{Val: cur.Val})
			cur = cur.Left
		} else {
			res = append(res, cur.Val)
			if st.Size() == 0 {
				break
			}
			cur = st.Pop().(*TreeNode)
		}
	}

	return res
}

// 中序递归
func inorderTraversal(root *TreeNode) []int {
	var res []int
	_inorderTraversal(root, &res)
	return res
}

func _inorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	_inorderTraversal(root.Left, res)
	*res = append(*res, root.Val)
	_inorderTraversal(root.Right, res)
}

// 中序迭代
func inorderTraversal1(root *TreeNode) []int {
	var res []int
	st := &MyStack{}
	cur := root
	for cur != nil || st.Size() > 0 {
		for cur != nil {
			st.Push(cur)
			cur = cur.Left
		}

		cur = st.Pop().(*TreeNode)
		res = append(res, cur.Val)
		cur = cur.Right
	}

	return res
}

// Morris中序遍历
func inorderTraversal2(root *TreeNode) []int {
	var res []int

	for root != nil {
		if root.Left != nil {
			p := root.Left
			for p.Right != nil && p.Right != root {
				p = p.Right
			}

			if p.Right == root {
				res = append(res, root.Val)
				p.Right = nil
				root = root.Right
			} else {
				p.Right = root
				root = root.Left
			}
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}

	return res
}

// 前序递归
func preorderTraversal(root *TreeNode) []int {
	var res []int
	_preorderTraversal(root, &res)
	return res
}

func _preorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	_preorderTraversal(root.Left, res)
	_preorderTraversal(root.Right, res)
}

// 前序迭代
func preorderTraversal1(root *TreeNode) []int {
	var res []int

	st := &MyStack{}

	for root != nil {
		if root.Right != nil {
			st.Push(root.Right)
		}

		res = append(res, root.Val)

		if root.Left != nil {
			root = root.Left
		} else {
			if st.Size() == 0 {
				break
			}
			root = st.Pop().(*TreeNode)
		}
	}

	return res
}

// Morris前序遍历
func preorderTraversal2(root *TreeNode) []int {
	var res []int

	for root != nil {
		if root.Left != nil {
			p := root.Left
			for p.Right != nil && p.Right != root {
				p = p.Right
			}

			if p.Right == nil {
				res = append(res, root.Val)
				p.Right = root
				root = root.Left
			} else {
				p.Right = nil
				root = root.Right
			}
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}

	return res
}

// 后序递归
func postorderTraversal(root *TreeNode) []int {
	var res []int
	_postorderTraversal(root, &res)
	return res
}

func _postorderTraversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	_postorderTraversal(root.Left, res)
	_postorderTraversal(root.Right, res)
	*res = append(*res, root.Val)
}

// 后序迭代
func postorderTraversal1(root *TreeNode) []int {
	var res []int
	st := &MyStack{}

	cur := root
	var last *TreeNode
	for cur != nil || st.Size() > 0 {
		for cur != nil {
			st.Push(cur)
			cur = cur.Left
		}

		cur = st.Pop().(*TreeNode)
		if cur.Right == nil || cur.Right == last {
			res = append(res, cur.Val)
			last = cur
			cur = nil
		} else {
			st.Push(cur)
			cur = cur.Right
		}
	}

	return res
}

// Morris后序遍历
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func postorderTraversal2(root *TreeNode) (res []int) {
	addPath := func(node *TreeNode) {
		resSize := len(res)
		for ; node != nil; node = node.Right {
			res = append(res, node.Val)
		}
		reverse(res[resSize:])
	}

	p1 := root
	for p1 != nil {
		if p2 := p1.Left; p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
			addPath(p1.Left)
		}
		p1 = p1.Right
	}
	addPath(root)
	return
}

func main() {

}
