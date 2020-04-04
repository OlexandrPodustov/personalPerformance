// interactive
package main

import (
	"fmt"
)

func main() {
	var iterations int
	_, err := fmt.Scan(&iterations)
	if err != nil {
		fmt.Println("parse testCasesAmount:", err)
		return
	}

	for testCaseNumber := 1; testCaseNumber <= iterations; testCaseNumber++ {
		var activitiesToAssign int
		_, err := fmt.Scan(&activitiesToAssign)
		if err != nil {
			fmt.Println(testCaseNumber, " activitiesToAssign:", activitiesToAssign, err)
			return
		}

		// result := assign(activitiesToAssign)
		result := "assign(activitiesToAssign)"
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}
