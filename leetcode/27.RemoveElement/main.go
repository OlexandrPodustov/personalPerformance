package main

func removeElement(nums []int, val int) int {
	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		}

		return 1
	}

	amountOfUnmatched := 0
	left := 0
	right := len(nums) - 1

	for left <= right {
		if nums[left] != val {
			amountOfUnmatched++
			left++
			continue
		}
		if nums[right] == val {
			right--
			continue
		}

		nums[left], nums[right] = nums[right], nums[left]
		left++
		amountOfUnmatched++
	}

	return amountOfUnmatched
}
