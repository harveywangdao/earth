package main

import (
	"fmt"
	"math"
	"time"
)

func pow1(x, y int) int {
	a := 1
	for i := 0; i < y; i++ {
		a = a * x
	}
	return a
}

func pow2(x, y int) int {
	if y == 0 {
		return 1
	}

	if y%2 == 0 {
		return pow2(x*x, y>>1)
	} else {
		return x * pow2(x*x, y>>1)
	}
}

func pow3(x, n int) int {
	if n == 0 {
		return 1
	} else {
		for n&1 == 0 {
			n >>= 1
			x *= x
		}
	}

	result := x
	n >>= 1
	for n != 0 {
		x *= x
		if n&1 != 0 {
			result *= x
		}
		n >>= 1
	}
	return result
}

func pow4(x, y int) int {
	ret := 1
	for y > 0 {
		if y&1 != 0 {
			ret *= x
		}
		y >>= 1
		x *= x
	}

	return ret
}

func div1(x, n int) int {
	for i := 0; i < n; i++ {
		x = x / 2
	}
	return x
}

func div2(x, n int) int {
	for i := 0; i < n; i++ {
		x >>= 1
	}
	return x
}

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1.0 / x
		n = -n
	}

	ret := 1.0
	base := x
	for i := n; i > 0; i = i >> 1 {
		if i%2 != 0 {
			ret = ret * base
			i--
		}
		base = base * base
	}

	return ret
}

func main() {
	now := time.Now()
	fmt.Println(math.Pow(2, 61), "cost:", time.Now().Sub(now))

	now = time.Now()
	fmt.Println(pow1(2, 61), "cost:", time.Now().Sub(now))

	now = time.Now()
	fmt.Println(pow2(2, 61), "cost:", time.Now().Sub(now))

	now = time.Now()
	fmt.Println(pow3(2, 61), "cost:", time.Now().Sub(now))

	now = time.Now()
	fmt.Println(pow4(2, 61), "cost:", time.Now().Sub(now))

	now = time.Now()
	fmt.Println(div1(math.MaxInt64, 60), "cost:", time.Now().Sub(now))

	now = time.Now()
	fmt.Println(div2(math.MaxInt64, 60), "cost:", time.Now().Sub(now))
}
