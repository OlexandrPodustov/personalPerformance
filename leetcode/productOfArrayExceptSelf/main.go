package main

import "fmt"

func main() {
	in := []int{1, 2, 3, 4}
	result := productExceptSelf(in)
	fmt.Println("final result", result)
}

// complexity O(n^2)
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	for i := range nums {
		res := 1
		for j := range nums {
			if i != j {
				res *= nums[j]
			}
		}

		result[i] = res
		// fmt.Println(i, "result", result[i])
	}

	return result
}
