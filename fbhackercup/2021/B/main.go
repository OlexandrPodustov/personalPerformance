package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var testCases int
	if _, err := fmt.Scanln(&testCases); err != nil {
		log.Fatal("fmt.Scanln(&t) testCases err", err)
	}

	for caseNumber := 1; caseNumber <= testCases; caseNumber++ {
		var rowsAmnt int
		if _, err := fmt.Scanln(&rowsAmnt); err != nil {
			log.Fatal("fmt.Scanln - err: ", err)
		}
		input := []string{}
		for i := 0; i < rowsAmnt; i++ {
			var row string
			if _, err := fmt.Scanln(&row); err != nil {
				log.Fatal("fmt.Scanln - err: ", err)
			}
			input = append(input, row)
		}

		res := solution(input)
		fmt.Printf("Case #%d: %v\n", caseNumber, res)
	}
}

func solution(input []string) string {
	// fmt.Println("len input", len(input))
	// fmt.Println("input", input)

	columns := extractColumns(input)
	input = append(input, columns...)

	// fmt.Println("input+columns", input)
	minAmntOfChanges := 1000_000_000

	solutions := map[int][]string{}
	for _, v := range input {
		if !strings.Contains(v, "O") {
			amntOfChanges := strings.Count(v, ".")
			if minAmntOfChanges > amntOfChanges {
				minAmntOfChanges = amntOfChanges
			}
			solutions[minAmntOfChanges] = append(solutions[minAmntOfChanges], v)
			// fmt.Println("possible solution", v, amntOfChanges)
		}
	}
	if minAmntOfChanges == 1000_000_000 {
		return "Impossible"
	}

	// fmt.Println("minAmntOfChanges", minAmntOfChanges)
	// fmt.Println("solutions", solutions)
	// fmt.Println("solutions with minAmntOfChanges", solutions[minAmntOfChanges])

	// 	differentSets := map[string]struct{}{}
	// 	for _, v := range solutions[minAmntOfChanges] {
	// 		differentSets[v] = struct{}{}
	// 	}
	// 	fmt.Println("differentSets", differentSets)
	// 	fmt.Println("len differentSets", len(differentSets))

	// int, // minimum number of additional Xs
	// int, // number of different sets of cells in which that minimum number of Xs could be placed
	// return fmt.Sprintf("%v %v", minAmntOfChanges, len(differentSets))
	return fmt.Sprintf("%v %v", minAmntOfChanges, len(solutions[minAmntOfChanges]))
}

func extractColumns(input []string) []string {
	result := make([]string, len(input))

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			// fmt.Println("input[i][j]", i, j, string(input[i][j]))
			result[j] += string(input[i][j])
		}
	}

	// fmt.Println("len result", len(result))
	// fmt.Println("result", result)

	return result
}
