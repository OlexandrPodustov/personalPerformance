package main

func main() {
}

func runningSum(nums []int) []int {
	result := make([]int, len(nums))

	total := 0
	for i := range nums {
		total += nums[i]
		result[i] = total
	}

	return result
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	nodesToIterate := []*TreeNode{root}

	for len(nodesToIterate) != 0 {
		if nodesToIterate[0] != nil && nodesToIterate[0].Left != nil {
			nodesToIterate = append(nodesToIterate, nodesToIterate[0].Left)
		}
		if nodesToIterate[0] != nil && nodesToIterate[0].Right != nil {
			nodesToIterate = append(nodesToIterate, nodesToIterate[0].Right)
		}
		nodesToIterate[0].Left, nodesToIterate[0].Right = nodesToIterate[0].Right, nodesToIterate[0].Left

		nodesToIterate = nodesToIterate[1:]
	}

	return root
}

func maximumWealth(accounts [][]int) int {
	mv := 0

	for i := 0; i < len(accounts); i++ {
		localMax := 0
		for j := 0; j < len(accounts[i]); j++ {
			localMax += accounts[i][j]
		}
		if localMax > mv {
			mv = localMax
		}
	}

	return mv
}
