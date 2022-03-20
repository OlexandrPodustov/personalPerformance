package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var testCases int
	if _, err := fmt.Scanln(&testCases); err != nil {
		log.Fatal("testCases err", err)
	}

	for caseNumber := 1; caseNumber <= testCases; caseNumber++ {
		var given int
		if _, err := fmt.Scanln(&given); err != nil {
			log.Fatal("given", err)
		}

		res := solution(given)
		fmt.Printf("Case #%d: %v\n", caseNumber, res)
	}
}

func solution(given int) int {
	if given == 0 {
		return 90
	}
	var sum int
	str := strconv.Itoa(given)
	for i := 0; i < len(str); i++ {
		res, err := strconv.Atoi(string(str[i]))
		if err != nil {
			log.Fatal("waaaaat", err)
		}
		sum += res
	}

	var newStr string
	rem := 9 - sum%9
	if rem == 9 {
		rem = 0
	}
	if len(str) == 1 && given > rem {
		newStr = fmt.Sprint(rem, str)
		// println("1 newStr ", newStr)
	} else {

		// println("str", str, rem)
		newStr = fmt.Sprint(str, rem)
		// println("newStr", newStr)
	}

	result, err := strconv.Atoi(newStr)
	if err != nil {
		log.Fatal("final waaaaat", err)
	}
	return result
}
