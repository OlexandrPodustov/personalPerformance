package main

func removeDuplicates(nums []int) int {
	var j int
	for i := range nums {
		if nums[j] == nums[i] {
			continue
		}

		j++
		nums[j] = nums[i]
	}

	return j + 1
}
