package main

import "fmt"

func run(step int) int {
	if step == 0 {
		return 1
	}
	ans := 0
	if step >= 2 {
		ans += run(step - 2)
	}
	if step >= 1 {
		ans += run(step - 1)
	}
	return ans
}

func main() {
	fmt.Println(run(6))
}
