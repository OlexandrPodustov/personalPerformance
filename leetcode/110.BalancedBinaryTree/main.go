package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lh := findHeight(root.Left)
	rh := findHeight(root.Right)

	// fmt.Println("main", lh, rh, abs(lh-rh))

	if abs(lh-rh) > 1 {
		return false
	}

	return true
}

func findHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	height := 1
	lh, rh := 0, 0
	if root.Left != nil {
		lh = findHeight(root.Left)
	}
	if root.Right != nil {
		rh = findHeight(root.Right)
	}

	if lh > rh {
		height += lh
	} else {
		height += rh
	}
	// fmt.Println("findHeight", root.Val, height, lh, rh)

	return height
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}
