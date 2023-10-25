package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := height(root, 1)
	return result
}

func height(tn *TreeNode, h int) int {
	if tn == nil {
		return 0
	}
	lh := height(tn.Left, h+1)
	rh := height(tn.Right, h+1)
	if lh > rh {
		return lh
	}

	return rh
}
