package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

var count int

func coinChange(coins []int, amount int) int {
	count = 0
	if amount < 1 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}

	sort.Ints(coins)
	res := math.MaxInt32
	_coinChange(coins, amount, 0, &res)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func _coinChange(coins []int, amount, cur int, res *int) {
	lastIndex := len(coins) - 1
	for num := amount / coins[lastIndex]; num >= 0; num-- {
		count++
		remain := amount - num*coins[lastIndex]
		if remain > 0 {
			if lastIndex == 0 {
				break
			}

			if cur+num+1 >= *res {
				break
			}

			_coinChange(coins[:lastIndex], remain, cur+num, res)
		} else {
			*res = min(*res, cur+num)
			break
		}
	}
}

func coinChange2(coins []int, amount int) int {
	count = 0
	if amount < 1 {
		return 0
	}

	memo := make([]int, amount)
	return _coinChange2(coins, amount, memo)
}

func _coinChange2(coins []int, amount int, memo []int) int {
	min, num, amo := -1, 0, 0
	for i := 0; i < len(coins); i++ {
		count++
		amo = amount - coins[i]
		if amo > 0 {
			if memo[amo-1] != 0 {
				num = memo[amo-1]
			} else {
				num = _coinChange2(coins, amo, memo)
			}

			if num != -1 && (min == -1 || num < min) {
				min = num
			}
		} else if amo == 0 {
			min = 0
			break
		}
	}

	if min == -1 {
		memo[amount-1] = -1
	} else {
		memo[amount-1] = min + 1
	}

	return memo[amount-1]
}

func coinChange3(coins []int, amount int) int {
	count = 0
	if amount < 1 {
		return 0
	}

	memo := make([]int, amount+1)
	var amo int
	for i := 1; i <= amount; i++ {
		memo[i] = amount + 1
		for j := 0; j < len(coins); j++ {
			count++
			amo = i - coins[j]
			if amo >= 0 && memo[amo]+1 < memo[i] {
				memo[i] = memo[amo] + 1
			}
		}
	}

	if memo[amount] > amount {
		return -1
	}
	return memo[amount]
}

func coinChange4(coins []int, amount int) int {
	count = 0
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			count++
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func coinChange5(coins []int, amount int) int {
	count = 0
	res := math.MaxInt32
	sort.Ints(coins)

	backtrack(coins, len(coins)-1, amount, 0, &res)
	if res == math.MaxInt32 {
		return -1
	}

	return res
}

func backtrack(coins []int, coinIndex int, amount int, cur int, res *int) {
	if coinIndex < 0 {
		return
	}

	for n := amount / coins[coinIndex]; n >= 0; n-- {
		count++
		remain := amount - n*coins[coinIndex]
		curNum := cur + n

		if remain == 0 {
			*res = min(*res, curNum)
			break
		}

		if coinIndex == 0 {
			break
		}

		if curNum+1 >= *res {
			break
		}

		backtrack(coins, coinIndex-1, remain, curNum, res)
	}
}

func main() {
	var now time.Time

	now = time.Now()
	fmt.Println("coinChange F(6249) =", coinChange([]int{186, 419, 83, 408}, 6249), "cost:", time.Now().Sub(now), "count:", count)

	now = time.Now()
	fmt.Println("coinChange2 F(6249) =", coinChange2([]int{186, 419, 83, 408}, 6249), "cost:", time.Now().Sub(now), "count:", count)

	now = time.Now()
	fmt.Println("coinChange3 F(6249) =", coinChange3([]int{186, 419, 83, 408}, 6249), "cost:", time.Now().Sub(now), "count:", count)

	now = time.Now()
	fmt.Println("coinChange4 F(6249) =", coinChange4([]int{186, 419, 83, 408}, 6249), "cost:", time.Now().Sub(now), "count:", count)

	now = time.Now()
	fmt.Println("coinChange5 F(6249) =", coinChange5([]int{186, 419, 83, 408}, 6249), "cost:", time.Now().Sub(now), "count:", count)

	now = time.Now()
	fmt.Println("coinChange F(9208) =", coinChange([]int{288, 160, 10, 249, 40, 77, 314, 429}, 9208), "cost:", time.Now().Sub(now), "count:", count)

	now = time.Now()
	fmt.Println("coinChange2 F(9208) =", coinChange2([]int{288, 160, 10, 249, 40, 77, 314, 429}, 9208), "cost:", time.Now().Sub(now), "count:", count)

	now = time.Now()
	fmt.Println("coinChange3 F(9208) =", coinChange3([]int{288, 160, 10, 249, 40, 77, 314, 429}, 9208), "cost:", time.Now().Sub(now), "count:", count)

	now = time.Now()
	fmt.Println("coinChange4 F(9208) =", coinChange4([]int{288, 160, 10, 249, 40, 77, 314, 429}, 9208), "cost:", time.Now().Sub(now), "count:", count)

	now = time.Now()
	fmt.Println("coinChange5 F(9208) =", coinChange5([]int{288, 160, 10, 249, 40, 77, 314, 429}, 9208), "cost:", time.Now().Sub(now), "count:", count)
}
