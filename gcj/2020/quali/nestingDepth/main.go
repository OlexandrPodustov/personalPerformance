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

	originalFinalPosition := make([]int, len(input))
	fmt.Printf("originalFinalPosition %v\n", originalFinalPosition)

	result := ""
	for i := 0; i < len(ss); i++ {
		v := ss[i]
		if v == "0" {
			result += v
		} else {
			result += fmt.Sprintf("(%s)", v)

			originalFinalPosition[i] = 1
		}
	}

	fmt.Printf("originalFinalPosition %v\n", originalFinalPosition)

	if len(input) > 1 {
		result = trimRedundand(input, result, originalFinalPosition)
	}

	return result
}

func trimRedundand(originalS, input string, ofp [][]string) string {
	result := ""

	var ossInt []int
	oss := strings.Split(originalS, "")
	for _, v := range oss {
		intVal, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("convert intVal:", v, err)
			return ""
		}

		ossInt = append(ossInt, intVal)
	}

	fmt.Println("ossInt:", ossInt)

	for i := 0; i < len(ossInt); i++ {
		current := ossInt[i]
		if i+1 >= len(ossInt) {
			break
		}

		next := ossInt[i+1]
		if current != 0 && next != 0 && current-next == 0 {
			padd := ofp[i]
			currentIndex := i

		}
	}

	return result
}
