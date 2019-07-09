package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal2(root *TreeNode, out *[]int) {
	if root.Left != nil {
		inorderTraversal2(root.Left, out)
	}
	*out = append(*out, root.Val)
	if root.Right != nil {
		inorderTraversal2(root.Right, out)
	}
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	out := make([]int, 0, 100)
	inorderTraversal2(root, &out)
	return out
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.Right = &TreeNode{2, nil, nil}
	root.Right.Left = &TreeNode{3, nil, nil}

	out := inorderTraversal(root)
}
