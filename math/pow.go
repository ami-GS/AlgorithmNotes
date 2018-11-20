package main

import "fmt"

//https://leetcode.com/problems/powx-n/
func pow(x float64, n int) float64 {
	if n == 1 {
		return x
	} else if n == 0 {
		return 1
	}

	if n%2 == 1 {
		return pow(x, n-1) * x
	}
	tmp := pow(x, n/2)
	return tmp * tmp

}

func myPow(x float64, n int) float64 {
	sign := 1
	if n < 0 {
		sign = -1
	}
	out := pow(x, sign*n)
	if n < 0 {
		return 1 / out
	}
	return out
}

func main() {
	fmt.Println(myPow(2.000, 10))
	fmt.Println(myPow(2.100, 3))
	fmt.Println(myPow(2.000, -2))
	fmt.Println(myPow(2.000, 0))

}
