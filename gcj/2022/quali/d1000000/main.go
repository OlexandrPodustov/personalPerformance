package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	var testCasesAmount int
	if _, err := fmt.Scan(&testCasesAmount); err != nil {
		log.Fatal("parse testCasesAmount", err)
	}
	for testCaseNumber := 1; testCaseNumber <= testCasesAmount; testCaseNumber++ {
		var diceNumber int
		if _, err := fmt.Scan(&diceNumber); err != nil {
			log.Println(testCaseNumber, "  dimensions:", err)
		}
		dices := []int{}
		for i := 0; i < diceNumber; i++ {
			var dn int
			if _, err := fmt.Scan(&dn); err != nil {
				log.Println(testCaseNumber, "  dimensions:", err)
			}
			dices = append(dices, dn)
		}

		result := solve(dices)
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

func solve(dices []int) int {
	if !sort.IntsAreSorted(dices) {
		sort.Ints(dices)
	}
	maxLen := 0
	for i := 0; i < len(dices); i++ {
		if dices[i] > maxLen {
			maxLen++
		}
	}

	return maxLen
}
