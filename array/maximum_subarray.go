package main

import "fmt"

func calc(arr []int) int {
	max := 0
	out := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < max+arr[i] {
			max += arr[i]
		} else {
			max = arr[i]
		}
		if out < max {
			out = max
		}
	}
	return out
}

func main() {
	fmt.Println(calc([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
