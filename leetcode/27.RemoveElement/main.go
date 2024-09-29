package main

func removeElement(nums []int, val int) int {
	insertAt := 0

	for i, v := range nums {
		if v == val {
			continue
		}
		nums[insertAt] = nums[i]
		insertAt++
	}

	return insertAt
}
