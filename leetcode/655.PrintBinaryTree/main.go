package main

import (
	"fmt"
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
	result := initMatrix(h+1, columnsAmount)

	fmt.Println("h", h)
	fillMatrix(root, result, h, columnsAmount)
	fmt.Println("result", result)

	return result
}

func fillMatrix(root *TreeNode, result [][]string, h, columnsAmount int) {
	r, c := 0, (columnsAmount-1)/2

	if root == nil {
		return
	}
	if r == 0 {
		result[r][c] = strconv.Itoa(root.Val) // root node placement
	}
	if root.Left != nil {
		result[r+1][max(c-1<<(h-r-1), 0)] = strconv.Itoa(root.Left.Val)
	}
	if root.Right != nil {
		result[r+1][c+1<<(h-r-1)] = strconv.Itoa(root.Right.Val)
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
	if root.Left == nil && root.Right == nil {
		return 0
	}

	var lh, rh int
	if root.Left != nil {
		lh = findHeight(root.Left) + 1
	}
	if root.Right != nil {
		rh = findHeight(root.Right) + 1
	}

	return max(lh, rh)
}
