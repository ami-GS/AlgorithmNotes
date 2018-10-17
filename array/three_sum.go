package main

import (
	"fmt"
	"sort"
)

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

}
