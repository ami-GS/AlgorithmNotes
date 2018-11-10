package main

import "fmt"

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	max := 0
	subStart := 0
	subEnd := 1

	for i := 1; i < len(s); i++ {
		c := s[i]

		isRepeating := false
		for j := subStart; j < subEnd; j++ {
			if s[j] == c {
				isRepeating = true
				break
			}
		}

		if isRepeating {
			tmp := subEnd - subStart
			if max < tmp {
				max = tmp
			}
			subStart++
		}
		subEnd++
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("au"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
