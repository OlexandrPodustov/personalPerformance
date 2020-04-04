package main

import (
	"fmt"
	"log"
)

func main() {
	var iterations int
	_, err := fmt.Scan(&iterations)
	if err != nil {
		log.Println("parse testCasesAmount:", err)
	}

	for testCaseNumber := 1; testCaseNumber <= iterations; testCaseNumber++ {
		var dimensions int
		_, err := fmt.Scan(&dimensions)
		if err != nil {
			log.Println(testCaseNumber, "  dimensions:", err)
		}

		var trace int
		var rowsRepeatedElements int
		columns := make([][]int, dimensions)
		for rn := 0; rn < dimensions; rn++ {
			col := make([]int, dimensions)
			columns[rn] = col
		}

		for rn := 0; rn < dimensions; rn++ {
			var row []int
			for e := 0; e < dimensions; e++ {
				var element int
				_, err := fmt.Scan(&element)
				if err != nil {
					log.Println(testCaseNumber, " row:", err)
				}

				columns[e][rn] = element
				row = append(row, element)
			}

			amnt := check(row)
			rowsRepeatedElements += amnt

			trace += row[rn]
		}

		var columnRepeatedElements int
		for _, col := range columns {
			amnt := check(col)
			columnRepeatedElements += amnt
		}

		fmt.Printf("Case #%v: %v %v %v\n", testCaseNumber, trace, rowsRepeatedElements, columnRepeatedElements)
	}
}

func check(input []int) int {
	uniqueMap := make(map[int]struct{}, len(input))

	for _, v := range input {
		if _, ok := uniqueMap[v]; !ok {
			uniqueMap[v] = struct{}{}
		} else {
			return 1
		}
	}

	return 0
}
