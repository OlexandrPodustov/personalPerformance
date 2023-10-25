package main

func containsDuplicate(nums []int) bool {
	distinctNums := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := distinctNums[v]; ok {
			return true
		}
		distinctNums[v] = struct{}{}
	}

	return false
}
