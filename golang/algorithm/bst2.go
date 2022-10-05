package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findmid(head *ListNode) (*ListNode, *ListNode) {
	p := head
	mid := head
	var prev *ListNode
	for p != nil && p.Next != nil {
		prev = mid
		mid = mid.Next
		p = p.Next.Next
	}
	return prev, mid
}

func sortedListToBST(head *ListNode) *TreeNode {
	prev, cur := findmid(head)
	if cur == nil {
		return nil
	}
	if prev != nil {
		prev.Next = nil
	} else {
		head = nil
	}

	head2 := cur.Next
	cur.Next = nil

	n := &TreeNode{Val: cur.Val}

	n.Left = sortedListToBST(head)
	n.Right = sortedListToBST(head2)

	return n
}

func ptl(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func _pttree(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%d ", root.Val)
	_pttree(root.Left)
	_pttree(root.Right)
}

func pttree(root *TreeNode) {
	_pttree(root)
	fmt.Println()
}

func main() {
	head := &ListNode{Val: 1}
	p := head
	p.Next = &ListNode{Val: 2}
	p = p.Next
	p.Next = &ListNode{Val: 3}
	p = p.Next
	p.Next = &ListNode{Val: 4}
	p = p.Next
	p.Next = &ListNode{Val: 5}
	p = p.Next

	tr := sortedListToBST(head)

	pttree(tr)
}
