package main

import "fmt"

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

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(2, 1))
	fmt.Println(add(1, 99))
	fmt.Println(add(99, 1))
	fmt.Println(add(11, 0))
	fmt.Println(add(0, 11))
	fmt.Println(add(0xff, 0x1))
}
