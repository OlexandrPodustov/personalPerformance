package main

import (
	"fmt"
)

func circularArrayLoop(arr []int) bool {
	currentDirection := false
	for i, v := range arr {
		if abs(arr[i]) == len(arr) {
			continue
		}
		if v >= 0 {
			currentDirection = true
		} else {
			currentDirection = false
		}
		slowPtr := i
		fastPtr := i

		fmt.Printf("\n\t\tInitially slow pointer = %d, Fast pointer = %d\n", slowPtr, fastPtr)
		fmt.Printf("\t\tStarting from value = %d\n\n", v)

		for slowPtr != fastPtr || slowPtr != -1 || fastPtr != -1 {
			slowPtr = nextStep(arr, slowPtr, currentDirection)
			if slowPtr == -1 {
				break
			}

			fastPtr = nextStep(arr, fastPtr, currentDirection)

			fmt.Printf("\t\tSlow pointer = %d\n", slowPtr)
			fmt.Printf("\t\tFast pointer after first-step = %d\n", fastPtr)
			if fastPtr != -1 {
				fastPtr = nextStep(arr, fastPtr, currentDirection)
				fmt.Printf("\t\tFast pointer after second-step = %d\n", fastPtr)
			}
			if fastPtr == -1 || slowPtr == fastPtr {
				break
			}
		}
		if slowPtr == fastPtr && slowPtr != -1 {
			return true
		}
	}
	return false
}

func nextStep(array []int, currentIndex int, currentDirection bool) int {
	nextDirection := false
	if array[currentIndex] >= 0 {
		nextDirection = true
	}

	// Loop can't be found if the loop direction is changed
	if nextDirection != currentDirection || abs(array[currentIndex]%len(array)) == 0 {
		return -1
	}

	findStep := currentIndex + array[currentIndex]
	findStep = findStep % len(array)

	if findStep < 0 {
		findStep = findStep + len(array)
	}

	return findStep
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
