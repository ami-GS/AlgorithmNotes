package main

import (
	"fmt"
	"math"
)

//https://leetcode.com/problems/validate-binary-search-tree/
/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBSTUtil(root *TreeNode, min int, max int) bool {

	if root.Val < min || max < root.Val {
		return false
	}

	//return isValidBSTUtil(root.Left, min, root.Val) && isValidBSTUtil(root.Right, root.Val, max)
	result := true
	if root.Left != nil {
		result = isValidBSTUtil(root.Left, min, root.Val-1) && result

	}
	if root.Right != nil {
		result = isValidBSTUtil(root.Right, root.Val-1, max) && result
	}
	return result

}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isValidBSTUtil(root, math.MinInt32, math.MaxInt32)
}

func main() {
	root := &TreeNode{10, nil, nil}
	root.Left = &TreeNode{5, nil, nil}
	root.Right = &TreeNode{15, nil, nil}
	root.Right.Left = &TreeNode{6, nil, nil}
	root.Right.Right = &TreeNode{20, nil, nil}

	fmt.Println(isValidBST(root))
}
