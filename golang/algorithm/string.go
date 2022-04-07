package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func convert(s string, numRows int) string {
	i, k := 0, 0
	var cols [][]byte
	for {
		col := make([]byte, numRows)
		if k == 0 {
			for j := 0; j < numRows && i < len(s); i, j = i+1, j+1 {
				col[j] = s[i]
			}
		} else {
			col[numRows-1-k] = s[i]
			i++
		}

		cols = append(cols, col)
		if i == len(s) {
			break
		}

		k++
		if k >= numRows-1 {
			k = 0
		}
	}

	k = 0
	ns := make([]byte, len(s))
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(cols); j++ {
			if cols[j][i] != 0 {
				ns[k] = cols[j][i]
				k++
			}
		}
	}

	return string(ns)
}

func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	n := len(s)
	box := 2*numRows - 2
	cycle := n / box
	if n%box != 0 {
		cycle++
	}

	ns := make([]byte, n)
	i := 0
	for row := 0; row < numRows; row++ {
		for c := 0; c < cycle && row+box*c < n; c++ {
			ns[i], i = s[row+box*c], i+1

			if row > 0 && row < numRows-1 && box-row+box*c < n {
				ns[i], i = s[box-row+box*c], i+1
			}
		}
	}
	return string(ns)
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	neg := false
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		s = s[1:]
		neg = true
	}

	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}

		n = 10*n + int(s[i]-'0')
		if neg && -n < math.MinInt32 {
			return math.MinInt32
		} else if !neg && n > math.MaxInt32 {
			return math.MaxInt32
		}
	}

	if neg {
		return -n
	}
	return n
}

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func lengthOfLongestSubstring(s string) int {
	pos := func(ss []byte, c byte) int {
		for i := 0; i < len(ss); i++ {
			if ss[i] == c {
				return i
			}
		}
		return -1
	}

	var substr []byte
	max := 0
	for i := 0; i < len(s); i++ {
		if p := pos(substr, s[i]); p >= 0 {
			substr = substr[p+1:]
		}
		substr = append(substr, s[i])

		if len(substr) > max {
			max = len(substr)
		}
	}

	return max
}

func firstUniqChar1(s string) int {
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]] = mp[s[i]] + 1
	}
	for i := 0; i < len(s); i++ {
		n := mp[s[i]]
		if n == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar2(s string) int {
	arr := make([]int, 26)

	for i := 0; i < len(s); i++ {
		arr[int(s[i]-'a')]++
	}

	for i := 0; i < len(s); i++ {
		if arr[int(s[i]-'a')] == 1 {
			return i
		}
	}
	return -1
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	arr := make([]int, 26)
	for i := 0; i < len(s); i++ {
		arr[int(s[i]-'a')]++
	}

	for i := 0; i < len(t); i++ {
		arr[int(t[i]-'a')]--
	}

	for i := 0; i < 26; i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		ab, bb := byte(0), byte(0)
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			ab = s[i]
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			ab = s[i] - 'A' + 'a'
		} else {
			i++
			continue
		}

		if (s[j] >= 'a' && s[j] <= 'z') || (s[j] >= '0' && s[j] <= '9') {
			bb = s[j]
		} else if s[j] >= 'A' && s[j] <= 'Z' {
			bb = s[j] - 'A' + 'a'
		} else {
			j--
			continue
		}
		if ab != bb {
			return false
		}
		i++
		j--
	}

	return true
}

// 暴力算法
func strStr1(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	n := len(haystack)
	m := len(needle)

	for x := 0; x < n; x++ {
		if haystack[x] == needle[0] {
			if n-x < m {
				return -1
			}

			ok := true
			for i, j := x+1, 1; j < m; i, j = i+1, j+1 {
				if haystack[i] != needle[j] {
					ok = false
					break
				}
			}
			if ok {
				return x
			}
		}
	}

	return -1
}

// KMP
func KMP(haystack string, needle string) int {
	// next数组
	next := make([]int, len(needle))
	// 构建next数组的函数
	createNext := func() {
		// 当模式串长度为0，就直接返回
		if len(needle) == 0 {
			return
		}
		// 默认next[0]=-1
		k, index := -1, 0
		next[0] = -1
		// 使用模式串倒数第二个字符计算最后一个字符的next值，所以这里是<len(needle)-1
		for index < len(needle)-1 {
			// 当字符匹配，直接加1
			if k == -1 || needle[index] == needle[k] {
				index++
				k++
				next[index] = k
			} else {
				// 不匹配就更新k的值
				k = next[k]
			}
		}
	}
	// 模式匹配函数
	kmp := func() int {
		// 先处理特殊情况的空字符串
		if len(haystack) == 0 && len(needle) != 0 {
			return -1
		}
		if len(needle) == 0 {
			return 0
		}
		// 当连个字符串都不为空
		for h, n := 0, 0; ; {
			if haystack[h] == needle[n] {
				// 如果匹配，两个字符串的下标都加1
				n++
				h++
				// 是否到了模式串的最后
				if n == len(needle) {
					return h - n
				}
				// 是否到了被匹配字符串你的最后
				if h == len(haystack) {
					return -1
				}
			} else {
				// 遇到不匹配的字符
				n = next[n]
				// 从第一个字符就不匹配，那么就要将连个字符串都向后移动一个位置，因为只有next[0]=-1
				// 其他最小的是next的值是0，表示需要对模式串重新从最开始匹配
				if n < 0 {
					n++
					h++
					if h == len(haystack) {
						return -1
					}
				}
			}
		}
	}
	createNext()
	return kmp()
}

// Boyer-Moore Horspool Sunday
// Rabin Karp
func charToInt(idx int, s string) int {
	return (int)(s[idx] - 'a')
}

func RabinKarp(haystack string, needle string) int {
	L := len(needle)
	n := len(haystack)
	if L > n {
		return -1
	}

	// base value for the rolling hash function
	a := 26
	// modulus value for the rolling hash function to avoid overflow
	modulus := int(math.Pow(2, 31))

	// compute the hash of strings haystack[:L], needle[:L]
	h := 0
	ref_h := 0
	for i := 0; i < L; i++ {
		h = (h*a + charToInt(i, haystack)) % modulus
		ref_h = (ref_h*a + charToInt(i, needle)) % modulus
	}
	if h == ref_h {
		return 0
	}

	// const value to be used often : a**L % modulus
	aL := 1
	for i := 1; i <= L; i++ {
		aL = (aL * a) % modulus
	}

	for start := 1; start < n-L+1; start++ {
		// compute rolling hash in O(1) time
		h = (h*a - charToInt(start-1, haystack)*aL + charToInt(start+L-1, haystack)) % modulus
		if h == ref_h {
			return start
		}
	}
	return -1
}

// primeRK is the prime base used in Rabin-Karp algorithm.
const primeRK = 16777619

// hashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func hashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

func indexRabinKarp(s, substr string) int {
	// Rabin-Karp search
	hashss, pow := hashStr(substr)
	n := len(substr)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	if h == hashss && s[:n] == substr {
		return 0
	}
	for i := n; i < len(s); {
		h *= primeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashss && s[i-n:i] == substr {
			return i - n
		}
	}
	return -1
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}
	if len(haystack) < len(needle) {
		return -1
	}

	return indexRabinKarp(haystack, needle)
}

func countAndSay(n int) string {
	if n <= 1 {
		return "1"
	}

	s := countAndSay(n - 1)

	c := 1
	sb := bytes.NewBuffer(nil)
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			sb.WriteString(strconv.Itoa(c))
			sb.WriteByte(s[i-1])
			c = 1
		} else {
			c++
		}
	}
	sb.WriteString(strconv.Itoa(c))
	sb.WriteByte(s[len(s)-1])

	return sb.String()
}

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}

	i := 0
	for {
		if len(strs[0]) <= i {
			return strs[0][0:i]
		}

		for j := 1; j < n; j++ {
			if len(strs[j]) <= i {
				return strs[0][0:i]
			}
			if strs[j][i] != strs[0][i] {
				return strs[0][0:i]
			}
		}
		i++
	}

	return ""
}

func main() {
	fmt.Println(longestCommonPrefix(nil))
}
