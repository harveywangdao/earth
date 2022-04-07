package main

import (
	"fmt"
	"math"
	//"math/bits"
	"sort"
	"strconv"
)

func smallestGoodBase1(n string) string {
	m, _ := strconv.Atoi(n)
	for k := 2; k < m; k++ {
		base, sum := 1, 0
		for {
			sum += base
			base *= k

			if sum == m {
				return strconv.Itoa(k)
			} else if sum > m {
				break
			}
		}
	}
	return ""
}

func smallestGoodBase2(n string) string {
	m, _ := strconv.Atoi(n)
	for k := 2; k < m; k++ {
		a := (m-1)*(k-1) + k
		for a >= k {
			if a == k {
				return strconv.Itoa(k)
			}

			if a%k == 0 {
				a /= k
			} else {
				break
			}
		}
	}
	return ""
}

func smallestGoodBase(n string) string {
	num, _ := strconv.ParseUint(n, 10, 64)
	for bit := uint64(math.Log2(float64(num))); bit >= 1; bit-- {
		low, high := uint64(2), uint64(math.Pow(float64(num), 1.0/float64(bit)))
		for low <= high {
			fmt.Println(low, high, bit)
			mid := uint64(low + (high-low)>>1)
			sum := findBase(mid, bit)
			if sum == num {
				return strconv.FormatUint(mid, 10)
			} else if sum > num {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return strconv.FormatUint(num-1, 10)
}

func findBase(mid, bit uint64) uint64 {
	sum, base := uint64(1), uint64(1)
	for i := uint64(1); i <= bit; i++ {
		base *= mid
		sum += base
	}
	return sum
}

func hammingWeight1(num uint32) int {
	var n int
	for num > 0 {
		if num&1 == 1 {
			n++
		}
		num >>= 1
	}
	return n
}

func hammingWeight2(num uint32) int {
	mask, n := uint32(1), 0
	for i := 0; i < 32; i++ {
		if num&mask != 0 {
			n++
		}
		mask <<= 1
	}
	return n
}

func hammingWeight3(num uint32) int {
	n := 0
	for num > 0 {
		n++
		num &= (num - 1)
	}
	return n
}

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}

func isPowerOfTwo2(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(-n) == n
}

func isPowerOfTwo3(n int) bool {
	if n <= 0 {
		return false
	}
	return (1024*1024*1024)%n == 0
}

func findDisappearedNumbers1(nums []int) []int {
	var ns []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			find := false
			for j := 0; j < len(nums); j++ {
				if nums[j] == i+1 {
					nums[i], nums[j] = nums[j], nums[i]
					find = true
					break
				}
			}

			if !find {
				ns = append(ns, i+1)
			}
		}
	}

	return ns
}

func findDisappearedNumbers2(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var ns []int
	i := 1
	for _, n := range nums {
		if n == i {
			i++
			continue
		} else if n > i {
			for j := i; j < n; j++ {
				ns = append(ns, j)
			}
			i = n + 1
		}
	}

	for j := i; j <= len(nums); j++ {
		ns = append(ns, j)
	}

	return ns
}

func findDisappearedNumbers3(nums []int) []int {
	var ns []int
	m := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = true
	}

	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			ns = append(ns, i)
		}
	}

	return ns
}

func findDisappearedNumbers4(nums []int) []int {
	var ns []int
	arr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		arr[nums[i]-1] = nums[i]
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			ns = append(ns, i+1)
		}
	}

	return ns
}

func findDisappearedNumbers5(nums []int) []int {
	i := 0
	for i < len(nums) {
		if nums[i] == i+1 {
			i++
			continue
		}

		if nums[i] == nums[nums[i]-1] {
			i++
			continue
		}

		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		fmt.Println(nums)
	}

	var ns []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ns = append(ns, i+1)
		}
	}

	return ns
}

func findDisappearedNumbers6(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n < 0 {
			n = -n
		}

		if nums[n-1] > 0 {
			nums[n-1] = -nums[n-1]
		}
	}

	var ns []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ns = append(ns, i+1)
		}
	}

	return ns
}

func findErrorNums1(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	lost, cpy := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			cpy = nums[i]
			break
		}
	}

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n < 0 {
			n = -n
		}

		if nums[n-1] > 0 {
			nums[n-1] = -nums[n-1]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			lost = i + 1
			break
		}
	}

	return []int{cpy, lost}
}

func findErrorNums2(nums []int) []int {
	cpy, lost := 0, 0
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n < 0 {
			n = -n
		}
		if nums[n-1] > 0 {
			nums[n-1] = -nums[n-1]
		} else {
			cpy = n
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			lost = i + 1
			break
		}
	}

	return []int{cpy, lost}
}

func findErrorNums(nums []int) []int {
	xor := 0
	for i := 0; i < len(nums); i++ {
		xor = xor ^ nums[i]
		xor = xor ^ (i + 1)
	}
	//rightmostbit := xor & ^(xor - 1)
	rightmostbit := xor & (-xor)

	xor0, xor1 := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&rightmostbit != 0 {
			xor1 = xor1 ^ nums[i]
		} else {
			xor0 = xor0 ^ nums[i]
		}

		if (i+1)&rightmostbit != 0 {
			xor1 = xor1 ^ (i + 1)
		} else {
			xor0 = xor0 ^ (i + 1)
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == xor0 {
			return []int{xor0, xor1}
		}
	}
	return []int{xor1, xor0}
}

func multiply(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	mul := make([]byte, len(num1)+len(num2))
	res := make([][]byte, len(num2))
	for i := 0; i < len(res); i++ {
		res[i] = make([]byte, len(num1)+len(num2))
	}

	for i := len(num2) - 1; i >= 0; i-- {
		k := len(res[i]) - 1

		for j := 0; j <= k; j++ {
			res[i][j] = '0'
		}

		for j := 0; j < len(num2)-1-i; j++ {
			k--
		}

		m := byte(0)
		b := num2[i] - '0'
		for j := len(num1) - 1; j >= 0; j-- {
			a := num1[j] - '0'
			c := a*b + m

			res[i][k] = c%10 + '0'
			k--
			m = c / 10
		}

		if m != 0 {
			res[i][k] = m + '0'
		}
	}

	/*	for i := 0; i < len(res); i++ {
		fmt.Println(string(res[i]))
	}*/

	sum := 0
	for i := len(num1) + len(num2) - 1; i >= 0; i-- {
		for j := 0; j < len(num2); j++ {
			sum += int(res[j][i] - '0')
		}
		mul[i] = byte(sum%10) + '0'
		sum = sum / 10
	}

	zero := 0
	for zero < len(mul)-1 {
		if mul[zero] != '0' {
			break
		}
		zero++
	}

	return string(mul[zero:])
}

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	res := make([]byte, len(num1)+1)

	i, j := len(num1)-1, len(num2)-1
	sum := 0
	for k := len(res) - 1; k >= 0; k-- {
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		res[k] = byte(sum%10) + '0'
		sum /= 10
	}

	zero := 0
	for zero < len(res)-1 {
		if res[zero] != '0' {
			break
		}
		zero++
	}
	return string(res[zero:])
}

var memo map[int]int = map[int]int{}

func numSquares1(n int) int {
	if m, ok := memo[n]; ok {
		return m
	}

	fmt.Println(n)

	sq := int(math.Sqrt(float64(n)))
	if sq*sq == n {
		memo[n] = 1
		return 1
	}

	min := math.MaxInt32
	for i := sq; i >= 1; i-- {
		m := 1 + numSquares1(n-i*i)
		if m < min {
			min = m
		}
	}
	memo[n] = min
	return min
}

func numSquares(n int) int {
	m := n
	for m%4 == 0 {
		m /= 4
	}
	if m%8 == 7 {
		return 4
	}

	q := int(math.Sqrt(float64(n)))
	if q*q == n {
		return 1
	}

	for i := 1; i*i <= n; i++ {
		q := int(math.Sqrt(float64(n - i*i)))
		if q*q == n-i*i {
			return 2
		}
	}
	return 3
}

var sum int

func sumNums1(n int) int {
	sum = 0
	_sumNums1(n)
	return sum
}

func _sumNums1(n int) bool {
	sum += n
	return n > 1 && _sumNums1(n-1)
}

func fm1(a, b int) int {
	m := 0
	for i := 0; b>>i > 0; i++ {
		if (b>>i)&1 == 1 {
			m += (a << i)
		}
	}
	return m
}

func fm(a, b int) int {
	m := 0
	for ; b != 0; b >>= 1 {
		if b&1 == 1 {
			m += a
		}
		a <<= 1
	}
	return m
}

func sumNums2(n int) int {
	a := n + 1
	b := n
	m := 0

	var fn func(x, y int) bool
	fn = func(x, y int) bool {
		m += y
		return x&1 == 1 || fn(1, -y)
	}

	fn(b, a)
	b >>= 1
	a <<= 1

	fn(b, a)
	b >>= 1
	a <<= 1

	fn(b, a)
	b >>= 1
	a <<= 1

	fn(b, a)
	b >>= 1
	a <<= 1

	fn(b, a)
	b >>= 1
	a <<= 1

	fn(b, a)
	b >>= 1
	a <<= 1

	fn(b, a)
	b >>= 1
	a <<= 1

	fn(b, a)
	b >>= 1
	a <<= 1

	fn(b, a)
	b >>= 1
	a <<= 1

	fn(b, a)
	b >>= 1
	a <<= 1

	fn(b, a)
	b >>= 1
	a <<= 1

	fn(b, a)
	b >>= 1
	a <<= 1

	fn(b, a)
	b >>= 1
	a <<= 1

	fn(b, a)

	return m >> 1
}

func sumNums(n int) int {
	a := n + 1
	b := n
	m := 0

	fn := func() bool {
		m += a
		return true
	}

	_ = (b&1 == 1) && fn()
	b >>= 1
	a <<= 1

	_ = (b&1 == 1) && fn()
	b >>= 1
	a <<= 1

	_ = (b&1 == 1) && fn()
	b >>= 1
	a <<= 1

	_ = (b&1 == 1) && fn()
	b >>= 1
	a <<= 1

	_ = (b&1 == 1) && fn()
	b >>= 1
	a <<= 1

	_ = (b&1 == 1) && fn()
	b >>= 1
	a <<= 1

	_ = (b&1 == 1) && fn()
	b >>= 1
	a <<= 1

	_ = (b&1 == 1) && fn()
	b >>= 1
	a <<= 1

	_ = (b&1 == 1) && fn()
	b >>= 1
	a <<= 1

	_ = (b&1 == 1) && fn()
	b >>= 1
	a <<= 1

	_ = (b&1 == 1) && fn()
	b >>= 1
	a <<= 1

	_ = (b&1 == 1) && fn()
	b >>= 1
	a <<= 1

	_ = (b&1 == 1) && fn()
	b >>= 1
	a <<= 1

	_ = (b&1 == 1) && fn()

	return m >> 1
}

func countBits1(num int) []int {
	res := make([]int, num+1)
	for i := 0; i <= num; i++ {
		n := 0
		for j := i; j > 0; j >>= 1 {
			if j&1 == 1 {
				n++
			}
		}
		res[i] = n
	}
	return res
}

func countBits2(num int) []int {
	res := make([]int, num+1)
	for i := 0; i <= num; i++ {
		n := 0
		for j := i; j > 0; j &= (j - 1) {
			n++
		}
		res[i] = n
	}
	return res
}

func countBits3(num int) []int {
	res := make([]int, num+1)
	level := 2
	m := 0
	for i := 1; i <= num; i++ {
		if i == level {
			m <<= 1
			m++
			level *= 2
		}

		res[i] = 1 + res[i&m]
	}
	return res
}

func countBits4(num int) []int {
	res := make([]int, num+1)
	m := 1
	for i := 1; i <= num; i++ {
		if i == m*2 {
			m *= 2
		}
		res[i] = 1 + res[i-m]
	}
	return res
}

func countBits5(num int) []int {
	res := make([]int, num+1)
	m := 0
	for i := 1; i <= num; i++ {
		if i&(i-1) == 0 {
			m = i
		}
		res[i] = 1 + res[i-m]
	}
	return res
}

func countBits6(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = i&1 + res[i>>1]
	}
	return res
}

func countBits7(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = 1 + res[i&(i-1)]
	}
	return res
}

func main() {

}
