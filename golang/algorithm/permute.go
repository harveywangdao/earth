package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var combination [][]int
	var line []int

	if len(nums) == 1 {
		line = append(line, nums[0])
		combination = append(combination, line)
		return combination
	} else if len(nums) == 2 {
		line = make([]int, 0, len(nums))
		line = append(line, nums[0], nums[1])
		combination = append(combination, line)

		line = make([]int, 0, len(nums))
		line = append(line, nums[1], nums[0])
		combination = append(combination, line)
		return combination
	}

	for i := 0; i < len(nums); i++ {
		line = make([]int, 0, len(nums)-1)
		line = append(line, nums[:i]...)
		line = append(line, nums[i+1:]...)
		lines := permute(line)
		for j := 0; j < len(lines); j++ {
			lines[j] = append([]int{nums[i]}, lines[j]...)
		}

		combination = append(combination, lines...)
	}
	return combination
}

func contains(arr []int, n int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			return true
		}
	}
	return false
}

var result [][]int

func permute2(nums []int) [][]int {
	_permute2(nums, nil)
	return result
}

func _permute2(nums, track []int) {
	if len(nums) == len(track) {
		line := make([]int, len(track))
		copy(line, track)
		result = append(result, line)
		return
	}

	for i := 0; i < len(nums); i++ {
		if contains(track, nums[i]) {
			continue
		}
		track = append(track, nums[i])
		_permute2(nums, track)
		track = track[:len(track)-1]
	}
}

func permute3(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	used := make([]bool, len(nums))
	path := make([]int, 0, len(nums))

	_permute3(nums, path, used, &res)
	return res
}

func _permute3(nums []int, path []int, used []bool, res *[][]int) {
	if len(path) == len(nums) {
		line := make([]int, len(path))
		copy(line, path)
		*res = append(*res, line)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		_permute3(nums, path, used, res)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func main() {
	ss := permute([]int{1, 2, 3, 4})
	for i := 0; i < len(ss); i++ {
		fmt.Println(ss[i])
	}

	ss2 := permute2([]int{1, 2, 3, 4})
	for i := 0; i < len(ss2); i++ {
		fmt.Println(ss2[i])
	}

	ss3 := permute3([]int{1, 2, 3, 4})
	for i := 0; i < len(ss3); i++ {
		fmt.Println(ss3[i])
	}
}
