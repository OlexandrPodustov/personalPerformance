package main

import (
	"fmt"
	"log"
)

func main() {
	var testCasesAmount int
	if _, err := fmt.Scan(&testCasesAmount); err != nil {
		log.Fatal("parse testCasesAmount", err)
	}
	for testCaseNumber := 1; testCaseNumber <= testCasesAmount; testCaseNumber++ {
		var inp string
		if _, err := fmt.Scan(&inp); err != nil {
			log.Println(testCaseNumber, "  dimensions:", err)
		}
		result := solve(inp)
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

func solve(inp string) string {
	println("inp", inp)
	return ""
}
