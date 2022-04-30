package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var testCasesAmount int
	if _, err := fmt.Scan(&testCasesAmount); err != nil {
		log.Fatal("parse testCasesAmount", err)
	}
	for testCaseNumber := 1; testCaseNumber <= testCasesAmount; testCaseNumber++ {
		var numbersGiven, maxNumbersAdd int
		if _, err := fmt.Scan(&numbersGiven, &maxNumbersAdd); err != nil {
			log.Fatal(testCaseNumber, "", err)
		}

		numbers := make([]int, numbersGiven)
		for i := 0; i < numbersGiven; i++ {
			var number int
			if _, err := fmt.Scan(&number); err != nil {
				log.Fatal(testCaseNumber, "", err)
			}
			numbers[i] = number
		}

		result := solve(numbers, maxNumbersAdd)
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

func solve(numbers []int, maxNumbersAdd int) string {
	result := "IMPOSSIBLE"
	res := diff(numbers)
	if !equal(append(numbers, res)) {
		return result
	}
	result = strconv.Itoa(res)
	return result
}

func diff(numbers []int) int {
	sumOfSquares := 0
	sum := 0
	for _, v := range numbers {
		sum += v
		sumOfSquares += v * v
	}
	if sum == 0 {
		return 0
	}

	return (sumOfSquares - sum*sum) / (2 * sum)
}

func equal(numbers []int) bool {
	sumOfSquares := 0
	sum := 0
	for _, v := range numbers {
		sum += v
		sumOfSquares += v * v
	}

	return sumOfSquares == sum*sum
}
