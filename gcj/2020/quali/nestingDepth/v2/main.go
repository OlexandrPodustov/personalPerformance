package main

import (
	"fmt"
	"strconv"
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
		var stringOfDigits string
		_, err := fmt.Scan(&stringOfDigits)
		if err != nil {
			fmt.Println(testCaseNumber, "  stringOfDigits:", stringOfDigits, err)
			return
		}

		result := wrap(stringOfDigits)
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

func wrap(input string) string {
	result := ""
	ss := strings.Split(input, "")
	for i := 0; i < len(ss); i++ {
		v := ss[i]
		intVal, err := strconv.Atoi(v)
		if err != nil {
			panic("intVal, err := strconv.Atoi(v):" + v + err.Error())
		}

		bef := strings.Repeat("(", intVal)
		aft := strings.Repeat(")", intVal)
		result += fmt.Sprintf("%s%s%s", bef, v, aft)
	}

	if len(input) > 1 {
		result = trimRedundand(result)
	}

	return result
}

func trimRedundand(stringWithBraces string) string {
	for {
		oldLen := len(stringWithBraces)
		stringWithBraces = strings.Replace(stringWithBraces, ")(", "", -1)

		if oldLen == len(stringWithBraces) {
			break
		}
	}

	return stringWithBraces
}
