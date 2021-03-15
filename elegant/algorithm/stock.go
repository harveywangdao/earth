package main

import (
	"fmt"
	"math"
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

func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minprice, profit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-minprice > profit {
			profit = prices[i] - minprice
		}
		if prices[i] < minprice {
			minprice = prices[i]
		}
	}
	return profit
}

/*
dp[i][j] i天的利润
j:0 未持有股票
j:1 持有股票
dp[i][0] = max(dp[i-1][0], dp[i-1][1]+p[i])
dp[i][1] = max(dp[i-1][1], -p[i])
*/
func maxProfit1_dg(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, -prices[i])
	}
	return dp0
}

func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	profit := 0
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			profit += (prices[i] - prices[i-1])
		}
	}
	return profit
}

/*
dp[i][0] = max(dp[i-1][0], dp[i-1][1]+p[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0]-p[i])
*/
func maxProfit2_dg(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		tmp := dp0
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, tmp-prices[i])
	}
	return dp0
}

func maxProfit3_dg0(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2
}

func maxProfit3_dg1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp_k1_0 := 0
	dp_k1_1 := -prices[0]
	dp_k2_0 := 0
	dp_k2_1 := -prices[0]
	for i := 1; i < n; i++ {
		dp_k2_0 = max(dp_k2_0, dp_k2_1+prices[i])
		dp_k2_1 = max(dp_k2_1, dp_k1_0-prices[i])
		dp_k1_0 = max(dp_k1_0, dp_k1_1+prices[i])
		dp_k1_1 = max(dp_k1_1, -prices[i])
	}
	return dp_k2_0
}

//今天没有股票
//dp[i][k][0] = max(a, b)
//1.昨天也没有股票  dp[i-1][k][0]
//2.今天卖了股票    dp[i-1][k][1]  + prices[i]

//今天有股票
//dp[i][k][1] = max(a, b)
//1.昨天也有股票    dp[i-1][k][1]
//2.今天买了股票    dp[i-1][k-1][0] - prices[i]

func maxProfit4_dg0(k int, prices []int) int {
	n := len(prices)
	maxk := min(n/2, k)
	if maxk == 0 {
		return 0
	}

	dp := make([][][2]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][2]int, maxk+1)
	}

	for i := 0; i < n; i++ {
		for k := 1; k <= maxk; k++ {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}

			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}

	return dp[n-1][maxk][0]
}

func maxProfit4_dg1(k int, prices []int) int {
	n := len(prices)
	maxk := min(n/2, k)
	if maxk == 0 {
		return 0
	}
	dp := make([][2]int, maxk+1)
	for i := 0; i < n; i++ {
		for k := 1; k <= maxk; k++ {
			if i == 0 {
				dp[k][0] = 0
				dp[k][1] = -prices[i]
				continue
			}

			dp[k][0] = max(dp[k][0], dp[k][1]+prices[i])
			dp[k][1] = max(dp[k][1], dp[k-1][0]-prices[i])
		}
	}

	return dp[maxk][0]
}

/*
i天结束之后的总收益
有 dp[i][0] = max(dp[i-1][0], dp[i-1][2]-p[i])
冷 dp[i][1] = dp[i-1][0] + p[i]
无 dp[i][2] = max(dp[i-1][2], dp[i-1][1])
*/
func maxProfit5_dg0(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][3]int, n)
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][2], dp[i-1][1])
	}
	return max(dp[n-1][1], dp[n-1][2])
}

func maxProfit5_dg1(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp0 := -prices[0]
	dp1 := 0
	dp2 := 0
	for i := 1; i < n; i++ {
		tmp := dp2
		dp2 = max(dp2, dp1)
		dp1 = dp0 + prices[i]
		dp0 = max(dp0, tmp-prices[i])
	}
	return max(dp1, dp2)
}

/*
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i]-fee)
*/
func maxProfit6_dg0(prices []int, fee int) int {
	n := len(prices)
	maxk := min(n/2, math.MaxInt32)
	if maxk == 0 {
		return 0
	}
	dp := make([][2]int, maxk+1)
	for i := 0; i < n; i++ {
		for k := 1; k <= maxk; k++ {
			if i == 0 {
				dp[k][0] = 0
				dp[k][1] = -prices[i] - fee
				continue
			}
			dp[k][0] = max(dp[k][0], dp[k][1]+prices[i])
			dp[k][1] = max(dp[k][1], dp[k-1][0]-prices[i]-fee)
		}
	}
	return dp[maxk][0]
}

/*
dp[i][0] = max(dp[i-1][0], dp[i-1][1]+p[i]-fee)
dp[i][1] = max(dp[i-1][1], dp[i-1][0]-p[i])
*/
func maxProfit6_dg1(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		tmp := dp0
		dp0 = max(dp0, dp1+prices[i]-fee)
		dp1 = max(dp1, tmp-prices[i])
	}
	return dp0
}

/*
dp[i][0] = max(dp[i-1][0], dp[i-1][1]+p[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0]-p[i]-fee)
*/
func maxProfit6_dg2(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp0, dp1 := 0, -prices[0]-fee
	for i := 1; i < n; i++ {
		tmp := dp0
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, tmp-prices[i]-fee)
	}
	return dp0
}

func main() {
	fmt.Println(maxProfit4(4, []int{3, 3, 5, 0, 0, 3, 1, 4}))
}
