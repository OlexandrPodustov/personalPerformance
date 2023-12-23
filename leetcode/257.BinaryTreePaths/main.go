package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	result := []string{}
	solveBacktracking(root, strconv.Itoa(root.Val), &result)

	return result
}

func solve(root *TreeNode, path string, result *[]string) {
	if root == nil {
		return
	}

	element := ""
	if rv := strconv.Itoa(root.Val); rv == path {
		element = rv
	} else {
		element = fmt.Sprintf("%v->%v", path, root.Val)
	}
	fmt.Println("element", element)

	solve(root.Left, element, result)
	solve(root.Right, element, result)

	if root.Left == nil && root.Right == nil {
		*result = append(*result, element)
	}

	return
}

func solveBacktracking(root *TreeNode, path string, result *[]string) {
	if root == nil {
		return
	}
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprint(path))
	sb.WriteString(fmt.Sprint(`->`))
	sbLenBefore := sb.Len()
	if root.Left != nil {
		sb.WriteString(strconv.Itoa(root.Left.Val))
		lPath := sb.String()
		solveBacktracking(root.Left, lPath, result)
		sb.Reset()
		sb.Write([]byte(lPath[:sbLenBefore]))
	}

	if root.Right != nil {
		// sb.WriteByte(byte(root.Right.Val))
		sb.WriteString(strconv.Itoa(root.Right.Val))
		rPath := sb.String()
		solveBacktracking(root.Right, rPath, result)
		sb.Reset()
		sb.Write([]byte(rPath[:sbLenBefore]))
	}

	if root.Left == nil && root.Right == nil {
		*result = append(*result, path)
	}

	return
}

// for (direction) {
// visited.add(point)
// recursion(visited, blabla)
// visited.remove(point)
