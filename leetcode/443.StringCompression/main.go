package main

import "fmt"

func compress(chars []byte) int {
	if len(chars) == 0 {
		return 0
	}
	if len(chars) == 1 {
		return 1
	}

	result := ""
	prevChar := chars[0]
	amount := 0
	for _, v := range chars {
		// fmt.Println("v, prevChar", v, prevChar)

		if v == prevChar {
			amount++
		} else {
			if amount == 1 {
				result += string(prevChar)
			} else {
				result += fmt.Sprintf("%v%d", string(prevChar), amount)
				// fmt.Println("result", result)
				prevChar = 0
				amount = 1
			}
		}
		prevChar = v
	}

	if amount == 1 {
		result += string(prevChar)
	} else {
		result += fmt.Sprintf("%v%d", string(prevChar), amount)
	}

	for i := 0; i < len(result); i++ {
		chars[i] = result[i]
	}

	return len(result)
}
