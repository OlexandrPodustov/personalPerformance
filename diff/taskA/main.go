package main

import "fmt"

func main() {
	input := []int{5, 15, 43, 9, 2}
	res := check(input, 7)
	fmt.Println(res)
}

func check(input []int, target int) []int {
	var result []int

	inputMap := make(map[int]int, len(input))
	for i, v := range input {
		inputMap[v] = i

		diff := target - v
		index, ok := inputMap[diff]
		if ok {
			result = []int{i, index}
			break
		}
	}

	return result
}
