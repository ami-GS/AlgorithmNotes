package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/3sum/
func threeSum(arr []int) [][]int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	outMap := make(map[[3]int]struct{})
	for left := 0; left < len(arr); left++ {
		mid := left + 1
		for right := len(arr) - 1; mid < right; right-- {
			lrSum := arr[left] + arr[right]
			for mid < right {
				sum := lrSum + arr[mid]
				fmt.Println(len(arr), arr, left, mid, right, sum)
				if sum > 0 {
					break
				} else if sum < 0 {
					mid++
				} else {
					//out = append(out, []int{arr[left], arr[mid], arr[right]})
					outMap[[3]int{arr[left], arr[mid], arr[right]}] = struct{}{}
					mid++
				}
			}
		}
	}
	out := make([][]int, len(outMap))
	i := 0
	for k, _ := range outMap {
		out[i] = []int{k[0], k[1], k[2]}
		i++
	}
	return out
}

type myArr []int

func (p *myArr) Len() int {
	return len(*p)
}
func (p *myArr) Less(i, j int) bool {
	return (*p)[i] < (*p)[j]
}
func (p *myArr) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func calc(arr myArr, target int) [][3]int {
	out := make([][3]int, 0)
	sort.Sort(&arr)
	fmt.Println(arr)
	for i := 0; i < arr.Len(); i++ {
		k := arr.Len() - 1
		j := i + 1
		for ; j < k; k-- {
			tmp := arr[i] + arr[k]
			for {
				tmptmp := tmp + arr[j]
				if tmptmp > target {
					break
				} else if tmptmp < target {
					j++
				} else {
					out = append(out, [3]int{arr[i], arr[j], arr[k]})
					j++
				}
			}
		}
	}
	return out
}

func main() {
	arr := []int{2, 3, 4, -1, 5, -2, 6, -1, 0, 8}
	out := calc(arr, 10)
	fmt.Println(out)
	arr2 := []int{-1, 0, 1, 2, -1, -4}
	out2 := threeSum(arr2)
	fmt.Println(out2)
	out3 := threeSum(arr)
	fmt.Println(out3)

}
