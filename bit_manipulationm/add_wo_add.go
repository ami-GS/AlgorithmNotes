package main

import "fmt"

// temporal values are important
func add(a, b int) int {
	out := 0
	out = a ^ b
	b = (a & b) << 1
	a = out
	for b != 0 {
		out = a ^ b
		b = (a & b) << 1
		a = out
	}
	return out
}

func add2(left, right int) int {
	for right != 0 {
		left, right = left^right, (left&right)<<1
	}
	return left
}

func main() {
	fmt.Println(add2(1, 2))
	fmt.Println(add2(2, 1))
	fmt.Println(add2(1, 99))
	fmt.Println(add2(99, 1))
	fmt.Println(add2(11, 0))
	fmt.Println(add2(0, 11))
	fmt.Println(add2(0xff, 0x1))
}
