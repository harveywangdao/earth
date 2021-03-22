package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution1 struct {
	origin []int
	ran    *rand.Rand
}

func Constructor1(nums []int) Solution1 {
	origin := make([]int, len(nums))
	copy(origin, nums)

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	return Solution1{
		origin: origin,
		ran:    r,
	}
}

func (this *Solution1) Reset() []int {
	nums := make([]int, len(this.origin))
	copy(nums, this.origin)
	return nums
}

func (this *Solution1) Shuffle() []int {
	nums := make([]int, len(this.origin))
	rs := this.ran.Perm(len(this.origin))

	for i := 0; i < len(rs); i++ {
		nums[i] = this.origin[rs[i]]
	}

	return nums
}

type Solution struct {
	origin []int
	arr    []int
	ran    *rand.Rand
}

func Constructor(nums []int) Solution {
	origin := make([]int, len(nums))
	copy(origin, nums)

	return Solution{
		origin: origin,
		arr:    nums,
		ran:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *Solution) Reset() []int {
	this.arr = this.origin
	nums := make([]int, len(this.origin))
	copy(nums, this.origin)
	this.origin = nums
	return this.origin
}

// 洗牌算法
func (this *Solution) Shuffle() []int {
	n := len(this.origin)

	for i := 0; i < n; i++ {
		j := this.ran.Intn(n-i) + i
		this.arr[i], this.arr[j] = this.arr[j], this.arr[i]
	}

	return this.arr
}

func fizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		str := ""
		if i%3 == 0 {
			str = "Fizz"
		}
		if i%5 == 0 {
			str = str + "Buzz"
		}
		if str == "" {
			str = strconv.Itoa(i)
		}
		res = append(res, str)
	}
	return res
}

func missingNumber1(nums []int) int {
	n := len(nums)
	xor := 0
	for i := 0; i < n; i++ {
		xor ^= i
		xor ^= nums[i]
	}
	xor ^= n
	return xor
}

func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for i := 0; i < n; i++ {
		sum -= nums[i]
	}
	return sum
}

func findMin(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return nums[i+1]
		}
	}
	return nums[0]
}

func findMin1(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] >= nums[0] {
			l = mid + 1
		} else if nums[mid] < nums[0] {
			r = mid
		}
	}
	if l == len(nums) {
		return nums[0]
	}
	return nums[l]
}

func findMin2(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

func main() {

}
