package main

import (
	"math"
)

func thirdMax(nums []int) int {
	const minInt = math.MinInt
	var (
		max  = minInt
		max2 = minInt
		max3 = minInt
	)

	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	for _, v := range nums {
		if v < max && v > max2 {
			max2 = v
		}
	}

	for _, v := range nums {
		if v < max2 && v > max3 {
			max3 = v
		}
	}

	if max3 == minInt {
		return max
	}

	return max3
}
