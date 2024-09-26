package main

import (
	"fmt"
	"log"
	"math"
)

const (
	possible   = "YES"
	impossible = "NO"
)

func main() {
	var testCases int
	if _, err := fmt.Scanln(&testCases); err != nil {
		log.Fatal("fmt.Scanln(&t) testCases err", err)
	}

	for caseNumber := 1; caseNumber <= testCases; caseNumber++ {
		var travelers, timeLimit uint64
		if _, err := fmt.Scan(&travelers, &timeLimit); err != nil {
			log.Fatal("fmt.Scanln - err: ", err)
		}

		times := make([]uint64, 0, travelers)
		for i := 0; i < int(travelers); i++ {
			var tt uint64
			if _, err := fmt.Scan(&tt); err != nil {
				log.Fatal("fmt.Scanln - err: ", err)
			}
			times = append(times, tt)
		}

		fmt.Printf("Case #%d: %v\n", caseNumber, solve(travelers, timeLimit, times))
	}
}

func solve(travelers, timeLimit uint64, times []uint64) string {
	fastestTraveller := findMin(times)

	if travelers > 2 {
		return possibleIfTimeLimit(fastestTraveller*(travelers-2)*2+fastestTraveller, timeLimit)
	}

	return possibleIfTimeLimit(fastestTraveller, timeLimit)
}

func findMin(times []uint64) uint64 {
	var res uint64 = math.MaxUint64
	for _, v := range times {
		if v <= res {
			res = v
		}
	}
	return res
}

func possibleIfTimeLimit(totalTravelTime, timeLimit uint64) string {
	if totalTravelTime <= timeLimit && totalTravelTime > 1 {
		return possible
	}
	return impossible
}
