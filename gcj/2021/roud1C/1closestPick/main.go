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
		fmt.Println("n, k", n, k)

		ticketsSold := make([]int, 0, n)
		for i := 0; i < n; i++ {
			var n int
			if _, err := fmt.Scan(&n); err != nil {
				log.Println(testCaseNumber, "parse input:", err)
				return
			}

			ticketsSold = append(ticketsSold, n)
		}
		fmt.Println("ticketsSold", ticketsSold)

		result := check(n, k)
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

func check(n, k int) string {

	res := ""
	return res
}
