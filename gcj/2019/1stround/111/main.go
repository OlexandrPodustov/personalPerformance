package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var iterations int
	_, err := fmt.Scan(&iterations)
	if err != nil {
		log.Println("failed to parse testCasesAmount:", err)
	}

	for ii := 1; ii <= iterations; ii++ {
		result := check(ii)
		fmt.Printf("Case #%v: %v\n", ii, result)
	}

}

func check(ii int) string {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Println(ii, " failed to parse input:", err)
	}

	words := []string{}
	for i := 1; i <= n; i++ {
		var word string
		_, err := fmt.Scan(&word)
		if err != nil {
			log.Println(ii, " failed to parse input:", err)
		}
		words = append(words, word)
	}

	return strings.Join(words, ", ")
}
