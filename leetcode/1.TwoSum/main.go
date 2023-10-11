package main

func twoSum(nums []int, target int) []int {
	diffs := make(map[int]int)

	result := []int{}
	for i, v := range nums {
		if val, ok := diffs[v]; ok {
			result = []int{val, i}
			break
		}
		diffs[target-v] = i
	}

	return result
}
