package main

import "math"

func maxArea(height []int) int {
	result := math.MinInt

	start, end := 0, len(height)-1
	for start < end {
		area := min(height[start], height[end]) * (end - start)
		if area > result {
			result = area
		}

		if height[start] <= height[end] {
			start++
		} else {
			end--
		}
	}

	return result
}
