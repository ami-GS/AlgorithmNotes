package main

import "fmt"

func calc(nums []int) int {
	min := nums[0]
	subMax := 0
	for _, v := range nums {
		if v < min {
			min = v
		}
		if subMax < v-min {
			subMax = v - min
		}
	}
	return subMax
}

func main() {
	//nums := []int{45, 24, 35, 31, 40, 38, 11}
	nums := []int{44, 30, 24, 32, 35, 30, 40, 38, 15}
	fmt.Println(calc(nums))
}
