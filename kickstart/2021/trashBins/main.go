package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var testCases int
	if _, err := fmt.Scanln(&testCases); err != nil {
		log.Fatal("fmt.Scanln(&t) testCases err", err)
	}

	for caseNumber := 1; caseNumber <= testCases; caseNumber++ {

		var amountOfHouses int
		if _, err := fmt.Scanln(&amountOfHouses); err != nil {
			log.Fatal("fmt.Scanln - err: ", err)
		}
		fmt.Println("am", amountOfHouses)
		var greetingString string
		if _, err := fmt.Scanln(&greetingString); err != nil {
			log.Fatal("fmt.Scanln - err: ", err)
		}
		res := solution(greetingString)
		fmt.Printf("Case #%d: %v\n", caseNumber, res)
	}
}

func solution(input string) int {
	fmt.Println("input", input)
	ss := strings.Split(input, "")

	cans := make([]int, 0, len(input))
	for _, v := range ss {
		it, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		cans = append(cans, it)
	}

	result := 0
	for _, v := range cans {
		fmt.Println("v", v)
		result += findLenthToNearestTrash()
	}

	return result
}

func findLenthToNearestTrash() int {
	// find all distances
	// if distance to first known trash is 0 - return 0
	// if not - calculate all and find minimal

	minimumDistance := 1000_000_000
	_ = minimumDistance

	return 0
}
