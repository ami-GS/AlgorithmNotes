package main

import "fmt"

func helper0(nums []int, mask []byte, sub []int, out *[][]int) {
	if len(sub) == len(nums) {
		tmp := make([]int, len(sub))
		copy(tmp, sub)
		*out = append(*out, tmp)
		return
	}

	for i, num := range nums {
		if mask[i] == 0 {
			mask[i] = 1
			helper0(nums, mask, append(sub, num), out)
			mask[i] = 0
		}
	}
}

func permute(nums []int) [][]int {
	out := make([][]int, 0)
	helper0(nums, make([]byte, len(nums)), make([]int, 0), &out)
	return out
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
