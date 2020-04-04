package main

import (
	"fmt"
	"sort"
	"strings"
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

func assign(activitiesToAssign int) string {
	startDuplicates := make(map[int]int)
	endddDuplicates := make(map[int]int)

	act := make(As, 0)
	for l := 0; l < activitiesToAssign; l++ {
		var start, end int
		_, err := fmt.Scan(&start, &end)
		if err != nil {
			fmt.Println("activitiesToAssign:", l, activitiesToAssign, err)
			return ""
		}

		startDuplicates[start]++
		endddDuplicates[end]++

		act = append(act, Activity{S: start, E: end, index: l})
	}

	for _, v := range startDuplicates {
		if v > 2 {
			return "IMPOSSIBLE"
		}
	}

	for _, v := range endddDuplicates {
		if v > 2 {
			return "IMPOSSIBLE"
		}
	}

	sort.Sort(act)

	c, j := 0, 0
	result := make([]string, len(act))
	for _, v := range act {
		s := v.S
		e := v.E
		idx := v.index

		if s >= c {
			result[idx] = "C"
			c = e
		} else if s >= j {
			result[idx] = "J"
			j = e
		} else {
			return "IMPOSSIBLE"
		}
	}

	return strings.Join(result, "")
}

type Activity struct {
	S, E, index int
}

type As []Activity

func (a As) Len() int           { return len(a) }
func (a As) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a As) Less(i, j int) bool { return a[i].S < a[j].S }
