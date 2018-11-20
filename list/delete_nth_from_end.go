package main

import "fmt"

//https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := head
	prevP := head
	fastP := head
	for i := 0; i < n; i++ {
		fastP = fastP.Next
	}

	for fastP != nil {
		prevP = p
		p = p.Next
		fastP = fastP.Next
	}
	if p == head {
		head = head.Next
	} else {
		prevP.Next = p.Next
	}
	return head
}

func main() {
	head := ListNode{1, nil}
	//head.Next = &ListNode{2, nil}
	// head.Next.Next = &ListNode{3, nil}
	// head.Next.Next.Next = &ListNode{4, nil}
	// head.Next.Next.Next.Next = &ListNode{5, nil}
	// head.Next.Next.Next.Next.Next = &ListNode{6, nil}
	// head.Next.Next.Next.Next.Next.Next = &ListNode{7, nil}
	p := &head
	for p != nil {
		fmt.Printf("%d->", p.Val)
		p = p.Next
	}

	fmt.Println()
	after := removeNthFromEnd(&head, 1)
	p = after
	for p != nil {
		fmt.Printf("%d->", p.Val)
		p = p.Next
	}
}
