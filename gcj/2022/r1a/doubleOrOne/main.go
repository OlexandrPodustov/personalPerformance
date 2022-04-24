package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
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

	generateAllCombinations(inp)
	return ""
}

func generateAllCombinations(inp string) []string {
	set := []string{
		inp,
	}

	fmt.Println("       set", set)
	set = sort.StringSlice(set)
	fmt.Println("sorted set", set)

	res := "PEEEEL, PEEEELL, PEEEL, PEEELL, PEEL, PEELL, PPEEEEL, PPEEEELL, PPEEEL, PPEEELL, PPEEL, PPEELL"
	result := strings.Split(res, ", ")
	fmt.Println("result", len(result))
	fmt.Println("result", result)

	return nil
}
