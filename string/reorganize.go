package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/reorganize-string/
type pair struct {
	ch    byte
	count int
}

func reorganizeString(S string) string {
	charCount := make([]pair, 'z'-'a'+1)
	for i, _ := range charCount {
		charCount[i].ch = byte('a' + i)
	}

	for _, c := range S {
		charCount[c-'a'].count++
	}
	// for _, v := range charCount {
	// 	fmt.Printf("%c:%d ", v.ch, v.count)
	// }
	sort.Slice(charCount, func(i, j int) bool {
		return charCount[i].count < charCount[j].count
	})
	// fmt.Println()
	// for _, v := range charCount {
	// 	fmt.Printf("%c:%d ", v.ch, v.count)
	// }
	out := make([]byte, len(S))
	outIdx := 0
	for i := len(charCount) - 1; i > 0; i-- {
		if charCount[i].count == 0 {
			return string(out)
		}
		out[outIdx] = charCount[i].ch
		outIdx++
		charCount[i].count--

		for k := i - 1; k >= 0; k-- {
			if charCount[k].count == 0 && charCount[i].count != 0 {
				return ""
			}
			for l := 0; l < charCount[k].count; l++ {
				out[outIdx] = charCount[k].ch
				outIdx++
				if charCount[i].count > 0 {
					out[outIdx] = charCount[i].ch
					outIdx++
					charCount[i].count--
				}
			}
			charCount[k].count = 0

			if charCount[i].count == 0 {
				i = k
				break
			}
		}
	}

	return ""
}

func do() string {
	return ""
}

func main() {
	testSet := []string{"aab", "aaab", "aba", "aabb", "zqugrfbsznyiwbokwkpvpmeyvaosdkedbgjogzdpwawwl"}

	for _, s := range testSet {
		fmt.Println(reorganizeString(s))
	}
}
