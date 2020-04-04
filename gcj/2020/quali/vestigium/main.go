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
				fmt.Println("column", e, columns[e])

				row = append(row, element)
			}

			amnt := check(row)
			rowsRepeatedElements += amnt

			trace += row[rn]
		}
		fmt.Println("trace", trace)
		fmt.Println("rowsRepeatedElements", rowsRepeatedElements)

		var columnRepeatedElements int
		for i, col := range columns {
			fmt.Println("col", i, col)

			amnt := check(col)
			columnRepeatedElements += amnt
		}
		fmt.Println("columnRepeatedElements", columnRepeatedElements)

		// fmt.Printf("Case #%v: %v %v\n", testCaseNumber, result, amnt)
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
