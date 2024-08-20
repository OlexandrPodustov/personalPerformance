package main

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	overallMax := nums[0]

	slidingResult := 1
	prev := nums[0]
	anyOnes := prev == 1
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		if curr == 1 {
			anyOnes = true
		}
		if prev == curr && curr == 1 {
			slidingResult++
			prev = curr
			continue
		}
		overallMax = max(overallMax, slidingResult)
		slidingResult = 1
		prev = curr
	}

	if !anyOnes {
		return 0
	}
	overallMax = max(overallMax, slidingResult)

	return overallMax
}
