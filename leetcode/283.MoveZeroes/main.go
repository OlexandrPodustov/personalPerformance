package main

func moveZeroes(nums []int) {
	ln := len(nums)
	if ln < 2 {
		return
	}

	for i := 0; i < ln; i++ {
		if nums[i] == 0 {
			for j := i; j < ln; j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}
