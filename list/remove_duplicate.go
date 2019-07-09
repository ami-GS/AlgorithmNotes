package main

import "fmt"

func removeDuplicates(nums []int) int {

	out := len(nums)
	prvCount := 0
	for i := 1; i < out; {
		if nums[i-1] == nums[i] {
			prvCount++
		} else {
			ii := i
			i -= prvCount
			out -= prvCount
			for j := i; ii < out+prvCount; {
				nums[j] = nums[ii]
				j++
				ii++
			}
			prvCount = 0
		}
		i++
		for k := 0; k < out; k++ {
			fmt.Print(nums[k])
		}
		fmt.Println()
	}
	if prvCount != 0 {
		return out - prvCount
	}

	return out
}

func removeDuplicates2(nums []int) int {
	prvCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			prvCount++
		} else {
			nums[i-prvCount] = nums[i]
		}
	}
	return len(nums) - prvCount
}

func main() {
	nums := []int{1, 1, 2, 2, 3}
	length := removeDuplicates2(nums)
	for i := 0; i < length; i++ {
		fmt.Print(nums[i])
	}
}
