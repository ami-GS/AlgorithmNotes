package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func kthSmallest(root *Node, k int) int {

	if root.left != nil {
		return kthSmallest(root.left, k)
	}
	if root.right != nil {
		return kthSmallest(root.right, k)
	}
	return 1
}

func makeBST(root *Node, nodes []*Node) {
	for i := 0; i < len(nodes); i++ {
		p := root
		n := nodes[i]
		for {
			if p.val <= n.val {
				if p.right == nil {
					p.right = n
					break
				} else {
					p = p.right
				}
			} else {
				if p.left == nil {
					p.left = n
					break
				} else {
					p = p.left
				}
			}
		}
	}
}

func printTree(root *Node) {
	if root == nil {
		return
	}
	printTree(root.left)
	printTree(root.right)
	fmt.Println(root.val)
}

func main() {
	root := &Node{1, nil, nil}
	nodes := []*Node{&Node{8, nil, nil}, &Node{9, nil, nil}, &Node{5, nil, nil}, &Node{6, nil, nil}, &Node{20, nil, nil}, &Node{3, nil, nil}}
	makeBST(root, nodes)
	printTree(root)
	//fmt.Println(kthSmallest(root, 3))
}
