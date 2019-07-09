package main

import "fmt"

func search(nums []int, target int) int {

	pivot := len(nums) / 2
	left := 0
	right := len(nums) - 1
	for {
		fmt.Println(left, pivot, right)
		if nums[pivot] == target {
			return pivot
		} else if nums[pivot] != target && pivot == 0 {
			return -1
		} else if nums[left] <= target && target < nums[pivot] {
			right = pivot
			pivot = (right - left) / 2
		} else if nums[pivot] < target && target <= nums[right] {
			left = pivot
			pivot = (right-left)/2 + left
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
