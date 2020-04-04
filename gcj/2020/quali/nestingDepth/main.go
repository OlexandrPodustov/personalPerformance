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

		result := zeroOrOne(stringOfDigits)
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

func zeroOrOne(input string) string {
	result := ""
	ss := strings.Split(input, "")
	for i := 0; i < len(ss); i++ {
		v := ss[i]
		if v == "0" {
			result += v
		} else {
			newPart := fmt.Sprintf("(%s)", v)
			result += newPart
		}
	}

	if len(input) > 1 {
		result = trimRedundand(result)
	}

	return result
}

func trimRedundand(stringWithBraces string) string {
	r := strings.NewReplacer("(", "", ")", "")
	stringWithoutBraces := r.Replace(stringWithBraces)
	// originn := strings.Split(origin, "")

	fmt.Println("full origin:", stringWithoutBraces)
	fmt.Println("stringWithBraces:", stringWithBraces)

	result := ""
	for len(stringWithoutBraces) > 0 {
		current := string(stringWithoutBraces[0])
		if len(stringWithoutBraces) == 1 {
			break
		}
		next := string(stringWithoutBraces[1])

		currIndex := strings.Index(stringWithBraces, current)
		if currIndex < 0 {
			fmt.Println("currIndex:", currIndex, stringWithBraces, current)
			break
		}
		nextIndex := strings.Index(stringWithBraces[currIndex+1:], next)
		if nextIndex < 0 {
			fmt.Println("nextIndex:", nextIndex, stringWithBraces, next)
			break
		}

		cur, _ := strconv.Atoi(current)
		nex, _ := strconv.Atoi(next)

		if cur != 0 && nex != 0 && cur-nex == 0 {
			fmt.Println(stringWithBraces, " - cur, n values", cur, nex, " - curr, n idx", currIndex, currIndex+nextIndex+1)
			bef := stringWithBraces[:currIndex+1]
			after := stringWithBraces[currIndex+nextIndex+1]
			result += bef + string(after)
		} else {
			// result += stringWithBraces[:currIndex+1]
		}

		stringWithBraces = stringWithBraces[currIndex+1:]
		stringWithoutBraces = stringWithoutBraces[1:]
	}

	return result
}
