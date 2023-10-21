package main

func findDuplicate(nums []int) int {
	dictionary := make(map[int]int)
	for _, v := range nums {
		if _, ok := dictionary[v]; ok {
			return v
		}
		dictionary[v]++
	}

	return 1
}
