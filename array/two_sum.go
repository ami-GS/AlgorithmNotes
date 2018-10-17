package main

import "fmt"

func calc(arr []int, arrNum int, target int) [][2]int {
	// want Set, but go doesn't have
	tmpMap := make(map[int]int)
	var out [][2]int

	for i := 0; i < arrNum; i++ {
		_, ok := tmpMap[target-arr[i]]
		if ok {
			tmpMap[arr[i]] = 1
			out = append(out, [2]int{arr[i], target - arr[i]})
		} else {
			tmpMap[arr[i]] = 1
		}
	}
	return out
}

func main() {
	array := []int{2, 3, 4, 5, 6, 7, 9, 11, 15}
	out := calc(array, 9, 9)
	fmt.Println(out)
}
