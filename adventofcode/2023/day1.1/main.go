package main

import (
	"fmt"
	"strconv"
)

func solve(lines []string) int {
	result := 0

	for _, v := range lines {
		a := findFirstLast(v)
		result += a

	}

	return result
}

func findFirstLast(str string) int {
	fmt.Println("raw", str)

	numbers := make([]rune, 0)
	for _, v := range str {
		if v >= '0' && v <= '9' {
			numbers = append(numbers, v-'0')
		}
	}

	fmt.Println("numbers", numbers)

	nn := fmt.Sprintf("%v%v", numbers[0], numbers[len(numbers)-1])
	nnn, _ := strconv.Atoi(nn)

	fmt.Println("numbers", nn, nnn)

	return nnn
}
