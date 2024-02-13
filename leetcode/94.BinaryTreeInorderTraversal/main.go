package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	var traverse func(root *TreeNode) []int

	traverse = func(root *TreeNode) []int {
		if root == nil {
			return nil
		}

		traverse(root.Left)
		result = append(result, root.Val)
		traverse(root.Right)

		return result
	}
	traverse(root)

	return result
}
