package main

import "fmt"

// https://leetcode.com/explore/featured/card/top-interview-questions-easy/127/strings/882/

func isAnagram(s string, t string) bool {
	var cMap [26]int
	for _, c := range s {
		cMap[c-'a']++
	}
	for _, c := range t {
		cMap[c-'a']--
	}
	for _, v := range cMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
	fmt.Println(isAnagram("zlap", "kcqx"))
}
