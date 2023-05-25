package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var testCases int
	if _, err := fmt.Scanln(&testCases); err != nil {
		log.Fatal("fmt.Scanln(&t) testCases err", err)
	}

	for caseNumber := 1; caseNumber <= testCases; caseNumber++ {
		var amountOfLetters int
		if _, err := fmt.Scanln(&amountOfLetters); err != nil {
			log.Fatal("fmt.Scanln - err: ", err)
		}

		// fmt.Println("amountOfLetters", amountOfLetters)

		var greetingString string
		if _, err := fmt.Scanln(&greetingString); err != nil {
			log.Fatal("fmt.Scanln - err: ", err)
		}
		res := solution(greetingString)
		fmt.Printf("Case #%d: %d\n", caseNumber, res)
	}
}

func solution(input string) int {
	// fmt.Println("input 1", input)

	rf := strings.NewReplacer("F", "")
	input = rf.Replace(input)
	// fmt.Println("input 2", input)

	xx := strings.NewReplacer("XX", "X")
	for strings.Contains(input, "XX") {
		input = xx.Replace(input)
		// fmt.Println("input 3", input)
	}
	// fmt.Println("input 4", input)

	oo := strings.NewReplacer("OO", "O")
	for strings.Contains(input, "OO") {
		input = oo.Replace(input)
		// fmt.Println("input 5", input)
	}
	// fmt.Println("input 6", input)

	result := len(input) - 1
	if result < 0 {
		result = 0
	}

	fmt.Println("result1", result)
	modulo := 1_000_000_007
	result %= modulo
	fmt.Println("result2", result)

	return result
}
