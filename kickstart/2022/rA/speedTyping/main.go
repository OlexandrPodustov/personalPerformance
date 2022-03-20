package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var testCases int
	if _, err := fmt.Scanln(&testCases); err != nil {
		log.Fatal("testCases err", err)
	}

	for caseNumber := 1; caseNumber <= testCases; caseNumber++ {

		var ideal, pfact string
		if _, err := fmt.Scanln(&ideal); err != nil {
			log.Fatal("fmt.Scan bags", err)
		}
		if _, err := fmt.Scanln(&pfact); err != nil {
			log.Fatal("fmt.Scan bags", err)
		}
		res := solution(ideal, pfact)
		if res < 0 {
			fmt.Printf("Case #%d: %v\n", caseNumber, "IMPOSSIBLE")
			continue
		}
		fmt.Printf("Case #%d: %v\n", caseNumber, res)
	}
}

func solution(ideal, fact string) int {
	if len(fact) < len(ideal) {
		return -1
	}
	if fact == ideal {
		return 0
	}
	if len(fact) == len(ideal) && fact != ideal {
		return -1
	}
	if strings.Contains(fact, ideal) {
		return len(fact) - len(ideal)
	}

	var amountToDelete int
	var fi int
	for i := 0; i < len(ideal); i++ {
		for {
			if len(fact) <= fi {
				return -1
			}
			if ideal[i] == fact[fi] {
				fi++
				break
			}
			amountToDelete++
			fi++
		}

	}
	if amountToDelete > len(fact)-len(ideal) {
		return -1
	}

	return len(fact) - len(ideal)
}
