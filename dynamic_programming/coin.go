package main

import "fmt"

func coinChange(coins []int, target int) int {
	dp := make([]int, target+1)
	for i, _ := range dp {
		dp[i] = 1 << 21
	}
	dp[0] = 0
	fmt.Println(dp)
	for i := 1; i < target+1; i++ {
		for _, coin := range coins {
			//fmt.Println(coin)
			if i >= coin {
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	fmt.Println(dp)
	if dp[target] == 1<<21 {
		return -1
	}
	return dp[target]
}

func main() {
	coins := []int{1, 2, 5}
	target := 11
	fmt.Println(coinChange(coins, target))
}
