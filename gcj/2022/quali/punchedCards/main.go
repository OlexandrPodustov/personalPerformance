package main

import (
	"fmt"
	"log"
)

const (
	dot    string = "."
	plus   string = "+"
	pillar string = "|"
)

func main() {
	var testCasesAmount int
	if _, err := fmt.Scan(&testCasesAmount); err != nil {
		log.Fatal("parse testCasesAmount", err)
	}
	for testCaseNumber := 1; testCaseNumber <= testCasesAmount; testCaseNumber++ {
		var rows int
		var columns int
		if _, err := fmt.Scan(&rows, &columns); err != nil {
			log.Println(testCaseNumber, "  dimensions:", err)
		}
		result := build(rows, columns)
		fmt.Printf("Case #%v:\n%v\n", testCaseNumber, result)
	}
}

func build(rows, columns int) string {
	result := ""
	for i := 0; i < rows; i++ {
		result += lineDivider(i, columns) + "\n"
		result += rowData(i, columns) + "\n"
	}
	result += lineDivider(10, columns)
	return result
}

func lineDivider(row, columnsAmount int) string {
	startEndDivider := plus
	if row == 0 {
		startEndDivider = dot
	}
	body := ""
	for i := 0; i < columnsAmount; i++ {
		if row == 0 && i == 0 {
			body += dot + plus
			continue
		}
		body += "-" + plus
	}
	return fmt.Sprintf("%v%v", startEndDivider, body)
}

func rowData(row, columnsAmount int) string {
	body := ""
	for i := 0; i < columnsAmount; i++ {
		if row == 0 && i == 0 {
			body += dot + dot
			continue
		}
		body += pillar + dot
	}
	return body + pillar
}
