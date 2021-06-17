package main

import "fmt"

func main() {
	// input := []int{3, -6, 4, -3, 7, 4}
	input := []int{0, 0}
	// input := []int{0}
	// input := []int{0, 1}
	res := check(input)
	fmt.Println(res)
}

func check(input []int) bool {
	if len(input) < 2 {
		return false
	}

	m := make(map[int]int, len(input))
	for _, v := range input {
		m[v]++
	}

	for _, v := range input {
		vv := -1 * v
		mv, ok := m[vv]

		if v == 0 && mv == 1 {
			continue
		}
		if ok {
			return true
		}
	}

	return false
}
