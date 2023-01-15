package main

import (
	"sort"
)

func findSumOfThree(nums []int, target int) bool {
	sort.Sort(sort.IntSlice(nums))

	low := 0
	high := len(nums) - 1

	for i := 0; i < len(nums)-2; i++ {
		low = i + 1

		for low < high {
			total := nums[i] + nums[low] + nums[high]

			switch {
			case total == target:
				return true
			case total > target:
				high--
			case total < target:
				low++
			}
		}
	}

	return false
}
