package main

import "fmt"

func swap(lv, rv *int) {
	if *lv != *rv {
		*lv ^= *rv
		*rv ^= *lv
		*lv ^= *rv
	}
}

func main() {
	a, b := 2, 4
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}
