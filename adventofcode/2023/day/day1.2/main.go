package main

import (
	"strings"
)

func solve(lines []string) int {
	result := 0

	for _, v := range lines {
		v = replaceWords(v)
		a := findFirstLast(v)
		result += a
	}

	return result
}

func findFirstLast(str string) int {
	numbers := make([]rune, 0)
	for _, v := range str {
		if v >= '0' && v <= '9' {
			numbers = append(numbers, v-'0')
		}
	}
	nn := 10*numbers[0] + numbers[len(numbers)-1]

	return int(nn)
}

const (
	one   = "one"
	two   = "two"
	three = "three"
	four  = "four"
	five  = "five"
	six   = "six"
	seven = "seven"
	eight = "eight"
	nine  = "nine"
)

func replaceWords(str string) string {
	rr := strings.NewReplacer(
		one, "o1e",
		two, "t2o",
		three, "t3e",
		four, "f4r",
		five, "f5e",
		six, "s6x",
		seven, "s7n",
		eight, "e8t",
		nine, "n9e",
	)

	return rr.Replace(rr.Replace(str))
}
