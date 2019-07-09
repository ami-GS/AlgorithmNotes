package main

import "fmt"

func combination(arr []byte, subArr []byte, ret *[][]byte, depth int, maxDepth int) {
	if depth == maxDepth {
		subRet := make([]byte, maxDepth)
		copy(subRet, subArr)
		*ret = append(*ret, subRet)
		return
	}
	for i := 0; i < len(arr); i++ {
		subArr[depth] = arr[i]
		combination(arr[i+1:], subArr, ret, depth+1, maxDepth)
	}

}

func permutation(arr []byte, subArr []byte, ret *[][]byte, depth int) {
	if depth == len(arr) {
		subRet := make([]byte, len(subArr))
		copy(subRet, subArr)
		*ret = append(*ret, subRet)
		return
	}

	for i := 0; i < len(arr); i++ {
		subArr[depth] = arr[i]
		permutation(arr, subArr, ret, depth+1)
	}
}

func helper(candidates []int, target int, sub []int, out *[][]int) {
	if target == 0 {
		tmp := make([]int, len(sub))
		copy(tmp, sub)
		*out = append(*out, tmp)
	} else if target < 0 {
		return
	}
	for i, num := range candidates {
		if target >= num {
			helper(candidates[i:], target-num, append(sub, num), out)
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	out := make([][]int, 0)
	helper(candidates, target, make([]int, 0), &out)
	return out
}

func main() {
	//arr := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	// arr := []byte{'a', 'b', 'c'}
	// out := [][]byte{}
	// permutation(arr, make([]byte, len(arr)), &out, 0)
	//target := 2
	//combination(arr, make([]byte, target), &out, 0, target)
	out := combinationSum([]int{7, 3, 2}, 18)
	fmt.Println(out)
}
