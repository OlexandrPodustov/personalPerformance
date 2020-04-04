package main

import (
	"fmt"
	"sort"
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

		result := assign(activitiesToAssign)
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

// startDuplicates := make(map[int]int)
// endddDuplicates := make(map[int]int)
func assign(activitiesToAssign int) string {
	result := ""

	act := make([]activity, 0)
	for l := 0; l < activitiesToAssign; l++ {
		var start, end int
		_, err := fmt.Scan(&start, &end)
		if err != nil {
			fmt.Println("activitiesToAssign:", l, activitiesToAssign, err)
			return ""
		}

		act = append(act, activity{s: start, e: end})
	}

	sort.SliceStable(act, func(i, j int) bool {
		return act[i].s < act[j].s
	})
	// fmt.Println("sorted act:", act)

	c, j := 0, 0
	for _, v := range act {
		s := v.s
		e := v.e

		if s >= c {
			result += "C"
			c = e
		} else if s >= j {
			result += "J"
			j = e
		} else {
			return "IMPOSSIBLE"
		}
	}

	return result
}

type activity struct {
	s, e int
}
