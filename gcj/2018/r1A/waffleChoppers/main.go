package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var testCasesAmount int
	if _, err := fmt.Scan(&testCasesAmount); err != nil {
		fmt.Println("parse testCasesAmount:", err)
		return
	}

	for testCaseNumber := 1; testCaseNumber <= testCasesAmount; testCaseNumber++ {
		var rows, columns, horizontal, vertical int
		if _, err := fmt.Scan(&rows, &columns, &horizontal, &vertical); err != nil {
			fmt.Println("&rows, &columns, &horizontal, &vertical:", err)
			return
		}

		var ca int
		for i := 0; i < rows; i++ {
			var row string
			if _, err := fmt.Scan(&row); err != nil {
				fmt.Printf("scan row: %v", err)
				os.Exit(1)
			}
			ca += strings.Count(row, "@")
		}

		fmt.Println("ca total", ca, "parts", rows, columns, horizontal, vertical, 4)
	}
}

// calculate how many parts will be after division
// divide total amount of chocolate chips by total parts
// check if this result is possible for each combination
