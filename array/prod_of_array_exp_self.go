package main

import "fmt"

func calc(nums []int) []int {
	out := make([]int, len(nums))
	for i, _ := range out {
		out[i] = 1
	}

	right := 1
	left := 1
	for i := 0; i < len(nums); i++ {
		out[i] *= left
		out[len(nums)-1-i] *= right
		left *= nums[i]
		right *= nums[len(nums)-1-i]
	}
	return out
}

func main() {
	fmt.Println(calc([]int{1, 2, 3, 4, 5, 6}))
}
