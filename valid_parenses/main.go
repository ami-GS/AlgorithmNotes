package main

import "fmt"

func genParensis(leftLest, rightLest int) int {
	if leftLest == 0 && rightLest == 0 {
		return 1
	}
	ans := 0
	if leftLest > 0 {
		ans += genParensis(leftLest-1, rightLest+1)
	}
	if rightLest > 0 {
		ans += genParensis(leftLest, rightLest-1)
	}
	return ans
}

func main() {
	fmt.Println(genParensis(9, 0))
}
