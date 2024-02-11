package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	var sum int
	var ft func(root *TreeNode) int

	ft = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lt := ft(root.Left)
		rt := ft(root.Right)
		sum += abs(lt - rt)

		// fmt.Println("lt, rt, sum", lt, rt, sum)

		return lt + rt + root.Val
	}

	ft(root)

	return sum
}

func abs(in int) int {
	if in < 0 {
		return -in
	}
	return in
}
