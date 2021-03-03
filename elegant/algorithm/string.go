package main

import (
	"fmt"
	"math"
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

func main() {
	fmt.Println(convert2("PAYPALISHIRING", 3))
	fmt.Println(myAtoi("4193 with words"))
}
