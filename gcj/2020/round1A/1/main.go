package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
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
	distinctPatterns := make(map[string]struct{})
	patternByLength := make(map[int]string)
	for i := 0; i < numberOfPatternsToMatch; i++ {
		var pattern string
		_, err := fmt.Scan(&pattern)
		if err != nil {
			log.Println(numberOfPatternsToMatch, "  numberOfPatternsToMatch:", err)
			return err.Error()
		}
		pattern = pattern[1:]

		if _, ok := distinctPatterns[pattern]; !ok {
			distinctPatterns[pattern] = struct{}{}

			if _, okk := patternByLength[len(pattern)]; !okk {
				patternByLength[len(pattern)] = pattern
			} else {
				return "*"
			}
		}
	}
	// fmt.Println("distinctPatterns", distinctPatterns)
	// fmt.Println("patternByLength", patternByLength)

	diffLengthes := make([]int, 0)
	for k := range patternByLength {
		diffLengthes = append(diffLengthes, k)
	}
	sort.Ints(diffLengthes)
	// fmt.Println("sorted ints", diffLengthes)

	shouldMatch := patternByLength[diffLengthes[0]]
	diffLengthes = diffLengthes[1:]
	// fmt.Println("diffLengthes", diffLengthes)

	for _, minLen := range diffLengthes {
		strByLen := patternByLength[minLen]

		if strings.HasSuffix(strByLen, shouldMatch) {
			shouldMatch = strByLen
		} else {
			return "*"
		}
	}

	return shouldMatch
}

// 3
// 2
// *XZ
// *XZ
// 5
// *CONUTS
// *COCONUTS
// *OCONUTS
// *CONUTS
// *S
// 2
// *XZ
// *XYZ
