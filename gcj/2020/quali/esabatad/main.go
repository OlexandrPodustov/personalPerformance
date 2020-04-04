// interactive
package main

import (
	"fmt"
)

func main() {
	var testCasesAmount, numberOfBitsInArray int
	_, err := fmt.Scan(&testCasesAmount, &numberOfBitsInArray)
	if err != nil {
		panic("parse testCasesAmount:" + err.Error())
	}
	fmt.Println("testCasesAmount, numberOfBitsInArray:", testCasesAmount, numberOfBitsInArray)

	for testCaseNumber := 1; testCaseNumber <= testCasesAmount; testCaseNumber++ {
		var activitiesToAssign int
		_, err := fmt.Scan(&activitiesToAssign)
		if err != nil {
			panic(testCaseNumber + " activitiesToAssign:" + activitiesToAssign + " " + err.Error())
		}

		result := "1"
		fmt.Println(result)
	}
}
