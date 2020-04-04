package main

import (
	"fmt"
)

func main() {
	var iterations int
	_, err := fmt.Scan(&iterations)
	if err != nil {
		fmt.Println("parse testCasesAmount:", err)
		return
	}

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
