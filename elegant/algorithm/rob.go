package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
dp[i][0]=max(dp[i-1][0], dp[i-1][1])
dp[i][1]=dp[i-1][0]+p[i]
*/
func rob1_dg0(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = nums[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return max(dp[n-1][0], dp[n-1][1])
}

func rob1_dg1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp0, dp1 := 0, nums[0]
	for i := 1; i < n; i++ {
		dp0, dp1 = max(dp0, dp1), dp0+nums[i]
	}
	return max(dp0, dp1)
}

/*
dp[i] = max(dp[i-2]+nums[i], dp[i-1])
*/
func rob1_dg2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	dp0, dp1 := nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp0, dp1 = dp1, max(dp0+nums[i], dp1)
	}
	return dp1
}

func rob2_dg0(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n < 3 {
		return rob1_dg1(nums)
	}

	r1 := rob1_dg1(nums[1 : n-1])
	r2 := nums[0] + rob1_dg1(nums[2:n-1])
	r3 := nums[n-1] + rob1_dg1(nums[1:n-2])
	return max(max(r1, r2), r3)
}

func rob2_dg1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n < 3 {
		return rob1_dg1(nums)
	}

	r1 := rob1_dg1(nums[1:])
	r2 := rob1_dg1(nums[:n-1])
	return max(r1, r2)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
dp[i] = max(dp[i-2]+nums[i], dp[i-1])
*/
func rob(root *TreeNode) int {
	n, _, _ := _rob(root)
	return n
}

func _rob(root *TreeNode) (int, int, int) {
	if root == nil {
		return 0, 0, 0
	}
	left, ll, lr := _rob(root.Left)
	right, rl, rr := _rob(root.Right)

	return max(left+right, ll+lr+rl+rr+root.Val), left, right
}

func main() {
	fmt.Println(rob2_dg0([]int{2, 3, 2}))
	fmt.Println(rob2_dg0([]int{1, 2, 3, 1}))
}
