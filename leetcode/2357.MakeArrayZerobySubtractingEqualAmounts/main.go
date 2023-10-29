package main

func minimumOperations(nums []int) int {
	distinctNumbers := make(map[int]struct{}, 0)
	for _, v := range nums {
		if v > 0 {
			distinctNumbers[v] = struct{}{}
		}
	}

	return len(distinctNumbers)
}
