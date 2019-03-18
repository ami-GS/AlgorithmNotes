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

func calcWithSet(arr []int, target int) [][2]int {
	sumSet := make(map[int]interface{})
	var out [][2]int

	for _, val := range arr {
		sub := target - val
		if _, ok := sumSet[sub]; ok {
			out = append(out, [2]int{sub, val})
		}
		sumSet[val] = struct{}{}
	}
	return out
}

func main() {
	array := []int{2, 3, 4, 5, 6, 7, 9, 11, 15}
	//out := calc(array, 9, 9)
	out := calcWithSet(array, 9)
	fmt.Println(out)
}
