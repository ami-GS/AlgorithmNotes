package main

import (
	"fmt"
	"math"
)

func myAtoi2(str string) int {
	numStarting := false
	out := int(0)
	sign := false
	for _, c := range str {
		switch {
		case c == ' ':
			if numStarting {
				break
			}
		case c == '-':
			sign = true
		case c == '+':
			continue
		case '0' <= c && c <= '9':
			numStarting = true
			out = out*10 + int(c-'0')
		default:
			if !numStarting {
				return 0
			} else {
				break
			}
		}
	}
	if out > 2147483648 && sign {
		return -2147483648
	}
	if out > 2147483647 {
		return 2147483647
	}
	if sign {
		return -out
	}
	return out
}

// https://leetcode.com/problems/string-to-integer-atoi/
func myAtoi(str string) int {
	sign := 1
	oneSign := false
	val := int64(0)
	digit := 0
	for idx, _ := range str {
		c := str[len(str)-1-idx]
		isNum := '0' <= c && c <= '9'
		if c == '-' {
			if oneSign {
				return 0
			}
			oneSign = true
			sign = -1
		} else if c == '+' {
			if oneSign {
				return 0
			}
			oneSign = true
		} else if c == '.' {
			val = 0
			digit = 0
		} else if digit != 0 && !isNum {
			if c != ' ' {
				return 0
			}
			//break
		} else if isNum {
			val += int64(c-'0') * int64(math.Pow(10, float64(digit)))
			digit++
		}
	}
	if val != 0 {
		out := int(val) * sign
		if out > math.MaxInt32 {
			return math.MaxInt32
		} else if out < math.MinInt32 {
			return math.MinInt32
		}
		return out
	}
	return 0
}

func main() {
	fmt.Println(myAtoi(""))
	fmt.Println(myAtoi("+1"))
	fmt.Println(myAtoi("+-1"))
	fmt.Println(myAtoi("++1"))
	fmt.Println(myAtoi("--1"))
	fmt.Println(myAtoi("-+1"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("    42"))
	fmt.Println(myAtoi("42   "))
	fmt.Println(myAtoi("-42   "))
	fmt.Println(myAtoi("  42  "))
	fmt.Println(myAtoi("  -42  "))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("91283472332"))
	fmt.Println(myAtoi("3.1415"))
	fmt.Println(myAtoi("a 3.1415"))
	fmt.Println(myAtoi("3.1415 a"))
}
