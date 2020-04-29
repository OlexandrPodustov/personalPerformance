package main

import (
	"fmt"
	"log"
)

func main() {
	var iterations int
	_, err := fmt.Scan(&iterations)
	if err != nil {
		log.Println("parse testCasesAmount:", err)
		return
	}

	for testCaseNumber := 1; testCaseNumber <= iterations; testCaseNumber++ {
		var numberOfPatternsToMatch int
		_, err := fmt.Scan(&numberOfPatternsToMatch)
		if err != nil {
			log.Println(testCaseNumber, "  numberOfPatternsToMatch:", err)
			return
		}

		result := check(numberOfPatternsToMatch)
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

func check(numberOfPatternsToMatch int) string {
	return ""
}
