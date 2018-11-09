package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/explore/featured/card/top-interview-questions-easy/127/strings/883/
func isPalindrome(s string) bool {
	ss := strings.ToLower(s)
	i := 0
	j := len(s) - 1
	for i < j {
		for i < len(ss)-1 {
			if 'a' <= ss[i] && ss[i] <= 'z' {
				break
			}
			i++
		}
		for 0 < j {
			if 'a' <= ss[j] && ss[j] <= 'z' {
				break
			}
			j--
		}
		if i == j {
			return true
		}

		if ss[i] != ss[j] {
			return false
		}
		j--
		i++
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(".,"))
}
