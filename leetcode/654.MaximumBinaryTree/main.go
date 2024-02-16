package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// nums := []int{3, 2, 1}
	nums := []int{3, 2, 1, 6, 0, 5}
	constructMaximumBinaryTree(nums)
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mi := findMaxNumberIndex(nums)
	left := constructMaximumBinaryTree(nums[:mi])

	var right *TreeNode
	if mi < len(nums) {
		right = constructMaximumBinaryTree(nums[mi+1:])
	}

	return &TreeNode{
		Val:   nums[mi],
		Left:  left,
		Right: right,
	}
}

func findMaxNumberIndex(nums []int) int {
	max := -1
	maxIndex := -1

	for i := range nums {
		if nums[i] > max {
			max = nums[i]
			maxIndex = i
		}
	}

	return maxIndex
}
