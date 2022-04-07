package main

import (
	"fmt"
)

func quicksort1(nums []int, low, high int) {
	if low >= high || low < 0 || high >= len(nums) {
		return
	}

	j := low
	for i := low; i < high; i++ {
		if nums[i] < nums[high] {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	nums[j], nums[high] = nums[high], nums[j]

	quicksort1(nums, low, j-1)
	quicksort1(nums, j+1, high)
}

func quicksort2(nums []int, low, high int) {
	if low >= high {
		return
	}

	start, end, key := low, high, nums[low]
	for start < end {
		for start < end && nums[end] >= key {
			end--
		}
		nums[start] = nums[end]

		for start < end && nums[start] <= key {
			start++
		}
		nums[end] = nums[start]
	}
	nums[start] = key

	quicksort2(nums, low, start-1)
	quicksort2(nums, start+1, high)
}

func quicksort3(arr []int, start, end int) {
	if start >= end {
		return
	}

	i, j := start, end
	key := arr[(start+end)/2]
	for i <= j {
		for arr[i] < key {
			i++
		}
		for arr[j] > key {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	quicksort3(arr, start, j)
	quicksort3(arr, i, end)
}

func partion(a []int, i, j int) int {
	for i < j {
		for i < j && a[i] < a[j] {
			j--
		}
		a[j], a[i] = a[i], a[j]
		for i < j && a[i] < a[j] {
			i++
		}
		a[j], a[i] = a[i], a[j]
	}
	return i
}

func quicksort4(a []int, left, right int) {
	if left >= right {
		return
	}
	mid := partion(a, left, right)
	quicksort4(a, left, mid-1)
	quicksort4(a, mid+1, right)
}

func heapsort(nums []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		_heapsort(nums, i, n)
	}

	for i := n - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		_heapsort(nums, 0, i)
	}
}

// 大顶堆
func _heapsort(nums []int, i, n int) {
	max := i
	left := 2*i + 1
	if left < n && nums[max] < nums[left] {
		max = left
	}

	right := left + 1
	if right < n && nums[max] < nums[right] {
		max = right
	}

	if max != i {
		nums[i], nums[max] = nums[max], nums[i]
		_heapsort(nums, max, n)
	}
}

func heapsort2(nums []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		_heapsort2(nums, i, n)
	}

	for i := n - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		_heapsort2(nums, 0, i)
	}
}

// 小顶堆
func _heapsort2(nums []int, i, n int) {
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
		_heapsort2(nums, min, n)
	}
}

func main() {
	var arr []int
	/*	arr = []int{5, 3, 4, 9, 2}
		quicksort1(arr, 0, len(arr)-1)
		fmt.Println(arr)

		arr = []int{5, 3, 4, 9, 2}
		quicksort2(arr, 0, len(arr)-1)
		fmt.Println(arr)

		arr = []int{5, 3, 4, 9, 2}
		quicksort3(arr, 0, len(arr)-1)
		fmt.Println(arr)

		arr = []int{5, 3, 4, 9, 2}
		quicksort4(arr, 0, len(arr)-1)
		fmt.Println(arr)*/

	arr = []int{5, 3, 4, 9, 2}
	heapsort(arr, len(arr))
	fmt.Println(arr)

	arr = []int{5, 3, 4, 9, 2}
	heapsort2(arr, len(arr))
	fmt.Println(arr)
}
