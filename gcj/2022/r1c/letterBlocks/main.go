package main

import (
	"fmt"
	"log"
	"sort"
)

const impossible = "IMPOSSIBLE"

func main() {
	var testCasesAmount int
	if _, err := fmt.Scan(&testCasesAmount); err != nil {
		log.Fatal("parse testCasesAmount", err)
	}
	for testCaseNumber := 1; testCaseNumber <= testCasesAmount; testCaseNumber++ {
		var wordsAmount int
		if _, err := fmt.Scan(&wordsAmount); err != nil {
			log.Fatal(testCaseNumber, "", err)
		}

		words := make([]string, wordsAmount)
		for i := 0; i < wordsAmount; i++ {
			var word string
			if _, err := fmt.Scan(&word); err != nil {
				log.Fatal(testCaseNumber, "", err)
			}
			words[i] = word
		}

		result := solve(words)
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

func solve(words []string) string {
	fmt.Println("words", words)
	// 1. check each chunk - if it has repetitive letters separated like in "HASH"
	for _, word := range words {
		if hasNonConsecutiveRepetitiveLetters(word) {
			return impossible
		}
	}

	return impossible
}

func hasNonConsecutiveRepetitiveLetters(word string) bool {
	letters := make(map[rune]int)
	for _, letter := range word {
		letters[letter]++
	}
	for letter, amount := range letters {
		if amount == 1 {
			delete(letters, letter)
		}
	}
	if len(letters) == 0 {
		return false
	}
	// for letter := range letters {
	//
	// }

	srted := sort.SliceIsSorted(word, func(i, j int) bool {
		return word[i] < word[j]
	})
	fmt.Println(srted)

	fmt.Println(len(letters))
	fmt.Println(letters)

	return false
}
