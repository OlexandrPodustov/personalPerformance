package main

import "fmt"

func main() {
	ints := []int{1, 2, 3}
	fmt.Println("sum", ints, sum(ints))
	ints = []int{2, 4, 6, 14, 3, 7}
	fmt.Println("sum", ints, sum(ints))
	fmt.Println("countNumbers", ints, countNumbers(ints))
	fmt.Println("maxNumber", ints, maxNumber(ints))
}

func sum(ints []int) int {
	if len(ints) == 0 {
		return 0
	}

	return ints[0] + sum(ints[1:])
}

func countNumbers(ints []int) int {
	if len(ints) == 0 {
		return 0
	}

	return 1 + countNumbers(ints[1:])
}

func maxNumber(ints []int) int {
	if len(ints) == 2 {
		if ints[0] > ints[1] {
			return ints[0]
		}
		return ints[1]
	}

	subMax := maxNumber(ints[1:])
	if ints[0] > subMax {
		return ints[0]
	}
	return subMax
}
