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
		var rows [][]int
		for rn := 0; rn < dimensions; rn++ {
			var row []int
			for e := 0; e < dimensions; e++ {
				var rowEl int
				_, err := fmt.Scan(&rowEl)
				if err != nil {
					log.Println(testCaseNumber, " row:", err)
				}

				row = append(row, rowEl)
			}

			rows = append(rows, row)

			trace += row[rn]
		}
		fmt.Println("trace", trace)

		for i, v := range rows {
			fmt.Println("rows", i, v)
		}

		// result, amnt := check()
		// fmt.Printf("Case #%v: %v %v\n", testCaseNumber, result, amnt)
	}
}

func check() (bool, int) {
	return false, 0
}
