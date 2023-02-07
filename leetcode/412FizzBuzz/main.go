package main

import "strconv"

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := range result {
		switch {
		case (i+1)%15 == 0:
			result[i] = "FizzBuzz"
		case (i+1)%5 == 0:
			result[i] = "Buzz"
		case (i+1)%3 == 0:
			result[i] = "Fizz"

		default:
			result[i] = strconv.Itoa(i + 1)
		}
	}

	return result
}
