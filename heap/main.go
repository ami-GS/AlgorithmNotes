package main

import (
	"fmt"

	"github.com/ami-GS/AlgorithmNotes/heap/go"
)

func main() {
	data := []int{2, 4, 5, 8, 9, 7, 14, 15, 1, 3, 10, 21, 13, 6, 11, 12}

	h := heap.NewFromSlice(data, heap.Max)
	for i := 0; i < len(data); i++ {
		fmt.Println(h.Pop())
	}
}
