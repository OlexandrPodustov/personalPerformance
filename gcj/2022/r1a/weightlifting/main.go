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
		var exercises, weightTypes int
		if _, err := fmt.Scan(&exercises, &weightTypes); err != nil {
			log.Println(testCaseNumber, "  dimensions:", err)
		}

		exWeights := make([][]int, 0, exercises)
		for i := 0; i < exercises; i++ {
			exerc := make([]int, 0, weightTypes)
			for j := 0; j < weightTypes; j++ {
				var numb int
				if _, err := fmt.Scan(&numb); err != nil {
					log.Println(testCaseNumber, "  dimensions:", err)
				}
				exerc = append(exerc, numb)
			}
			exWeights = append(exWeights, exerc)
		}

		result := solve(exWeights)
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

func solve(exWeights [][]int) string {
	fmt.Println(exWeights)
	mapWeights(exWeights)

	return ""
}

func mapWeights(ew [][]int) {
	commonWeights := make(map[int]struct{})
	for i := 0; i < len(ew); i++ {
		// fmt.Println("ew", i, ew[i])

		exerciseWeightsSet := make([]int, 0)
		for j := 0; j < len(ew[i]); j++ {
			wt := j + 1
			commonWeights[wt] = struct{}{}
			for jj := 1; jj <= ew[i][wt-1]; jj++ {
				exerciseWeightsSet = append(exerciseWeightsSet, wt)
			}
		}

		fmt.Println("exerciseWeightsSet", exerciseWeightsSet)
	}

	fmt.Println("commonWeights", commonWeights)
}
