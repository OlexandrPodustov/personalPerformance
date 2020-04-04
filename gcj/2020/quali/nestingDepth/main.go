package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var iterations int
	_, err := fmt.Scan(&iterations)
	if err != nil {
		fmt.Println("parse testCasesAmount:", err)
		return
	}

	// r := strings.Repeat("(", 3)
	// fmt.Println("repeat", r)

	for testCaseNumber := 1; testCaseNumber <= iterations; testCaseNumber++ {
		var stringOfDigits string
		_, err := fmt.Scan(&stringOfDigits)
		if err != nil {
			fmt.Println(testCaseNumber, "  stringOfDigits:", stringOfDigits, err)
			return
		}

		result := zeroOrOne(stringOfDigits)
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

func zeroOrOne(input string) string {
	fmt.Println("stringOfDigits:", input)

	ss := strings.Split(input, "")

	originalFinalPosition := make([][]string, len(ss))
	for i, v := range ss {
		originalFinalPosition[i] = []string{v}
	}

	fmt.Printf("originalFinalPosition %v\n", originalFinalPosition)

	result := ""
	for i := 0; i < len(ss); i++ {
		v := ss[i]
		if v == "0" {
			result += v

			old := originalFinalPosition[i][0]
			originalFinalPosition[i][0] += ", " + old
		} else {
			result += fmt.Sprintf("(%s)", v)

			old := originalFinalPosition[i][0]

			oldPosition, err := strconv.Atoi(old)
			if err != nil {
				fmt.Println("convert old position:", oldPosition, old, err)
				return ""
			}

			newPos := strconv.Itoa(1)
			originalFinalPosition[i][0] += ", " + newPos
		}
	}

	fmt.Printf("originalFinalPosition %v\n", originalFinalPosition)

	return result
}
