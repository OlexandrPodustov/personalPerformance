package main

func moveZeroes(nums []int) {
	ln := len(nums)
	if ln < 2 {
		return
	}

	lastNonZeroFoundAt := 0
	for i := 0; i < ln; i++ {
		if nums[i] != 0 {
			nums[lastNonZeroFoundAt] = nums[i]
			lastNonZeroFoundAt++
		}
	}

	for i := lastNonZeroFoundAt; i < ln; i++ {
		nums[i] = 0
	}
}
