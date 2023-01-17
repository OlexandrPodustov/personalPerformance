package main

import "fmt"

func happyNumber(num int) bool {
	slow := num
	fast := sumDigits(num)

	fmt.Println()
	fmt.Println("num, slow, fast", num, slow, fast)
	for {
		if fast == 1 {
			return true
		}
		if fast == slow {
			break
		}
		slow = sumDigits(slow)
		fast = sumDigits(sumDigits(fast))

		fmt.Println("slow, fast", slow, fast)
	}
	fmt.Println("cycle detected")

	return false
}

// pow calculates the power of the given digit
func pow(digit int, power int) int {
	res := 1
	for i := 0; i < power; i++ {
		res = res * digit
	}
	return res
}

// sumDigits is a helper function that calculates the sum of digits.
func sumDigits(number int) int {
	totalSum := 0
	for number > 0 {
		digit := number % 10
		number = number / 10
		totalSum += pow(digit, 2)
	}
	return totalSum
}
