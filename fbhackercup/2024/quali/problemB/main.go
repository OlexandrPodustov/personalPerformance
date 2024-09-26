package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var testCases int
	if _, err := fmt.Scanln(&testCases); err != nil {
		log.Fatal("fmt.Scanln(&t) testCases err", err)
	}

	for caseNumber := 1; caseNumber <= testCases; caseNumber++ {
		var lines, chance float64
		if _, err := fmt.Scan(&lines, &chance); err != nil {
			log.Fatal("fmt.Scanln - err: ", err)
		}

		fmt.Printf("Case #%d: %v\n", caseNumber, solve(lines, chance))
	}
}

func solve(lines, chance float64) float64 {
	ch := chance / 100
	probability := math.Pow(ch, (lines-1)/lines)

	return (probability - ch) * 100
}
