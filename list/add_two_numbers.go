package main

import "fmt"

// https://leetcode.com/problems/add-two-numbers/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	out := &ListNode{0, nil}
	outp := out
	befp := outp.Next
	for {
		if l1 == nil && l2 == nil {
			if outp.Val == 0 {
				befp.Next = nil
			}
			break
		}

		l1v := 0
		if l1 != nil {
			l1v = l1.Val
		}
		l2v := 0
		if l2 != nil {
			l2v = l2.Val
		}

		tmp := l1v + l2v + outp.Val
		outp.Val = tmp % 10
		outp.Next = &ListNode{tmp / 10, nil}
		befp = outp
		outp = outp.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return out
}

func main() {
	l1 := ListNode{2, nil}
	l1.Next = &ListNode{4, nil}
	l1.Next.Next = &ListNode{3, nil}
	l2 := ListNode{5, nil}
	l2.Next = &ListNode{6, nil}
	l2.Next.Next = &ListNode{4, nil}

	out := addTwoNumbers(&l1, &l2)
	outp := out
	for outp != nil {
		fmt.Printf("%d -> ", outp.Val)
		outp = outp.Next
	}

}
