package main

import "fmt"

func circularArrayLoop(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if hasCircle(nums, i) {
			return true
		}
	}

	return false
}

func hasCircle(nums []int, start int) bool {
	slow, fast1, fast2 := start, start, start
	for i := 0; i < len(nums); i++ {
		// oldSlow := slow
		slow = move(nums, slow)
		fast1 = move(nums, fast1)
		fast2 = move(nums, fast1)
		fmt.Println("hasCircle slow, _, fast", slow, fast1, fast2)

		// if slow == fast2 && fast1 != fast2 {
		if slow == fast2 {
			return true
		}
		fast1 = fast2

		// if oldSlow == slow {
		// 	return false
		// }
	}

	return false
}

func move(nums []int, startIndex int) int {
	fmt.Println("move start 1", startIndex)
	val := nums[startIndex]
	newIdx := (val + startIndex) % len(nums)

	if newIdx < 0 {
		newIdx = len(nums) + newIdx
	}
	fmt.Println("move start 3", newIdx)

	return newIdx
}
