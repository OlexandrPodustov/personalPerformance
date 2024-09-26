package main

func moveZeroes(nums []int) {
	insertAt := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[insertAt] = nums[i]
			insertAt++
		}
	}
	for i := insertAt; i < len(nums); i++ {
		nums[i] = 0
	}
}
