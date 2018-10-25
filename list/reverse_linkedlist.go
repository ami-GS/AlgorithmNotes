package main

import "fmt"

type Node struct {
	val int
	nxt *Node
}

func reverse(n *Node) *Node {
	bef := (*Node)(nil)
	tmp := n
	for tmp != nil {
		next := tmp.nxt
		tmp.nxt = bef
		//next.nxt = tmp
		bef = tmp
		tmp = next
	}
	return bef
}

func main() {
	N := 10
	n := Node{0, nil}
	np := &n
	for i := 1; i < N; i++ {
		np.nxt = &Node{i, nil}
		np = np.nxt
	}
	fmt.Println()
	np = &n
	for i := 0; i < N; i++ {
		fmt.Print(np.val, "->")
		np = np.nxt
	}
	fmt.Println()
	np = &n
	rnp := reverse(np)

	for i := 0; i < N; i++ {
		fmt.Print(rnp.val, "->")
		rnp = rnp.nxt
	}
}
