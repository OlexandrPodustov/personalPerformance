package main

import "fmt"

func main() {
	// in := []int{-3, 2, -3, 4, 2}
	// in := []int{1, -2, -3}
	in := []int{1, 2}

	res := minStartValue(in)
	fmt.Println(res)
}

func minStartValue(nums []int) int {
	minSum := 0
	sum := 0

	for _, n := range nums {
		sum += n
		if sum < minSum {
			minSum = sum
		}

	}
	// fmt.Println(minSum)
	x := 1 - minSum

	return x
}
