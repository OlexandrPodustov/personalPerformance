package main

import (
	"strconv"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	h := findHeight(root)
	columnsAmount := 1<<h - 1
	result := initMatrix(h, columnsAmount)

	fillMatrix(root, result, h-1, 0, columnsAmount/2)

	return result
}

func fillMatrix(root *TreeNode, result [][]string, height, row, col int) {
	if root == nil {
		return
	}
	result[row][col] = strconv.Itoa(root.Val)

	if root.Left != nil {
		fillMatrix(root.Left, result, height, row+1, col-1<<(max(height-row-1, 0)))
	}
	if root.Right != nil {
		fillMatrix(root.Right, result, height, row+1, col+1<<(max(height-row-1, 0)))
	}
}

func initMatrix(h, columnsAmount int) [][]string {
	result := make([][]string, 0, h)
	for i := 0; i < h; i++ {
		res := make([]string, 0, columnsAmount)
		for i := 0; i < columnsAmount; i++ {
			res = append(res, "")
		}
		result = append(result, res)
	}
	return result
}

func findHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(findHeight(root.Left), findHeight(root.Right)) + 1
}
