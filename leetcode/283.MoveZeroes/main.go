package main

func moveZeroes(nums []int) {
	ln := len(nums)
	if ln < 2 {
		return
	}

	lastNonZeroFoundAt := 0
	for cur := 0; cur < ln; cur++ {
		if nums[cur] != 0 {
			nums[lastNonZeroFoundAt], nums[cur] = nums[cur], nums[lastNonZeroFoundAt]
			lastNonZeroFoundAt++
		}
	}
}
