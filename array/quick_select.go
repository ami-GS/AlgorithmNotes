package main

import "fmt"

// reference https://www.slideserve.com/deirdre-lane/quick-sort-quick-select
func QuickSelect(array []int, left, right, k int) int {
	pivot := (right-left)/2 + left
	pv := array[pivot]
	fmt.Println(pv, array, array[left:right+1])
	//array[pivot], array[right] = array[right], array[pivot]

	l := left
	r := right
	for {
		for ; array[l] < pv; l++ {
		}
		for ; pv < array[r]; r-- {
		}

		// POINT: this is important
		if l >= r {
			break
		}
		array[l], array[r] = array[r], array[l]
		l++
		r--
	}

	// POINT: careful about 'k' and 'l' indices difference
	// l == r means same sa k-1 == l
	// usually r+1 == l, due to loop avobe
	if k-1 == l {
		fmt.Println(pv, array, array[left:right+1])
		return array[l]
	} else if k-1 <= l {
		// POINT: careful aboit l-1 and r+1, or infinit loop
		// r+1 == l, that's why -1 or +1 are needed
		return QuickSelect(array, left, l-1, k)
	} else {
		return QuickSelect(array, r+1, right, k)
	}
	panic("bad index")
}

func main() {

	//fmt.Println(3+1, QuickSelect(array, 0, len(array)-1, 3+1))
	array := []int{3, 2, 1, 5, 6, 4}
	N := 6
	for i := 0; i < N; i++ {
		//array := []int{9, 8, 7, 6, 5, 0, 1, 2, 3, 4}
		array = []int{3, 2, 1, 5, 6, 4}
		//array = []int{6, 5, 4, 3, 2, 1}
		fmt.Println(i+1, QuickSelect(array, 0, len(array)-1, i+1))
	}
}
