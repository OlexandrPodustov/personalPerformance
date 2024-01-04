package main

import (
	"math"
)

func myAtoi(s string) int {
	var result int
	var integersMet bool
	integerPostive := true
	signMet := false
	resultSet := make([]int, 0)
	for _, v := range s {
		vi := v - 48
		isInteger := vi >= 0 && vi < 10 // 48 to 57 are integers
		if integersMet && !isInteger {
			break
		}

		if v == '-' { // save sign somewhere
			if !signMet && !integersMet {
				signMet = true
				integerPostive = false
				continue
			} else {
				return 0
			}
		}
		if v == '+' {
			if !signMet && !integersMet {
				signMet = true
				continue
			} else {
				return 0
			}
		}
		if v == ' ' && !signMet && !integersMet { // ignore if it's leading or stop if not
			continue
		}
		if !integersMet && !isInteger {
			return 0
		}
		if isInteger {
			integersMet = true
			resultSet = append(resultSet, int(vi))
			if len(resultSet) > 0 && resultSet[0] == 0 {
				resultSet = resultSet[1:]
			}
		}
	}

	outOfBound := false
	if len(resultSet) > 10 {
		result = math.MaxInt32
		outOfBound = true
	} else {
		multiplier := 1
		for i := len(resultSet) - 1; i >= 0; i-- {
			result += resultSet[i] * multiplier
			multiplier *= 10
			if result > math.MaxInt32 {
				result = math.MaxInt32
				outOfBound = true
				break
			}
		}
	}

	if !integerPostive {
		remainder := 0
		if outOfBound {
			remainder = -1
		}

		result *= -1
		result += remainder

	}
	return result
}
