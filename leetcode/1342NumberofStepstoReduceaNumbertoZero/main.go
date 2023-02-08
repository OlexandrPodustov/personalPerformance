package main

func numberOfSteps(num int) int {
	numberOfSteps := 0
	for num != 0 {
		if num%2 == 0 {
			numberOfSteps++
			num /= 2
		} else {
			numberOfSteps++
			num--
		}
	}

	return numberOfSteps
}
