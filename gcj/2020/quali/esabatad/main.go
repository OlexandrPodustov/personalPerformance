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
	// for i := 0; i < 4; i++ {
	for j := 1; j <= numberOfBitsInArray; j++ {
		response[j-1] = getNthByte(j)
	}
	// }

	r := ""
	for _, v := range response {
		r += strconv.Itoa(v)
	}

	return r, nil
}

// func swap()           {}
// func complement()     {}
// func swapComplement() {}

func getNthByte(position int) int {
	fmt.Println(position)

	var pByte int
	if _, err := fmt.Scan(&pByte); err != nil {
		errVal := fmt.Sprintf("scan. received: %T, %#v, %v", pByte, pByte, err)
		panic(errVal)
	}

	return pByte
}
