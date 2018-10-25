package main

import "fmt"

// recursive method
func combinationSum_r(nums []int, target int) int {
	if target == 0 {
		return 1
	} else if target < 0 {
		return 0
	}
	tmp := 0
	for i := 0; i < len(nums); i++ {
		tmp += combinationSum_r(nums, target-nums[i])
	}
	return tmp
}

func combinationSum(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i < target+1; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(combinationSum_r(nums, 6))
	fmt.Println(combinationSum(nums, 6))
}
