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
		var amount int
		if _, err := fmt.Scan(&amount); err != nil {
			log.Println(testCaseNumber, "  dimensions:", err)
		}

		deliciousness := make([]int, 0, amount)
		for i := 0; i < amount; i++ {
			var deliciousnessLevel int
			if _, err := fmt.Scan(&deliciousnessLevel); err != nil {
				log.Println(testCaseNumber, "  dimensions:", err)
			}
			deliciousness = append(deliciousness, deliciousnessLevel)
		}

		result := solve(deliciousness)
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

func solve(deliciousness []int) int {
	// fmt.Println(deliciousness)
	previosDeliciousness := 0
	customersPaid := 0

	for len(deliciousness) > 0 {
		mi := findMinIdx(deliciousness)
		// fmt.Println("mi", mi)

		if deliciousness[mi] >= previosDeliciousness {
			previosDeliciousness = deliciousness[mi]
			customersPaid++
		}

		if mi == 0 {
			deliciousness = deliciousness[1:]
		} else {
			deliciousness = deliciousness[:mi]
		}
		// fmt.Println("deliciousness", deliciousness)
	}

	return customersPaid
}

func findMinIdx(deliciousness []int) int {
	if lastIdx := len(deliciousness) - 1; deliciousness[0] > deliciousness[lastIdx] {
		return lastIdx
	}
	return 0
}
