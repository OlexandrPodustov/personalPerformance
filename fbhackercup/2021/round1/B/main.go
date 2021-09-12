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
		var nRows, mCols, aTopLeft, bTopRight int
		if _, err := fmt.Scanln(&nRows, &mCols, &aTopLeft, &bTopRight); err != nil {
			log.Fatal("fmt.Scanln - err: ", err)
		}

		res := solution(nRows, mCols, aTopLeft, bTopRight)
		fmt.Printf("Case #%d: %v\n", caseNumber, res)
	}
}

func solution(nRows, mCols, aTopLeft, bTopRight int) string {
	minSteps := nRows + mCols - 1
	if minSteps > aTopLeft || minSteps > bTopRight {
		return "Impossible"
	}
	atl := aTopLeft / minSteps
	btr := bTopRight / minSteps

	minCommon := 0
	if atl > btr {
		minCommon = btr
	} else {
		minCommon = atl
	}

	var rows [][]int
	for i := 0; i < nRows; i++ {
		row := make([]int, mCols)
		rows = append(rows, row)
	}

	fillTheField(rows, minCommon)
	// fmt.Println("rows fillTheField", rows)

	maxNumberOnTheRoute := 1
	for !ATopLeftEnough(rows, aTopLeft) {
		fmt.Println("ATopLeftEnough", rows)

		maxNumberOnTheRoute++
		incrementOneOnAtopLeft(rows)
	}
	for !BTopRightEnough(rows, bTopRight) {
		maxNumberOnTheRoute++
		incrementOneOnBtopRight(rows)
	}

	setRest(rows, maxNumberOnTheRoute+1)
	fmt.Println("setRest", rows)

	end := ""
	for i, vv := range rows {
		elems := []string{}
		for _, v := range vv {
			elems = append(elems, strconv.Itoa(v))
		}
		end += strings.Join(elems, " ")

		if len(rows)-1 != i {
			end += "\n"
		}
	}

	return fmt.Sprintf("Possible\n%+v", end)
}

func incrementOneOnAtopLeft(input [][]int) {
	input[1][len(input[0])-1]++
}

func incrementOneOnBtopRight(input [][]int) {
	input[1][0]++
}

func ATopLeftEnough(input [][]int, aTopLeft int) bool {
	total := 0
	for j := 0; j < len(input[0]); j++ {
		total += input[0][j]
	}
	for i := 1; i < len(input); i++ {
		total += input[i][len(input[i])-1]
	}

	return total == aTopLeft
}

func BTopRightEnough(input [][]int, bTopRight int) bool {
	total := 0
	for j := 0; j < len(input[0]); j++ {
		total += input[0][j]
	}
	for i := 1; i < len(input); i++ {
		total += input[i][0]
	}

	return total == bTopRight
}

func fillTheField(input [][]int, minCommon int) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			input[i][j] = minCommon
		}
	}

	for i := 1; i < len(input); i++ {
		for j := 1; j < len(input[i])-1; j++ {
			input[i][j] = 1001
		}
	}
}

func setRest(input [][]int, numb int) {
	for i := 1; i < len(input); i++ {
		for j := 1; j < len(input[i])-1; j++ {
			input[i][j] = numb
		}
	}
}
