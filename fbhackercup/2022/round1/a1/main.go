package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	var testCases int
	if _, err := fmt.Scanln(&testCases); err != nil {
		log.Fatal("fmt.Scanln(&t) testCases err", err)
	}

	for caseNumber := 1; caseNumber <= testCases; caseNumber++ {
		var cardsAmount, Ktimes int
		if _, err := fmt.Scan(&cardsAmount, &Ktimes); err != nil {
			log.Fatal("fmt.Scanln - err: ", err)
		}

		fromAcards := make([]int, 0, cardsAmount)
		for i := 0; i < cardsAmount; i++ {
			var item int
			if _, err := fmt.Scan(&item); err != nil {
				log.Fatal("Scan: ", err)
			}

			fromAcards = append(fromAcards, item)
		}

		toBcards := make([]int, 0, cardsAmount)
		for i := 0; i < cardsAmount; i++ {
			var item int
			if _, err := fmt.Scan(&item); err != nil {
				log.Fatal("Scan: ", err)
			}

			toBcards = append(toBcards, item)
		}

		fmt.Printf("cardsAmount, Ktimes %v %v\n", cardsAmount, Ktimes)
		fmt.Printf("fromAcards %v\n", fromAcards)
		fmt.Printf("toBcards %v\n", toBcards)

		res := solution(fromAcards, toBcards, Ktimes)
		fmt.Printf("Case #%d: %v\n", caseNumber, res)
	}
}

const (
	negativeAnswer = "NO"
	positiveAnswer = "YES"
)

func solution(fromAcards, toBcards []int, kTimes int) string {
	if reflect.DeepEqual(fromAcards, toBcards) {
		return positiveAnswer
	}
	if kTimes == 0 {
		return negativeAnswer
	}

	return positiveAnswer
}
