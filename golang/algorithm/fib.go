package main

import (
	"fmt"
	"math"
	"time"
)

var count int

func fib1(n int) int {
	count++
	if n < 2 {
		return n
	}

	return fib1(n-1) + fib1(n-2)
}

func fib2_1(n int) int {
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		count++
		p = q
		q = r
		r = p + q
	}
	return r
}

func fib2(n int) int {
	if n < 2 {
		return n
	}

	n0, n1 := 0, 1
	for i := 1; i < n; i++ {
		count++
		n1, n0 = n1+n0, n1
	}
	return n1
}

func fib3(n int) int {
	sqrt5 := math.Sqrt(5)
	p1 := math.Pow((1+sqrt5)/2, float64(n))
	p2 := math.Pow((1-sqrt5)/2, float64(n))
	return int(math.Round((p1 - p2) / sqrt5))
}

type matrix [2][2]int

func multiply(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			count++
			c[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return
}

func pow(a matrix, n int) matrix {
	ret := matrix{{1, 0}, {0, 1}}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ret = multiply(ret, a)
		}
		a = multiply(a, a)
	}
	return ret
}

func fib4(n int) int {
	if n < 2 {
		return n
	}
	res := pow(matrix{{1, 1}, {1, 0}}, n-1)
	return res[0][0]
}

func fib5(n int) int {
	if n < 1 {
		return 0
	}

	memo := make([]int, n+1)
	return fib5_helper(memo, n)
}

func fib5_helper(memo []int, n int) int {
	count++
	if n == 1 || n == 2 {
		return 1
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = fib5_helper(memo, n-1) + fib5_helper(memo, n-2)
	return memo[n]
}

func main() {
	var now time.Time

	//count = 0
	//now = time.Now()
	//fmt.Println(fib1(40), "fib recursive cost:", time.Now().Sub(now), "count:", count)

	count = 0
	now = time.Now()
	fmt.Println(fib2_1(70), "fib cyclic 1 cost:", time.Now().Sub(now), "count:", count)

	count = 0
	now = time.Now()
	fmt.Println(fib2(70), "fib cyclic 2 cost:", time.Now().Sub(now), "count:", count)

	count = 0
	now = time.Now()
	fmt.Println(fib3(70), "fib math cost:", time.Now().Sub(now), "count:", count)

	count = 0
	now = time.Now()
	fmt.Println(fib4(70), "fib matrix cost:", time.Now().Sub(now), "count:", count)

	count = 0
	now = time.Now()
	fmt.Println(fib5(70), "fib memo cost:", time.Now().Sub(now), "count:", count)
}

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	n0, n1 := 1, 2
	for i := 2; i < n; i++ {
		n1, n0 = n1+n0, n1
	}
	return n1
}
