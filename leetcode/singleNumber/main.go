package main

import "fmt"

func main() {
	// fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}

func singleNumber(nums []int) int {
	numnum := make(map[int]int, len(nums))
	for i := range nums {
		numnum[nums[i]]++
	}

	for k, v := range numnum {
		if v == 1 {
			return k
		}
	}

	return 0
}
