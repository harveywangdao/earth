package main

import (
	"fmt"
)

func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	n1 := len(nums1)
	n2 := len(nums2)
	res := make([]int, n1)
	for i := 0; i < n1; i++ {
		res[i] = -1
		for j := 0; j < n2; j++ {
			if nums1[i] == nums2[j] {
				for x := j + 1; x < n2; x++ {
					if nums1[i] < nums2[x] {
						res[i] = nums2[x]
						break
					}
				}
				break
			}
		}
	}
	return res
}

func nextGreaterElement2(nums1 []int, nums2 []int) []int {
	n1 := len(nums1)
	n2 := len(nums2)
	res := make([]int, n1)

	mp := make(map[int]int)
	for i := 0; i < n2; i++ {
		mp[nums2[i]] = i
	}

	for i := 0; i < n1; i++ {
		res[i] = -1
		for j := mp[nums1[i]] + 1; j < n2; j++ {
			if nums1[i] < nums2[j] {
				res[i] = nums2[j]
				break
			}
		}
	}

	return res
}

func nextGreaterElement3(nums1 []int, nums2 []int) []int {
	n1 := len(nums1)
	n2 := len(nums2)
	mp := make(map[int]int)
	s := NewArrayStack(n2)
	for i := 0; i < n2; i++ {
		for s.Len() > 0 && s.Peek() < nums2[i] {
			mp[s.Pop()] = nums2[i]
		}
		s.Push(nums2[i])
	}

	for i := 0; i < n1; i++ {
		if x, ok := mp[nums1[i]]; ok {
			nums1[i] = x
		} else {
			nums1[i] = -1
		}
	}

	return nums1
}

func gcd1(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func gcd(a, b int) int {
	if b%a == 0 {
		return a
	}
	return gcd(b%a, a)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func containsDuplicate(nums []int) bool {
	mp := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := mp[nums[i]]; ok {
			return true
		}
		mp[nums[i]] = true
	}
	return false
}

func singleNumber(nums []int) int {
	m := 0
	for i := 0; i < len(nums); i++ {
		m ^= nums[i]
	}
	return m
}

func intersection(nums1 []int, nums2 []int) []int {
	n1 := len(nums1)
	n2 := len(nums2)
	mp := make(map[int]int)

	for i := 0; i < n1; i++ {
		mp[nums1[i]] = 0
	}

	var res []int
	for i := 0; i < n2; i++ {
		if v, ok := mp[nums2[i]]; ok && v == 0 {
			mp[nums2[i]] = 1
			res = append(res, nums2[i])
		}
	}
	return res
}

func intersect(nums1 []int, nums2 []int) []int {
	n1 := len(nums1)
	n2 := len(nums2)
	mp := make(map[int]int)

	for i := 0; i < n1; i++ {
		mp[nums1[i]] = mp[nums1[i]] + 1
	}

	var res []int
	for i := 0; i < n2; i++ {
		if v := mp[nums2[i]]; v > 0 {
			mp[nums2[i]] = v - 1
			res = append(res, nums2[i])
		}
	}
	return res
}

func plusOne(digits []int) []int {
	n := len(digits)

	m := 1
	for i := n - 1; i >= 0; i-- {
		if m == 0 {
			break
		}
		x := digits[i] + m
		digits[i] = x % 10
		m = x / 10
	}

	if m == 1 {
		res := make([]int, n+1)
		res[0] = 1
		copy(res[1:], digits)
		return res
	}

	return digits
}

func isValidSudoku1(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		mp := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			n := board[i][j] - '0'
			if _, ok := mp[n]; ok {
				return false
			}
			mp[n] = true
		}
	}

	for i := 0; i < 9; i++ {
		mp := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[j][i] == '.' {
				continue
			}
			n := board[j][i] - '0'
			if _, ok := mp[n]; ok {
				return false
			}
			mp[n] = true
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			mp := make(map[byte]bool)
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					if board[x+i*3][y+3*j] == '.' {
						continue
					}
					n := board[x+i*3][y+3*j] - '0'
					if _, ok := mp[n]; ok {
						return false
					}
					mp[n] = true
				}
			}
		}
	}
	return true
}

func isValidSudoku2(board [][]byte) bool {
	m_row := make([]map[byte]bool, 9)
	m_col := make([]map[byte]bool, 9)
	m_box := make([]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		m_row[i] = make(map[byte]bool)
		m_col[i] = make(map[byte]bool)
		m_box[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			n := board[i][j] - '0'
			if _, ok := m_row[i][n]; ok {
				return false
			}
			m_row[i][n] = true

			if _, ok := m_col[j][n]; ok {
				return false
			}
			m_col[j][n] = true

			index := (i/3)*3 + j/3
			if _, ok := m_box[index][n]; ok {
				return false
			}
			m_box[index][n] = true
		}
	}
	return true
}

func isValidSudoku3(board [][]byte) bool {
	var m_row, m_col, m_box int
	mm := make([]int, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			n := board[i][j] - '1'
			m_row = 1 << n
			m_col = 1 << 9 << n
			m_box = 1 << 9 << 9 << n

			if mm[i]&m_row != 0 {
				return false
			}
			mm[i] |= m_row

			if mm[j]&m_col != 0 {
				return false
			}
			mm[j] |= m_col

			index := (i/3)*3 + j/3
			if mm[index]&m_box != 0 {
				return false
			}
			mm[index] |= m_box
		}
	}
	return true
}

func rotate1(matrix [][]int) {
	n := len(matrix)

	matrix2 := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix2[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix2[j][n-1-i] = matrix[i][j]
		}
	}
	copy(matrix, matrix2)
}

func rotate2(matrix [][]int) {
	n := len(matrix)

	for cirle := 0; cirle < n/2; cirle++ {
		m := n - 2*cirle - 1

		for x := 0; x < m; x++ {
			i, j := cirle, cirle+x
			pre := matrix[i][j]

			for y := 0; y < 4; y++ {
				i, j = j, n-1-i
				matrix[i][j], pre = pre, matrix[i][j]
			}
		}
	}
}

func dailyTemperatures(T []int) []int {
	s := NewArrayStack(len(T))
	for i := 0; i < len(T); i++ {
		for s.Len() > 0 && s.Peek()%30000 < T[i] {
			t := s.Pop()
			T[t/30000] = i - t/30000
		}
		s.Push(i*30000 + T[i])
		T[i] = 0
	}
	return T
}

func firstBadVersion1(n int) int {
	arr := make([]int, n+1)
	l, r := 1, n
	for {
		mid := (l + r) / 2

		_badVersion(arr, mid)

		if arr[mid] == 1 {
			if mid == n {
				return mid + 1
			} else {
				_badVersion(arr, mid+1)

				if arr[mid+1] == 1 {
					l = mid + 1
				} else {
					return mid + 1
				}
			}
		} else {
			if mid == 1 {
				return mid
			} else {
				_badVersion(arr, mid-1)

				if arr[mid-1] == 1 {
					return mid
				} else {
					r = mid - 1
					continue
				}
			}
		}
	}
	return 0
}

func _badVersion1(arr []int, x int) {
	if arr[x] == 0 {
		if isBadVersion(x) {
			arr[x] = -1
		} else {
			arr[x] = 1
		}
	}
}

func _badVersion(x int) int {
	if isBadVersion(x) {
		return -1
	}
	return 1
}

func firstBadVersion2(n int) int {
	l, r := 1, n
	for {
		mid := (l + r) / 2

		mid_b := _badVersion(mid)

		if mid_b == 1 {
			if mid == n {
				return mid + 1
			} else {
				mid_r := _badVersion(mid + 1)

				if mid_r == 1 {
					l = mid + 1
				} else {
					return mid + 1
				}
			}
		} else {
			if mid == 1 {
				return mid
			} else {
				mid_l := _badVersion(mid - 1)

				if mid_l == 1 {
					return mid
				} else {
					r = mid - 1
					continue
				}
			}
		}
	}
	return 0
}

func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		mid := (l + r) / 2
		if !isBadVersion(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {
	fmt.Println(firstBadVersion(5))
}
