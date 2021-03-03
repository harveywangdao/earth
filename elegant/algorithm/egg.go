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

func superEggDrop1(K int, N int) int {
	if K == 1 {
		return N
	}

	dp := make([][]int, K+1)
	for i := 1; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}

	for i := 1; i <= K; i++ {
		for j := 1; j <= N; j++ {
			if i == 1 {
				dp[1][j] = j
			} else {
				if j == 1 {
					dp[i][1] = 1
				} else {
					dp[i][j] = math.MaxInt64
					for x := 1; x <= j; x++ {
						dp[i][j] = min(dp[i][j], max(1+dp[i-1][x-1], 1+dp[i][j-x]))
					}
				}
			}
		}
	}

	return dp[K][N]
}

func superEggDrop2(K int, N int) int {
	if K == 1 {
		return N
	}

	dp := make([][]int, K+1)
	for i := 1; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}

	for i := 1; i <= K; i++ {
		for j := 1; j <= N; j++ {
			if i == 1 {
				dp[1][j] = j
			} else {
				dp[i][j] = math.MaxInt64
				for x := 1; x <= j; x++ {
					dp[i][j] = min(dp[i][j], max(1+dp[i-1][x-1], 1+dp[i][j-x]))
				}
			}
		}
	}

	for i := 1; i < len(dp); i++ {
		fmt.Println(dp[i][1:])
	}

	return dp[K][N]
}

func superEggDrop3(K int, N int) int {
	if K == 1 {
		return N
	}

	dp := make([][]int, K+1)
	for i := 1; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}

	for i := 1; i <= K; i++ {
		for j := 1; j <= N; j++ {
			if i == 1 {
				dp[1][j] = j
			} else {
				if dp[i][j-1] == dp[i-1][j] {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = math.MaxInt64
					for x := 1; x <= j; x++ {
						dp[i][j] = min(dp[i][j], 1+max(dp[i-1][x-1], dp[i][j-x]))
					}
				}
			}
		}
	}

	return dp[K][N]
}

func superEggDrop4(K int, N int) int {
	if K == 1 {
		return N
	}

	dp := make([][]int, K+1)
	for i := 1; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}

	for i := 1; i <= K; i++ {
		for j := 1; j <= N; j++ {
			if i == 1 {
				dp[1][j] = j
			} else {
				if dp[i][j-1] == dp[i-1][j] {
					dp[i][j] = dp[i][j-1]
				} else {
					dp[i][j] = math.MaxInt64
					for x := 1; x <= j; x++ {
						dp[i][j] = min(dp[i][j], 1+max(dp[i-1][x-1], dp[i][j-x]))
						if dp[i][j] == dp[i][j-1] {
							break
						}
					}
				}
			}
		}
	}

	return dp[K][N]
}

func superEggDrop5(K int, N int) int {
	if K == 1 {
		return N
	}

	dp := make([][]int, K+1)
	for i := 1; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}

	for i := 1; i <= K; i++ {
		for j := 1; j <= N; j++ {
			if i == 1 {
				dp[1][j] = j
			} else {
				if dp[i][j-1] == dp[i-1][j] {
					dp[i][j] = dp[i][j-1]
				} else {
					x0, x1 := 1, j
					l, r := 1, j
					for l <= r {
						x := (l + r) / 2
						if dp[i-1][x-1] > dp[i][j-x] {
							r = x - 1
							x1 = x
						} else if dp[i-1][x-1] < dp[i][j-x] {
							l = x + 1
							x0 = x
						} else {
							x0 = x
							x1 = x
							break
						}
					}
					dp[i][j] = math.MaxInt64
					dp[i][j] = min(dp[i][j], 1+max(dp[i-1][x0-1], dp[i][j-x0]))
					dp[i][j] = min(dp[i][j], 1+max(dp[i-1][x1-1], dp[i][j-x1]))
				}
			}
		}
	}

	return dp[K][N]
}

func superEggDrop6(K int, N int) int {
	if K == 1 {
		return N
	}

	dp := make([][]int, K+1)
	for i := 1; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}

	for i := 1; i <= K; i++ {
		x0 := 1
		for j := 1; j <= N; j++ {
			if i == 1 {
				dp[1][j] = j
			} else {
				for ; x0 < j && max(dp[i-1][x0-1], dp[i][j-x0]) > max(dp[i-1][x0], dp[i][j-x0-1]); x0++ {

				}
				dp[i][j] = 1 + max(dp[i-1][x0-1], dp[i][j-x0])
			}
		}
	}

	return dp[K][N]
}

func superEggDrop7(K int, N int) int {
	if K == 1 {
		return N
	}
	if N == 1 {
		return 1
	}

	dp := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		dp[i] = make([]int, K+1)
	}

	for j := 1; j <= K; j++ {
		dp[1][j] = 1
	}

	for i := 2; i <= N; i++ {
		for j := 1; j <= K; j++ {
			dp[i][j] = 1 + dp[i-1][j-1] + dp[i-1][j]
			if dp[i][j] >= N {
				return i
			}
		}
	}
	return 0
}

func superEggDrop8(K int, N int) int {
	if K == 1 {
		return N
	}
	if N == 1 {
		return 1
	}

	dp := make([]int, K+1)
	for j := 1; j <= K; j++ {
		dp[j] = 1
	}

	var last int
	for i := 2; i <= N; i++ {
		last = 0
		for j := 1; j <= K; j++ {
			dp[j], last = 1+last+dp[j], dp[j]
			if dp[j] >= N {
				return i
			}
		}
	}
	return 0
}

func superEggDrop(K int, N int) int {
	if K == 1 {
		return N
	}

	dp := make([]int, N+1)
	var sum, x int
	for i := 1; i <= K; i++ {
		sum = 0
		for j := 1; j <= N; j++ {
			if i == 1 {
				dp[j] = 1
			} else {
				if j == 1 {
					dp[j] = 1
					x = 1
				} else {
					//dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
					dp[j], x = dp[j-1]+x, dp[j]
				}
			}
			sum += dp[j]
			if sum >= N {
				if i == K {
					//fmt.Println(dp)
					return j
				}
				break
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(superEggDrop(2, 100))
	fmt.Println(superEggDrop1(2, 100))
	fmt.Println(superEggDrop2(2, 100))
	fmt.Println(superEggDrop3(2, 100))
	fmt.Println(superEggDrop4(2, 100))
	fmt.Println(superEggDrop5(2, 100))
	fmt.Println(superEggDrop6(2, 100))
	fmt.Println(superEggDrop7(2, 100))
	fmt.Println(superEggDrop8(2, 100))
}
