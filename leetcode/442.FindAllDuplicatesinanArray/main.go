package main

import "sort"

func findDuplicates(nums []int) []int {
	var result []int
	dictionary := make(map[int]int)

	for _, v := range nums {
		dictionary[v]++
	}
	for key, v := range dictionary {
		if v > 1 {
			result = append(result, key)
		}
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}
