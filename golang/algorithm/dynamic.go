package main

import (
	"fmt"
	"math"
)

func maxSubArray1(nums []int) int {
	n := len(nums)

	max := math.MinInt64
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if max < sum {
				max = sum
			}
		}
	}
	return max
}

func maxSubArray2(nums []int) int {
	n := len(nums)

	lower := nums[0]
	height := nums[0]
	if nums[0] > 0 {
		lower = 0
	}

	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]

		if nums[i]-lower > height {
			height = nums[i] - lower
		}

		if nums[i] < lower {
			lower = nums[i]
		}
	}

	return height
}

func maxSubArray3(nums []int) int {
	n := len(nums)

	dp := make([][2]int, n)
	dp[0][0] = math.MinInt32
	dp[0][1] = nums[0]

	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = max(nums[i], dp[i-1][1]+nums[i])
	}

	return max(dp[n-1][0], dp[n-1][1])
}

func maxSubArray4(nums []int) int {
	n := len(nums)

	dp0 := math.MinInt32
	dp1 := nums[0]

	for i := 1; i < n; i++ {
		dp0 = max(dp0, dp1)
		dp1 = max(nums[i], dp1+nums[i])
	}

	return max(dp0, dp1)
}

func maxSubArray5(nums []int) int {
	n := len(nums)
	// dp[i] = max(dp[i-1]+nums[i], nums[i])
	res := nums[0]
	for i := 1; i < n; i++ {
		nums[i] = max(nums[i-1]+nums[i], nums[i])
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

func main() {

}
