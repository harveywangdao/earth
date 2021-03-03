package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

func (this *Solution) GetRandom() int {
	p := this.head
	val, i := 0, 0
	for p != nil {
		i++
		if rand.Intn(i) == 0 {
			val = p.Val
		}
		p = p.Next
	}
	return val
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	if p == nil {
		return head
	}
	for p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}

func glue() {
	ll := list.New()
	ll.PushFront(11)
	ll.PushFront(22)
	ll.PushBack(33)
	for l := ll.Front(); l != nil; l = l.Next() {
		fmt.Println(l.Value)
	}
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Println() {
	fmt.Println(h)
}

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func calendar() {
	h := &IntHeap{2, 1, 5, 100, 3, 6, 4, 5}

	heap.Init(h)
	h.Println()

	heap.Push(h, 17)
	h.Println()
	heap.Push(h, -1)
	h.Println()

	(*h)[4] = 20
	heap.Fix(h, 4)
	h.Println()

	heap.Pop(h)
	h.Println()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	p := head
	var lp, rp *ListNode
	for p != nil {
		rp = p.Next
		p.Next = lp
		lp = p
		p = rp
	}
	return lp
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var p, head, tmp *ListNode
	for {
		if l1 == nil {
			if p == nil {
				head = l2
			} else {
				p.Next = l2
			}
			break
		} else if l2 == nil {
			if p == nil {
				head = l1
			} else {
				p.Next = l1
			}
			break
		}

		if l1.Val < l2.Val {
			tmp = l1
			l1 = l1.Next
		} else {
			tmp = l2
			l2 = l2.Next
		}

		if p != nil {
			p.Next = tmp
			p = tmp
		} else {
			p, head = tmp, tmp
		}
	}
	return head
}

func mergeKLists1(lists []*ListNode) *ListNode {
	var p, head, tmp *ListNode

	for {
		pos := -1
		for i := 0; i < len(lists); i++ {
			if pos != -1 {
				if lists[i] != nil && lists[i].Val < lists[pos].Val {
					pos = i
				}
			} else {
				if lists[i] != nil {
					pos = i
				}
			}
		}

		if pos == -1 {
			break
		}

		tmp = lists[pos]
		lists[pos] = lists[pos].Next

		if p != nil {
			p.Next = tmp
			p = tmp
		} else {
			p, head = tmp, tmp
		}
	}
	return head
}

func _heapsort(nums []*ListNode, i, n int) {
	min := i
	left := 2*i + 1
	if left < n && nums[min].Val > nums[left].Val {
		min = left
	}

	right := left + 1
	if right < n && nums[min].Val > nums[right].Val {
		min = right
	}

	if min != i {
		nums[i], nums[min] = nums[min], nums[i]
		_heapsort(nums, min, n)
	}
}

func mergeKLists2(lists []*ListNode) *ListNode {
	var p, head *ListNode
	n := len(lists)
	i, j := 0, n-1
	for i <= j {
		if lists[i] == nil {
			lists[i], lists[j] = lists[j], lists[i]
			j--
		} else {
			i++
		}
	}

	n = i
	for i := n/2 - 1; i >= 0; i-- {
		_heapsort(lists, i, n)
	}

	for n > 0 {
		if n == 1 {
			if p == nil {
				head = lists[0]
			} else {
				p.Next = lists[0]
			}
			break
		}

		if p != nil {
			p.Next = lists[0]
			p = lists[0]
		} else {
			p, head = lists[0], lists[0]
		}

		if lists[0].Next != nil {
			lists[0] = lists[0].Next
		} else {
			lists[0], lists[n-1] = lists[n-1], lists[0]
			n--
		}

		_heapsort(lists, 0, n)
	}
	return head
}

func mergeKLists(lists []*ListNode) *ListNode {
	return _mergeKLists(lists, 0, len(lists)-1)
}

func _mergeKLists(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	} else if l > r {
		return nil
	}
	mid := (l + r) / 2
	return mergeTwoLists(_mergeKLists(lists, l, mid), _mergeKLists(lists, mid+1, r))
}

func rotateRight1(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	n := 1
	p := head
	for p.Next != nil {
		n++
		p = p.Next
	}
	tail := p

	k %= n
	if k == 0 {
		return head
	}

	p = head
	for i := 1; i < n-k; i++ {
		p = p.Next
	}

	newhead := p.Next
	p.Next = nil
	tail.Next = head
	return newhead
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	old_tail := head
	n := 0
	for n = 1; old_tail.Next != nil; n++ {
		old_tail = old_tail.Next
	}
	old_tail.Next = head

	new_tail := head
	for i := 0; i < n-k%n-1; i++ {
		new_tail = new_tail.Next
	}
	new_head := new_tail.Next
	new_tail.Next = nil
	return new_head
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for {
				if p == slow {
					return p
				}
				p = p.Next
				slow = slow.Next
			}
		}
	}
	return nil
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reorderList1(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	rp := reverseList(mid)

	p, lp := &ListNode{}, head
	for p != lp {
		p.Next, lp = lp, lp.Next
		p = p.Next

		if p == rp {
			break
		}
		p.Next, rp = rp, rp.Next
		p = p.Next
	}
	p.Next = nil
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	right := reverseList(mid.Next)
	mid.Next = nil

	p := &ListNode{}
	l1, l2 := head, right
	for {
		if l1 == nil {
			p.Next = l2
			break
		}
		if l2 == nil {
			p.Next = l1
			break
		}

		p.Next, l1 = l1, l1.Next
		p = p.Next

		p.Next, l2 = l2, l2.Next
		p = p.Next
	}
}

func main() {

}
