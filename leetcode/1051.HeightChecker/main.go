package main

import (
	"sort"
)

func heightChecker(heights []int) int {
	res := 0

	hs := make([]int, len(heights))
	copy(hs, heights)
	sort.Ints(hs)

	for i, v := range heights {
		if v != hs[i] {
			res++
		}
	}

	return res
}
