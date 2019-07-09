package main

import "fmt"

func calc(arr []int) int {
	min := arr[0]
	max := arr[0]
	out := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < 0 {
			min, max = max, min
		}

		if tmp := max * arr[i]; tmp > arr[i] {
			max = tmp
		}
		if tmp := min * arr[i]; tmp < arr[i] {
			min = tmp
		}
		if out < max {
			out = max
		}
	}
	return out
}

func main() {
	out := calc([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(out)
}
