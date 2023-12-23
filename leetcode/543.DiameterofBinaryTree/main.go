package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var maxmax int
	maxHeight(root, &maxmax)

	return maxmax
}

func maxHeight(tree *TreeNode, maxmax *int) int {
	if tree == nil {
		return 0
	}
	ml := maxHeight(tree.Left, maxmax)
	mr := maxHeight(tree.Right, maxmax)

	*maxmax = max(*maxmax, ml+mr)

	return max(ml, mr) + 1
}
