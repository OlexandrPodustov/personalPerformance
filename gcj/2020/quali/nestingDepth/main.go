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
		intVal, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Sprintf("intVal, err := strconv.Atoi(v): %v, %v", v, err)
		}

		bef := strings.Repeat("(", intVal)
		aft := strings.Repeat(")", intVal)
		newPart := fmt.Sprintf("%s%s%s", bef, v, aft)
		result += newPart
	}

	if len(input) > 1 {
		result = trimRedundand(result)
	}

	return result
}

func trimRedundand(stringWithBraces string) string {
	r := strings.NewReplacer("(", "", ")", "")
	stringNoBraces := r.Replace(stringWithBraces)
	// fmt.Println("stringWithBraces:", stringWithBraces)

	result := ""
	for len(stringNoBraces) > 0 {
		current := string(stringNoBraces[0])
		if len(stringNoBraces) == 1 {
			result += stringWithBraces
			break
		}
		next := string(stringNoBraces[1])

		currIndex := strings.Index(stringWithBraces, current)
		if currIndex < 0 {
			fmt.Println("currIndex:" + strconv.Itoa(currIndex) + " stringWithBraces " + stringWithBraces + " current " + current)
			break
		}
		nextIndex := strings.Index(stringWithBraces[currIndex+1:], next)
		if nextIndex < 0 {
			fmt.Println("nextIndex:" + strconv.Itoa(nextIndex) + " stringWithBraces[currIndex+1:] " + stringWithBraces[currIndex+1:] + " next " + next)
			break
		}

		cur, err := strconv.Atoi(current)
		if err != nil {
			fmt.Println("cur, err := strconv.Atoi(current):", err)
			return ""
		}
		nex, err := strconv.Atoi(next)
		if err != nil {
			fmt.Println("cur, err := strconv.Atoi(current):", err)
			return ""
		}

		// fmt.Println("A", stringWithBraces, " - cur, n values", cur, nex, " - curr, n idx", currIndex, currIndex+nextIndex+1)

		// if cur-nex == 0 {
		if cur != 0 && nex != 0 && cur-nex == 0 {
			// fmt.Println("B", stringWithBraces, " - cur, n values", cur, nex, " - curr, n idx", currIndex, currIndex+nextIndex+1)
			result += stringWithBraces[:currIndex+1]
			stringWithBraces = stringWithBraces[currIndex+nextIndex+1:]
			// } else if cur-nex != 0 {
		} else if cur != 0 && nex != 0 && cur-nex != 0 {
			// fmt.Println("C", stringWithBraces, " - cur, n values", cur, nex, " - curr, n idx", currIndex, currIndex+nextIndex+1)

			intermidPart := stringWithBraces[currIndex+1 : currIndex+nextIndex+1]

			for {
				oldLen := len(intermidPart)
				intermidPart = strings.Replace(intermidPart, ")(", "", -1)

				if oldLen == len(intermidPart) {
					break
				}
			}

			before := stringWithBraces[:currIndex+1]
			affter := stringWithBraces[currIndex+nextIndex+1:]

			stringWithBraces = before + intermidPart + affter
			// } else if cur == 0 && nex == 0 {
		} else {
			result += stringWithBraces[:1]
			stringWithBraces = stringWithBraces[1:]
		}
		// fmt.Println("D", stringWithBraces)

		stringNoBraces = stringNoBraces[1:]
		// fmt.Println("stringNoBraces", stringNoBraces)
		// fmt.Println()
	}

	return result
}

// 8
// 0000
// 101
// 111000
// 1
// 021
// 312
// 4
// 221
