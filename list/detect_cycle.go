package main

import "fmt"

type Node struct {
	val int
	nxt *Node
}

func hasCycle(root *Node) bool {
	if root == nil || root.nxt == nil {
		return false
	}

	slow := root
	fast := root.nxt
	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.nxt
		fast = fast.nxt.nxt
	}
	return false
}

func main() {
	root := &Node{0, nil}
	np := root
	for i := 0; i < 10; i++ {
		np.nxt = &Node{0, nil}
		np = np.nxt
	}
	np = root
	//np.nxt.nxt.nxt.nxt.nxt.nxt = np.nxt.nxt
	fmt.Println(hasCycle(root))
}
