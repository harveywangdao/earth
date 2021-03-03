package main

import (
	"fmt"
)

func bag(n, w int, wt, vt []int) int {
	count := 1 << n

	maxv := 0
	for i := 0; i < count; i++ {
		ww, vv := 0, 0
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				ww += wt[j]
				vv += vt[j]
			}
		}

		if ww <= w && vv > maxv {
			maxv = vv
		}
	}

	return maxv
}

func knapsack(N, W int, wt, val []int) int {
	// vector 全填入 0，base case 已初始化
	dp := make([][]int, N+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, W+1)
	}
	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if w-wt[i-1] < 0 {
				// 当前背包容量装不下，只能选择不装入背包
				dp[i][w] = dp[i-1][w]
			} else {
				// 装入或者不装入背包，择优
				dp[i][w] = max(dp[i-1][w-wt[i-1]]+val[i-1], dp[i-1][w])
			}
		}
	}

	return dp[N][W]
}

func knapsack2(N, W int, wt, val []int) int {
	// vector 全填入 0，base case 已初始化
	dp := make([][]int, N+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, W+1)
	}
	for w := 1; w <= W; w++ {
		for i := 1; i <= N; i++ {
			if w-wt[i-1] < 0 {
				// 当前背包容量装不下，只能选择不装入背包
				dp[i][w] = dp[i-1][w]
			} else {
				// 装入或者不装入背包，择优
				dp[i][w] = max(dp[i-1][w-wt[i-1]]+val[i-1], dp[i-1][w])
			}
		}
	}

	return dp[N][W]
}

func knapsack3(N, W int, wt, val []int) int {
	var dp [10][10]int
	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if w-wt[i-1] < 0 {
				// 当前背包容量装不下，只能选择不装入背包
				dp[i][w] = dp[i-1][w]
			} else {
				// 装入或者不装入背包，择优
				dp[i][w] = max(dp[i-1][w-wt[i-1]]+val[i-1], dp[i-1][w])
			}
		}
	}

	return dp[N][W]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(bag(3, 4, []int{2, 1, 4}, []int{4, 2, 3}))
	fmt.Println(knapsack(3, 4, []int{2, 1, 4}, []int{4, 2, 3}))
	fmt.Println(knapsack2(3, 4, []int{2, 1, 4}, []int{4, 2, 3}))
	fmt.Println(knapsack3(3, 4, []int{2, 1, 4}, []int{4, 2, 3}))
}
