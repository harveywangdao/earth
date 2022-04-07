package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func findMedinSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total == 0 {
		return 0.0
	}

	nums := make([]int, total)
	i, j, k := 0, 0, 0
	for {
		if i == len(nums1) {
			for j < len(nums2) {
				nums[k] = nums2[j]
				k++
				j++
			}
			break
		}

		if j == len(nums2) {
			for i < len(nums1) {
				nums[k] = nums1[i]
				k++
				i++
			}
			break
		}

		if nums1[i] <= nums2[j] {
			nums[k] = nums1[i]
			i++
			k++
		} else {
			nums[k] = nums2[j]
			j++
			k++
		}
	}

	if total%2 == 0 {
		return float64(nums[total/2]+nums[total/2-1]) / 2
	} else {
		return float64(nums[total/2])
	}
}

func findMedinSortedArrays2(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total == 0 {
		return float64(0)
	} else if total == 1 {
		if len(nums1) == 1 {
			return float64(nums1[0])
		} else {
			return float64(nums2[0])
		}
	}

	var i, j int
	var left, right int
	for {
		if i == len(nums1) {
			if i+j == (total+1)/2 {
				right = nums2[j]
			} else {
				left = nums2[(total-1)/2-i]
				right = nums2[(total-1)/2-i+1]
			}
			break
		}

		if j == len(nums2) {
			if i+j == (total+1)/2 {
				right = nums1[i]
			} else {
				left = nums1[(total-1)/2-j]
				right = nums1[(total-1)/2-j+1]
			}
			break
		}

		if i+j == (total+1)/2 {
			right = min(nums1[i], nums2[j])
			break
		}

		if nums1[i] <= nums2[j] {
			left = nums1[i]
			i++
		} else {
			left = nums2[j]
			j++
		}
	}

	if total%2 == 0 {
		return float64(left+right) / 2
	}
	return float64(left)
}

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

func findMedinSortedArrays3(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	var i, j int
	var medin1, medin2 int
	mid := total/2 + 1
	for i+j < mid {
		medin1 = medin2

		if i < len(nums1) && ((j < len(nums2) && nums1[i] <= nums2[j]) || j == len(nums2)) {
			medin2 = nums1[i]
			i++
		} else if j < len(nums2) && ((i < len(nums1) && nums1[i] > nums2[j]) || i == len(nums1)) {
			medin2 = nums2[j]
			j++
		} else {
			break
		}
	}

	if total%2 == 0 {
		return float64(medin1+medin2) / 2
	}
	return float64(medin2)
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	m, n = m-1, n-1

	for ; n >= 0; i, n = i-1, n-1 {
		for ; m >= 0 && nums1[m] > nums2[n]; i, m = i-1, m-1 {
			nums1[i] = nums1[m]
		}
		nums1[i] = nums2[n]
	}

	fmt.Println("nums1:", nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for {
		if i < 0 {
			copy(nums1, nums2[:j+1])
			break
		} else if j < 0 {
			break
		}

		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	fmt.Println("nums1:", nums1)
}

func findMedinSortedArrays4(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	total := len(nums1) + len(nums2)
	if total == 0 {
		return float64(0)
	} else if len(nums1) == 0 {
		if len(nums2)%2 == 0 {
			return float64(nums2[(total-1)/2]+nums2[total/2]) / 2
		} else {
			return float64(nums2[total/2])
		}
	}

	i := (len(nums1) - 1) / 2
	j := (total+1)/2 - i - 2
	for i >= 0 && i < len(nums1) && j >= 0 {
		if nums1[i] > nums2[j+1] {
			i, j = i-1, j+1
		} else if i+1 < len(nums1) && nums1[i+1] < nums2[j] {
			i, j = i+1, j-1
		} else {
			break
		}
	}

	var left_i, left_j, right_i, right_j int
	if i < 0 {
		left_i = math.MinInt32
	} else {
		left_i = nums1[i]
	}

	if j < 0 {
		left_j = math.MinInt32
	} else {
		left_j = nums2[j]
	}

	if i+1 >= len(nums1) {
		right_i = math.MaxInt32
	} else {
		right_i = nums1[i+1]
	}

	if j+1 >= len(nums2) {
		right_j = math.MaxInt32
	} else {
		right_j = nums2[j+1]
	}

	//fmt.Println(i, j, left_i, left_j, right_i, right_j)

	left := max(left_i, left_j)
	right := min(right_i, right_j)

	if total%2 == 0 {
		return float64(left+right) / 2
	} else {
		return float64(min(left, right))
	}
}

func findMedinSortedArrays5(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	total := len(nums1) + len(nums2)
	if total == 0 {
		return float64(0)
	}

	mid := (total + 1) / 2
	l, r := 0, len(nums1)
	for l < r {
		i := (r-l+1)/2 + l
		j := mid - i

		//fmt.Println(i, j, l, r)
		if nums1[i-1] > nums2[j] {
			r = i - 1
		} else {
			l = i
		}
	}

	i, j := l, mid-l
	//fmt.Println(i, j)
	var left_i, left_j, right_i, right_j int
	if i == 0 {
		left_i = math.MinInt32
	} else {
		left_i = nums1[i-1]
	}

	if j == 0 {
		left_j = math.MinInt32
	} else {
		left_j = nums2[j-1]
	}

	if i == len(nums1) {
		right_i = math.MaxInt32
	} else {
		right_i = nums1[i]
	}

	if j == len(nums2) {
		right_j = math.MaxInt32
	} else {
		right_j = nums2[j]
	}

	left := max(left_i, left_j)
	right := min(right_i, right_j)

	if total%2 == 0 {
		return float64(left+right) / 2
	} else {
		return float64(left)
	}
}

func findMedinSortedArrays6(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	total := len(nums1) + len(nums2)
	if total == 0 {
		return float64(0)
	}

	if total%2 != 0 {
		return float64(getKthElement(nums1, nums2, (total+1)/2))
	} else {
		return float64(getKthElement(nums1, nums2, total/2)+getKthElement(nums1, nums2, total/2+1)) / 2
	}
}

func getKthElement(nums1 []int, nums2 []int, k int) int {
	left, right := 0, len(nums1)
	for left < right {
		i := (right-left+1)/2 + left
		j := k - i

		if j < len(nums2) && nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			left = i
		}
	}

	i, j := left, k-left
	var left_i, left_j int
	if i == 0 {
		left_i = math.MinInt32
	} else {
		left_i = nums1[i-1]
	}

	if j == 0 {
		left_j = math.MinInt32
	} else {
		left_j = nums2[j-1]
	}

	return max(left_i, left_j)
}

func merge33(intervals [][]int) [][]int {
	var A [][]int
	for i := 0; i < len(intervals); i++ {
		A = _merge(A, intervals[i])
	}
	return A
}

func _merge(A [][]int, B []int) [][]int {
	n := len(A)
	if n == 0 || A[n-1][1] < B[0] {
		A = append(A, B)
		return A
	}

	if A[n-1][0] > B[1] {
		tmp := []int{A[n-1][0], A[n-1][1]}
		left := _merge(A[:n-1], B)
		A = append(left, tmp)
		return A
	}

	A[n-1][0] = min(A[n-1][0], B[0])
	A[n-1][1] = max(A[n-1][1], B[1])
	return _merge(A[:n-1], A[n-1])
}

func merge(intervals [][]int) [][]int {
	/*for i := 0; i < len(intervals); i++ {
		for j := i; j > 0; j-- {
			if intervals[j][0] < intervals[j-1][0] {
				intervals[j], intervals[j-1] = intervals[j-1], intervals[j]
			} else {
				break
			}
		}
	}*/

	//_sort(intervals, 0, len(intervals)-1)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		n := len(res)
		if res[n-1][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			res[n-1][0] = min(res[n-1][0], intervals[i][0])
			res[n-1][1] = max(res[n-1][1], intervals[i][1])
		}
	}

	return res
}

func _sort(nums [][]int, low, high int) {
	if low >= high || low < 0 || high >= len(nums) {
		return
	}

	j := low
	for i := low; i < high; i++ {
		if nums[i][0] < nums[high][0] {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	nums[j], nums[high] = nums[high], nums[j]

	_sort(nums, low, j-1)
	_sort(nums, j+1, high)
}

func _sort2(nums [][]int, low, high int) {
	if low >= high {
		return
	}

	start, end, key := low, high, nums[low]
	for start < end {
		for start < end && nums[end][0] >= key[0] {
			end--
		}
		nums[start] = nums[end]

		for start < end && nums[start][0] <= key[0] {
			start++
		}
		nums[end] = nums[start]
	}
	nums[start] = key

	_sort2(nums, low, start-1)
	_sort2(nums, start+1, high)
}

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return (intervals[i][0] < intervals[j][0]) || (intervals[i][0] == intervals[j][0] && intervals[i][1] > intervals[j][1])
	})

	j, n := 0, 1
	for i := 1; i < len(intervals); i++ {
		if intervals[j][1] >= intervals[i][1] {
			continue
		} else {
			j, n = i, n+1
		}
	}
	return n
}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		if firstList[i][1] < secondList[j][0] {
			i++
			continue
		}

		if firstList[i][0] > secondList[j][1] {
			j++
			continue
		}

		if firstList[i][0] < secondList[j][0] {
			if firstList[i][1] < secondList[j][1] {
				res = append(res, []int{secondList[j][0], firstList[i][1]})
				i++
			} else {
				res = append(res, []int{secondList[j][0], secondList[j][1]})
				j++
			}
		} else {
			if firstList[i][1] > secondList[j][1] {
				res = append(res, []int{firstList[i][0], secondList[j][1]})
				j++
			} else {
				res = append(res, []int{firstList[i][0], firstList[i][1]})
				i++
			}
		}
	}

	return res
}

func intervalIntersection2(firstList [][]int, secondList [][]int) [][]int {
	var res [][]int
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		l := max(firstList[i][0], secondList[j][0])
		r := min(firstList[i][1], secondList[j][1])

		if l <= r {
			res = append(res, []int{l, r})
		}

		if firstList[i][1] > secondList[j][1] {
			j++
		} else {
			i++
		}
	}

	return res
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for z1 := 0; z1 < n-3; z1++ {
		if z1 > 0 && nums[z1] == nums[z1-1] {
			continue
		}

		for z2 := z1 + 1; z2 < n-2; z2++ {
			if z2 > z1+1 && nums[z2] == nums[z2-1] {
				continue
			}

			z4 := n - 1
			t := target - nums[z1] - nums[z2]
			for z3 := z2 + 1; z3 < n-1; z3++ {
				if z3 > z2+1 && nums[z3] == nums[z3-1] {
					continue
				}

				for z3 < z4 && nums[z3]+nums[z4] > t {
					z4--
				}
				if z4 == z3 {
					break
				}

				if nums[z3]+nums[z4] == t {
					ans = append(ans, []int{nums[z1], nums[z2], nums[z3], nums[z4]})
				}
			}
		}
	}
	return ans
}

func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	return nsum(nums, 3, 0, 0)
}

func fourSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nsum(nums, 4, 0, target)
}

func nsum(nums []int, n, start, target int) [][]int {
	res := make([][]int, 0)
	sz := len(nums)

	if n < 2 {
		return nil
	} else if n == 2 {
		l, r := start, sz-1
		for l < r {
			sum := nums[l] + nums[r]
			if sum > target {
				r--
			} else if sum < target {
				l++
			} else {
				res = append(res, []int{nums[l], nums[r]})
				l, r = l+1, r-1
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	} else {
		for i := start; i < sz-2; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			ret := nsum(nums, n-1, i+1, target-nums[i])
			for j := 0; j < len(ret); j++ {
				c := []int{nums[i]}
				c = append(c, ret[j]...)
				res = append(res, c)
			}
		}
	}

	return res
}

// 一般法
func countPrimes1(n int) int {
	var count int
	for i := 2; i < n; i++ {
		if isprime1(i) {
			count++
		}
	}
	return count
}

func isprime1_1(n int) bool {
	sqrtn := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtn; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isprime1_2(n int) bool {
	if n < 2 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}

	sqrtn := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtn; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 六素数法
func isprime1(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	k := int(math.Sqrt(float64(n))) + 1
	for i := 5; i < k; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// 一般法2
func countPrimes2(n int) int {
	if n < 3 {
		return 0
	}

	nums := []int{2}
	for i := 3; i < n; i = i + 2 {
		if isprime2(i, nums) {
			nums = append(nums, i)
		}
	}

	return len(nums)
}

func isprime2(n int, nums []int) bool {
	sqrtn := int(math.Sqrt(float64(n)))
	for i := 0; nums[i] <= sqrtn; i++ {
		if n%nums[i] == 0 {
			return false
		}
	}
	return true
}

// 埃氏筛变种
func countPrimes3(n int) int {
	composite := make([]bool, n)
	count := 0
	if n >= 3 {
		count++
		for j := 2; 2*j < n; j++ {
			composite[2*j] = true
		}
	}

	for i := 3; i < n; i += 2 {
		if !composite[i] {
			count++
			for j := 2; i*j < n; j++ {
				composite[i*j] = true
			}
		}
	}

	return count
}

// 埃氏筛
func countPrimes4(n int) int {
	isprime := make([]bool, n)
	for i := 0; i < n; i++ {
		isprime[i] = true
	}

	count := 0
	for i := 2; i < n; i++ {
		if isprime[i] {
			count++
			for j := 2 * i; j < n; j = j + i {
				isprime[j] = false
			}
		}
	}

	return count
}

// 埃氏筛变种2
func countPrimes5(n int) int {
	if n < 3 {
		return 0
	}

	isComposite := make([]bool, n)
	count := n / 2

	for i := 3; i*i < n; i += 2 {
		if isComposite[i] {
			continue
		}

		for j := i * i; j < n; j += 2 * i {
			//fmt.Println(i, j)
			if !isComposite[j] {
				count--

				isComposite[j] = true
			}
		}
	}

	return count
}

// 埃氏筛变种3
func countPrimes6(n int) int {
	isNotPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if isNotPrime[i] {
			continue
		}
		for j := i * i; j < n; j = j + i {
			isNotPrime[j] = true
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if !isNotPrime[i] {
			count++
		}
	}
	return count
}

// 线性筛
func countPrimes(n int) int {
	isprime := make([]bool, n)
	for i := 0; i < len(isprime); i++ {
		isprime[i] = true
	}
	primes := make([]int, 0, n/2)

	for i := 2; i < n; i++ {
		if isprime[i] {
			primes = append(primes, i)
		}

		for _, p := range primes {
			if i*p >= n {
				break
			}
			isprime[i*p] = false
			if i%p == 0 {
				break
			}
		}
	}
	return len(primes)
}

func superPow2(a int, b []int) int {
	p := 1337
	base := a % p

	n := 0
	for i := 0; i < len(b); i++ {
		n = 10*n + b[i]
	}

	res := 1
	for n > 0 {
		if n%2 != 0 {
			res = (res * base) % p
		}
		base = (base * base) % p
		n >>= 1
	}

	return res
}

func superPow3(a int, b []int) int {
	p := 1337
	base := a % p

	res := 1
	for {
		zero := true
		for i := 0; i < len(b); i++ {
			if b[i] != 0 {
				zero = false
				b = b[i:]
				break
			}
		}
		if zero {
			break
		}

		if b[len(b)-1]%2 != 0 {
			res = (res * base) % p
		}
		base = (base * base) % p

		m := 0
		for i := 0; i < len(b); i++ {
			b[i], m = (b[i]+10*m)/2, (b[i]+10*m)%2
		}
	}

	return res
}

// a ^ b % p
func superPow4(a int, b []int) int {
	bn := len(b)
	if bn == 0 {
		return 1
	}
	p := 1337

	left := npow(a, b[bn-1])
	m := superPow4(a, b[:bn-1])
	right := npow(m, 10)

	return (right * left) % p
}

func superPow5(a int, b []int) int {
	res := 1
	for i := 0; i < len(b); i++ {
		x := npow(res, 10)
		y := npow(a, b[i])
		res = (x * y) % 1337
	}
	return res
}

// a ^ n % p
func npow(a int, n int) int {
	p := 1337
	base := a % p

	res := 1
	for n > 0 {
		if n%2 != 0 {
			res = (res * base) % p
		}
		base = (base * base) % p
		n >>= 1
	}

	return res
}

func superPow(a int, b []int) int {
	base := a
	res := 1
	for i := len(b) - 1; i >= 0; i-- {
		res = (res * npow(base, b[i])) % 1337
		base = npow(base, 10)
	}

	return res
}

func trailingZeroes2(n int) int {
	if n < 5 {
		return 0
	}

	nfive := 0
	for i := 5; i <= n; i += 5 {
		nfive += fivecount(i)
	}
	return nfive
}

func fivecount(n int) int {
	if n < 5 {
		return 0
	}

	if n%5 != 0 {
		return 0
	}

	return 1 + fivecount(n/5)
}

func trailingZeroes3(n int) int {
	if n < 5 {
		return 0
	}
	return n/5 + trailingZeroes(n/5)
}

func trailingZeroes(n int) int {
	c := 0
	for n > 4 {
		n /= 5
		c += n
	}
	return c
}

func preimageSizeFZF2(K int) int {
	for n := 5 * (K + 1); n >= 0; n -= 5 {
		m := trailingZeroes(n)
		if m == K {
			return 5
		} else if m < K {
			break
		}
	}
	return 0
}

func preimageSizeFZF3(K int) int {
	n := 4 * K
	z := trailingZeroes(n)
	if z == K {
		return 5
	}
	for {
		n++
		if n%5 == 0 {
			break
		}
	}

	for {
		z = trailingZeroes(n)
		if z == K {
			return 5
		} else if z > K {
			return 0
		}
		n += 5
	}

	return 0
}

func preimageSizeFZF4(K int) int {
	l, r := 4*K, 5*K+1

	for l <= r {
		mid := (l + r) / 2
		z := trailingZeroes(mid)
		if z == K {
			return 5
		} else if z > K {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return 0
}

func preimageSizeFZF(K int) int {
	base := 0
	for base < K {
		base = base*5 + 1
	}
	for K > 0 {
		base = (base - 1) / 5
		if K/base == 5 {
			return 0
		}
		K %= base
	}
	return 5
}

type Solution struct {
	ns []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

// 蓄水池抽样算法
func (this *Solution) Pick(target int) int {
	n, index := 0, 0
	for i := 0; i < len(this.ns); i++ {
		if this.ns[i] == target {
			n++
			if rand.Intn(n) == 0 {
				index = i
			}
		}
	}
	return index
}

func removeDuplicates(nums []int) int {
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return i + 1
}

func removeElement(nums []int, val int) int {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}

func moveZeroes(nums []int) {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	for ; i < j; i++ {
		nums[i] = 0
	}
}

func canWinNim(n int) bool {
	if n <= 3 {
		return true
	}
	if n%4 == 0 {
		return false
	}
	return true
}

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

func subarraySum1(nums []int, k int) int {
	c := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				c++
			}
		}
	}
	return c
}

func subarraySum(nums []int, k int) int {
	c, pre := 0, 0
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre = pre + nums[i]
		c = c + m[pre-k]
		m[pre] = m[pre] + 1
	}
	return c
}

func corpFlightBookings1(bookings [][]int, n int) []int {
	arr := make([]int, n)
	for i := 0; i < len(bookings); i++ {
		for j := bookings[i][0]; j <= bookings[i][1]; j++ {
			arr[j-1] += bookings[i][2]
		}
	}
	return arr
}

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n+1)
	for i := 0; i < len(bookings); i++ {
		start := bookings[i][0]
		end := bookings[i][1]

		diff[start-1] += bookings[i][2]
		diff[end] -= bookings[i][2]
	}

	arr := make([]int, n)
	arr[0] = diff[0]
	for i := 1; i < len(diff)-1; i++ {
		arr[i] = arr[i-1] + diff[i]
	}

	return arr
}

func findKthLargest1(nums []int, k int) int {
	for i := 0; i < k; i++ {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	return nums[k-1]
}

func _heapsort(nums []int, i, n int) {
	min := i
	left := 2*i + 1
	if left < n && nums[min] > nums[left] {
		min = left
	}

	right := left + 1
	if right < n && nums[min] > nums[right] {
		min = right
	}

	if min != i {
		nums[i], nums[min] = nums[min], nums[i]
		_heapsort(nums, min, n)
	}
}

func findKthLargest2(nums []int, k int) int {
	for i := k/2 - 1; i >= 0; i-- {
		_heapsort(nums, i, k)
	}

	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i]
			_heapsort(nums, 0, k)
		}
	}

	return nums[0]
}

func _quicksort(nums []int, start, end int) {
	if start >= end {
		return
	}

	key := nums[start]
	i, j := start, end
	for i < j {
		for i < j && nums[j] >= key {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= key {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = key

	_quicksort(nums, i+1, end)
	_quicksort(nums, start, i-1)
}

func findKthLargest3(nums []int, k int) int {
	_findKthLargest(nums, 0, len(nums)-1, k)
	return nums[len(nums)-k]
}

func _findKthLargest(nums []int, start, end, k int) {
	if start > end {
		return
	}

	key := nums[start]
	i, j := start, end
	for i < j {
		for i < j && nums[j] >= key {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= key {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = key

	if end-i+1 == k {
		return
	} else if end-i+1 > k {
		_findKthLargest(nums, i+1, end, k)
	} else if end-i+1 < k {
		_findKthLargest(nums, start, i-1, k-(end-i+1))
	}
}

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := randomPartition(a, l, r)
	if q == index {
		return a[q]
	} else if q < index {
		return quickSelect(a, q+1, r, index)
	}
	return quickSelect(a, l, q-1, index)
}

func randomPartition(a []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

func partition(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func canJump1(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			if !_canJump(nums, i) {
				return false
			}
		}
	}
	return true
}

func _canJump(nums []int, zeroIndex int) bool {
	for i := zeroIndex - 1; i >= 0; i-- {
		if nums[i] > zeroIndex-i {
			return true
		}
	}
	return false
}

func canJump(nums []int) bool {
	maxstep := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxstep < i {
			return false
		}
		if nums[i]+i > maxstep {
			maxstep = i + nums[i]
		}
		if maxstep >= len(nums)-1 {
			return true
		}
	}
	return true
}

func jump1(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 1; j <= nums[i] && i+j < n; j++ {
			if dp[i+j] == 0 {
				dp[i+j] = dp[i] + 1
			}
		}
		if dp[n-1] != 0 {
			break
		}
	}
	return dp[n-1]
}

func jump2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	right := 0
	for i := 0; i < n; i++ {
		for j := right + 1; j <= i+nums[i] && j < n; j++ {
			dp[j] = dp[i] + 1
		}
		if dp[n-1] != 0 {
			break
		}
		if i+nums[i] > right {
			right = i + nums[i]
		}
	}
	return dp[n-1]
}

func jump(nums []int) int {
	n := len(nums)
	right1 := 0
	right2 := 0
	step := 0
	for i := 0; i < n-1; i++ {
		if i+nums[i] > right2 {
			right2 = i + nums[i]
		}
		if i >= right1 {
			right1 = right2
			step++
		}
		if right1 >= n-1 {
			break
		}
	}
	return step
}

func maxScore1(cardPoints []int, k int) int {
	n := len(cardPoints)
	left, right, max := 0, 0, 0
	for i := 0; i <= k; i++ {
		right = 0
		for j := n - 1; j >= n-k+i && j >= i; j-- {
			right += cardPoints[j]
		}

		if left+right > max {
			max = left + right
		}

		if i < n {
			left += cardPoints[i]
		}
	}

	return max
}

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)

	var left, right, max int
	for i := 0; i < k && i < n; i++ {
		left += cardPoints[i]
	}
	max = left
	if n <= k {
		return max
	}

	for i, j := k-1, n-1; i >= 0; i, j = i-1, j-1 {
		left -= cardPoints[i]
		right += cardPoints[j]

		if left+right > max {
			max = left + right
		}
	}

	return max
}

func search1(nums []int, target int) int {
	n, pos := len(nums), 0
	for i := 1; i < n; i++ {
		if nums[i-1] > nums[i] {
			pos = i
			break
		}
	}

	k := bsearch(nums, 0, pos-1, target)
	if k != -1 {
		return k
	}
	return bsearch(nums, pos, n-1, target)
}

func bsearch(nums []int, l, r, target int) int {
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[l] {
			//[l:mid]
			if nums[mid] < target {
				l = mid + 1
			} else {
				if target < nums[l] {
					l = mid + 1
				} else if target > nums[l] {
					r = mid - 1
				} else {
					return l
				}
			}
		} else {
			//[mid:r]
			if nums[mid] > target {
				r = mid - 1
			} else {
				if target > nums[r] {
					r = mid - 1
				} else if target < nums[r] {
					l = mid + 1
				} else {
					return r
				}
			}
		}
	}
	return -1
}

func rotate1(nums []int, k int) {
	n := len(nums)
	k = k % n

	for i := 0; i < k; i++ {
		m := nums[n-1]
		for j := 0; j < n; j++ {
			nums[j], m = m, nums[j]
		}
	}
}

func rotate2(nums []int, k int) {
	n := len(nums)
	k = k % n
	nums2 := make([]int, n)
	copy(nums2, nums)

	for i := 0; i < n-k; i++ {
		nums[k+i] = nums2[i]
	}

	for j := 0; j < k; j++ {
		nums[j] = nums2[n-k+j]
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}

	j, m, i := k, nums[0], 0
	for c := 0; c < n; c++ {
		nums[j], m = m, nums[j]
		if j == i {
			j++
			i++
			m = nums[j]
		}

		j = (j + k) % n
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	var arr []int
	arr = []int{1, 2, 2, 3, 4, 5, 4}
	fmt.Println(maxScore(arr, 3))
}
