package main

func removeElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[count] = nums[i]
			count++
		}
	}

	return count
}