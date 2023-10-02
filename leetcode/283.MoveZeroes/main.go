package main

func moveZeroes(nums []int) {
	ln := len(nums)
	if ln < 2 {
		return
	}

	j := 0
	for i := 0; i < ln && j < ln; i++ {
		if nums[i] != 0 {
			continue
		}
		for j = i + 1; j < ln; j++ {
			if nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
}
