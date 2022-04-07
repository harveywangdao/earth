package main

import (
	"fmt"
	"time"
)

var count int

// 暴力枚举法
func longestPalindrome(s string) string {
	count = 0

	h, t, max := 0, 0, 0
	for i := 0; i < len(s)-1-max; i++ {
		for j := len(s) - 1; j > i && j-i > max; j-- {
			if s[i] == s[j] {
				ok := true
				for head, tail := i+1, j-1; head < tail; head, tail = head+1, tail-1 {
					count++
					if s[head] != s[tail] {
						ok = false
						break
					}
				}

				if ok {
					h, t = i, j
					max = t - h
					break
				}
			}
		}
	}

	return s[h : t+1]
}

// 自创 记录位置法
func longestPalindrome2(s string) string {
	count = 0

	table := make(map[byte][]int)
	for i := len(s) - 1; i >= 0; i-- {
		table[s[i]] = append(table[s[i]], i)
	}

	h, t, max := 0, 0, 0
	for i := 0; i < len(s)-1-max; i++ {
		tt := table[s[i]]
		for pos := 0; pos < len(tt); pos++ {
			count++

			if tt[pos]-i <= max {
				break
			}

			ok := true
			for head, tail := i+1, tt[pos]-1; head < tail; head, tail = head+1, tail-1 {
				if s[head] != s[tail] {
					ok = false
					break
				}
			}

			if ok {
				h, t = i, tt[pos]
				max = t - h
				break
			}
		}
	}

	return s[h : t+1]
}

// 中心扩散法
func longestPalindrome3(s string) string {
	count = 0
	h, max := 0, 1
	for i := 0; i < len(s); i++ {
		l, r := i, i+1
		for ; l >= 0 && r < len(s); l, r = l-1, r+1 {
			count++
			if s[l] != s[r] {
				break
			}
		}

		if r-l-1 > max {
			h, max = l+1, r-l-1
		}

		l, r = i, i
		for ; l >= 0 && r < len(s); l, r = l-1, r+1 {
			count++
			if s[l] != s[r] {
				break
			}
		}

		if r-l-1 > max {
			h, max = l+1, r-l-1
		}
	}

	return s[h : h+max]
}

// 动态规划
func longestPalindrome4(s string) string {
	count = 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	head, max := 0, 1

	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			count++
			if i == j {
				dp[i][j] = true
			} else if j-i < 3 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
			}

			if dp[i][j] && j-i+1 > max {
				head, max = i, j-i+1
			}
		}
	}

	return s[head : head+max]
}

// Manacher算法
func longestPalindrome5(s string) string {
	start, end := 0, -1
	t := "#"
	for i := 0; i < len(s); i++ {
		t += string(s[i]) + "#"
	}
	t += "#"
	s = t
	arm_len := []int{}
	right, j := -1, -1
	for i := 0; i < len(s); i++ {
		var cur_arm_len int
		if right >= i {
			i_sym := j*2 - i
			min_arm_len := min(arm_len[i_sym], right-i)
			cur_arm_len = expand(s, i-min_arm_len, i+min_arm_len)
		} else {
			cur_arm_len = expand(s, i, i)
		}
		arm_len = append(arm_len, cur_arm_len)
		if i+cur_arm_len > right {
			j = i
			right = i + cur_arm_len
		}
		if cur_arm_len*2+1 > end-start {
			start = i - cur_arm_len
			end = i + cur_arm_len
		}
	}
	ans := ""
	for i := start; i <= end; i++ {
		if s[i] != '#' {
			ans += string(s[i])
		}
	}
	return ans
}

func expand(s string, left, right int) int {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return (right - left - 2) / 2
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Manacher算法
func longestPalindrome7(s string) string {
	count = 0
	ns := make([]byte, 2*len(s)+1)
	for i := 0; i < len(s); i++ {
		ns[2*i], ns[2*i+1] = '#', s[i]
	}
	ns[len(ns)-1] = '#'
	p := make([]int, len(ns))

	maxRight, center, maxLen, begin := 0, 0, 1, 0
	for i := 0; i < len(ns); i++ {
		if i < maxRight {
			mirror := 2*center - i
			p[i] = min(maxRight-i, p[mirror])
		}

		left, right := i-(1+p[i]), i+(1+p[i])
		for left >= 0 && right < len(ns) && ns[left] == ns[right] {
			count++
			p[i]++
			left--
			right++
		}

		if i+p[i] > maxRight {
			maxRight = i + p[i]
			center = i
		}

		if p[i] > maxLen {
			maxLen = p[i]
			begin = (i - maxLen) / 2
		}
	}
	return s[begin : begin+maxLen]
}

// 优化的中心扩散法
func longestPalindrome6(s string) string {
	var index = 0
	var max = 0
	var size = len(s)
	var res string
	// 中心点扩散检测，分为两种中心点，一种是各个字母为中心点，
	// 一种是各个字母之间作为中心点，例如5个字母的字符串，一共有5+4个中心点
	for index < size {
		left := index
		right := index
		// 左右相同的情况
		for left >= 0 && s[left] == s[index] { // 如果中心点左边和中心点相同，则向左边扩散检测
			left--
		}
		// 左右相同的情况 比如 cbbd, 其中的bb
		for right < size && s[right] == s[index] { // 如果中心点右边和中心点相同，则向右边扩散
			right++
		}
		next := right                                          // 下一个点向右移动
		for left >= 0 && right < size && s[left] == s[right] { // 如果左右相等，则向两边扩散
			left--
			right++
		}
		index = next
		midMaxLen := right - left + 1
		if midMaxLen > max {
			res = s[left+1 : right]
			max = midMaxLen
		}
	}
	return res
}

func main() {
	var now time.Time
	now = time.Now()
	fmt.Println(longestPalindrome7("ansdsjkask"), "cost:", time.Now().Sub(now), "count:", count)

	now = time.Now()
	fmt.Println(longestPalindrome7("abba"), "cost:", time.Now().Sub(now), "count:", count)

	now = time.Now()
	fmt.Println(longestPalindrome7("abcba"), "cost:", time.Now().Sub(now), "count:", count)

	now = time.Now()
	fmt.Println(longestPalindrome7("abcbaeabcdefdfedcba"), "cost:", time.Now().Sub(now), "count:", count)

	now = time.Now()
	fmt.Println(longestPalindrome7("asdf"), "cost:", time.Now().Sub(now), "count:", count)

	now = time.Now()
	fmt.Println(longestPalindrome7("asdhhf"), "cost:", time.Now().Sub(now), "count:", count)
}
