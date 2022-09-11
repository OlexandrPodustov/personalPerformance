package main

import (
	"fmt"
	"log"
)

func main() {
	var testCases int
	if _, err := fmt.Scanln(&testCases); err != nil {
		log.Fatal("fmt.Scanln(&t) testCases err", err)
	}

	for caseNumber := 1; caseNumber <= testCases; caseNumber++ {
		var preOwnedClockPartsAmnt, capable_of_holding_at_most int
		if _, err := fmt.Scan(&preOwnedClockPartsAmnt, &capable_of_holding_at_most); err != nil {
			log.Fatal("fmt.Scanln - err: ", err)
		}

		preOwnedClockParts := make([]int, 0, preOwnedClockPartsAmnt)
		preOwnedClockPartsCount := make(map[int]int)
		for i := 0; i < preOwnedClockPartsAmnt; i++ {
			var item int
			if _, err := fmt.Scan(&item); err != nil {
				log.Fatal("Scan: ", err)
			}
			preOwnedClockPartsCount[item]++

			preOwnedClockParts = append(preOwnedClockParts, item)
		}

		res := solution(preOwnedClockParts, preOwnedClockPartsCount, capable_of_holding_at_most)
		fmt.Printf("Case #%d: %v\n", caseNumber, res)
	}
}

const negativeAnswer = "NO"

func solution(preOwnedClockParts []int, preOwnedClockPartsCount map[int]int, capable_of_holding_at_most int) string {
	// fmt.Printf("input %v %v\n", preOwnedClockParts, capable_of_holding_at_most)
	// fmt.Printf("preOwnedClockPartsCount %v \n", preOwnedClockPartsCount)

	if len(preOwnedClockParts) > capable_of_holding_at_most*2 {
		return negativeAnswer
	}
	for _, v := range preOwnedClockPartsCount {
		if v > 2 {
			return negativeAnswer
		}
	}

	return "YES"
}
