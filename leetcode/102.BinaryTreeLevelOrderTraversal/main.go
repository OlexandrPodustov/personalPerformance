package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	result := make([][]int, 0)
	for len(queue) > 0 {
		currentLevel := []int{}
		queueLen := len(queue)
		for i := 0; i < queueLen; i++ {
			currentElement := queue[0]
			// fmt.Println("currentElement inside", currentElement)
			queue = queue[1:]
			if currentElement.Left != nil {
				queue = append(queue, currentElement.Left)
			}
			if currentElement.Right != nil {
				queue = append(queue, currentElement.Right)
			}
			currentLevel = append(currentLevel, currentElement.Val)
		}

		result = append(result, currentLevel)
	}

	return result
}
