package main

import (
	"fmt"
)

type Node struct {
	Val       int
	Neighbors []*Node
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

func findTargetSumWays1(nums []int, S int) int {
	count := 0
	_findTargetSumWays1(nums, S, &count)
	return count
}

func _findTargetSumWays1(nums []int, S int, count *int) {
	if len(nums) == 1 {
		if nums[0] == S {
			(*count)++
		}
		if -nums[0] == S {
			(*count)++
		}
		return
	}

	_findTargetSumWays1(nums[1:], S-nums[0], count)
	_findTargetSumWays1(nums[1:], S+nums[0], count)
}

func findTargetSumWays2(nums []int, S int) int {
	n := len(nums)
	if S < 0 {
		S = -S
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if S > sum {
		return 0
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, sum+1)
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][nums[i]] += 1
			if nums[i] == 0 {
				dp[i][nums[i]] += 1
			}
			continue
		}

		for j := 0; j <= sum; j++ {
			if dp[i-1][j] == 0 {
				continue
			}
			dp[i][j+nums[i]] += dp[i-1][j]
			if j-nums[i] >= 0 {
				dp[i][j-nums[i]] += dp[i-1][j]
			}
			if j > 0 && -j+nums[i] >= 0 {
				dp[i][-j+nums[i]] += dp[i-1][j]
			}
			if j > 0 && -j-nums[i] >= 0 {
				dp[i][-j-nums[i]] += dp[i-1][j]
			}
		}
	}

	return dp[n-1][S]
}

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	if S < 0 {
		S = -S
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if S > sum {
		return 0
	}

	dp := make([]int, sum+1)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[nums[i]] += 1
			if nums[i] == 0 {
				dp[nums[i]] += 1
			}
			continue
		}

		dp2 := make([]int, sum+1)
		for j := 0; j <= sum; j++ {
			if dp[j] == 0 {
				continue
			}

			dp2[j+nums[i]] += dp[j]
			if j-nums[i] >= 0 {
				dp2[j-nums[i]] += dp[j]
			}
			if j > 0 && -j+nums[i] >= 0 {
				dp2[-j+nums[i]] += dp[j]
			}
			if j > 0 && -j-nums[i] >= 0 {
				dp2[-j-nums[i]] += dp[j]
			}
		}
		dp = dp2
	}

	return dp[S]
}

func main() {
	fmt.Println(findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}
