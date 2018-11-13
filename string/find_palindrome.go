package main

import "fmt"

func longestPalindrome(s string) string {
	evenIDs := make([]int, 0, 1000)
	oddIDs := make([]int, 0, 1000)
	if len(s) == 0 {
		return s
	}

	maxLen := 0
	maxSubStr := ""
	for i := 0; i < len(s)-1; i++ {
		if i != 0 && s[i-1] == s[i+1] {
			maxLen = 3
			maxSubStr = s[i-1 : i+2]
			evenIDs = append(evenIDs, i)
		}
		if s[i] == s[i+1] {
			if maxLen == 0 {
				maxLen = 2
				maxSubStr = s[i : i+2]
			}
			oddIDs = append(oddIDs, i)
		}
	}

	for _, idx := range evenIDs {
		for j := 2; 0 <= idx-j && idx+j < len(s); j++ {
			if s[idx-j] == s[idx+j] {
				if maxLen < 1+j*2 {
					maxLen = 1 + j*2
					maxSubStr = s[idx-j : idx+j+1]
				}
			} else {
				break
			}
		}
	}

	for _, idx := range oddIDs {
		for j := 1; 0 <= idx-j && idx+j+1 < len(s); j++ {
			if s[idx-j] == s[idx+1+j] {
				if maxLen < 2+j*2 {
					maxLen = 2 + j*2
					maxSubStr = s[idx-j : idx+1+j+1]
				}
			} else {
				break
			}
		}
	}
	if maxLen == 0 {
		return string(s[0])
	}
	return maxSubStr
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("aaaa"))
	fmt.Println(longestPalindrome("aa"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ac"))
	fmt.Println(longestPalindrome("babadkfjadlkjnbdafbannbbnnabnfnnbkdfnbkjdn"))
	fmt.Println(longestPalindrome("ababababa"))
	fmt.Println(longestPalindrome("aaaabbbbbbbbbbccccccccccddddddddddeeeeeeeeeeffffffffffgggggggggghhhhhhhhhhiiiiiiiiiijjjjjjjjjjkkkkkkkkkkllllllllllmmmmmmmmmmnnnnnnnnnnooooooooooppppppppppqqqqqqqqqqrrrrrrrrrrssssssssssttttttttttuuuuuuuuuuvvvvvvvvvvwwwwwwwwwwxxxxxxxxxxyyyyyyyyyyzzzzzzzzzzyyyyyyyyyyxxxxxxxxxxwwwwwwwwwwvvvvvvvvvvuuuuuuuuuuttttttttttssssssssssrrrrrrrrrrqqqqqqqqqqppppppppppoooooooooonnnnnnnnnnmmmmmmmmmmllllllllllkkkkkkkkkkjjjjjjjjjjiiiiiiiiiihhhhhhhhhhggggggggggffffffffffeeeeeeeeeeddddddddddccccccccccbbbbbbbbbbaaaaaaaabbbbbbbbbbccccccccccddddddddddeeeeeeeeeeffffffffffgggggggggghhhhhhhhhhiiiiiiiiiijjjjjjjjjjkkkkkkkkkkllllllllllmmmmmmmmmmnnnnnnnnnnooooooooooppppppppppqqqqqqqqqqrrrrrrrrrrssssssssssttttttttttuuuuuuuuuuvvvvvvvvvvwwwwwwwwwwxxxxxxxxxxyyyyyyyyyyzzzzzzzzzzyyyyyyyyyyxxxxxxxxxxwwwwwwwwwwvvvvvvvvvvuuuuuuuuuuttttttttttssssssssssrrrrrrrrrrqqqqqqqqqqppppppppppoooooooooonnnnnnnnnnmmmmmmmmmmllllllllllkkkkkkkkkkjjjjjjjjjjiiiiiiiiiihhhhhhhhhhggggggggggffffffffffeeeeeeeeeeddddddddddccccccccccbbbbbbbbbbaaaa"))
}
