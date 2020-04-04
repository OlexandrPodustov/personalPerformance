// interactive
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var testCasesAmount, numberOfBitsInArray int
	if _, err := fmt.Scan(&testCasesAmount, &numberOfBitsInArray); err != nil {
		panic("parse testCasesAmount:" + err.Error())
	}

	for testCaseNumber := 1; testCaseNumber <= testCasesAmount; testCaseNumber++ {
		err := solve(numberOfBitsInArray)
		if err != nil {
			fmt.Printf("failed to solve tc num: %v, err: %v", testCaseNumber, err)
			break
		}
	}
}

func solve(numberOfBitsInArray int) error {
	r, err := makeRequest(numberOfBitsInArray)
	if err != nil {
		return fmt.Errorf("makeRequest. resp: %v, %v", r, err)
	}
	fmt.Println(r) // give response to the judge

	var judgeDecision string
	if _, err = fmt.Scan(&judgeDecision); err != nil {
		return fmt.Errorf("scan judgeDecision. received: %T, %#v", judgeDecision, err)
	}
	if judgeDecision == "N" {
		return fmt.Errorf("received negative judgeDecision: %v", judgeDecision)
	}

	return nil
}

func makeRequest(numberOfBitsInArray int) (string, error) {
	response := make([]int, numberOfBitsInArray)

	// requestLimit := 150 / numberOfBitsInArray
	for i := 1; i <= numberOfBitsInArray; i++ {
		fmt.Println(i)

		var pByte int
		if _, err := fmt.Scan(&pByte); err != nil {
			return "", fmt.Errorf("scan. received: %T, %#v, %v", pByte, pByte, err)
		}

		response[i-1] = pByte
	}

	r := ""
	for _, v := range response {
		r += strconv.Itoa(v)
	}

	return r, nil
}
