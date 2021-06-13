package main

import (
	"fmt"
	"log"
)

func main() {
	var iterations int
	if _, err := fmt.Scanln(&iterations); err != nil {
		log.Println("parse testCasesAmount:", err)
		return
	}
	// fmt.Println("iterations", iterations)

	for testCaseNumber := 1; testCaseNumber <= iterations; testCaseNumber++ {
		var n, k int
		if _, err := fmt.Scanln(&n, &k); err != nil {
			log.Println(testCaseNumber, "parse n+k:", err)
			return
		}
		var str string
		if _, err := fmt.Scanln(&str); err != nil {
			log.Println(testCaseNumber, "parse str:", err)
			return
		}

		fmt.Println("n, k, str", n, k, str)

		result := check(n, k, str)
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

func check(n, k int, str string) string {

	return "" // modulo 10^9+7
}
