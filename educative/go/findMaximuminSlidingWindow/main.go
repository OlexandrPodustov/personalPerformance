package main

func findMaxSlidingWindow(nums []int, windowSize int) []int {
	result := make([]int, 0)

	if windowSize > len(nums) {
		windowSize = len(nums)
	}

	deq := NewDeque()
	for i := 0; i < windowSize; i++ {
		for !deq.Empty() && nums[i] >= nums[deq.Back()] {
			deq.PopBack()
		}
		deq.PushBack(i)
	}
	result = append(result, nums[deq.Front()])

	for i := windowSize; i < len(nums); i++ {
		for !deq.Empty() && nums[i] >= nums[deq.Back()] {
			deq.PopBack()
		}

		if !deq.Empty() && deq.Front() <= (i-windowSize) {
			deq.PopFront()
		}

		deq.PushBack(i)
		result = append(result, nums[deq.Front()])
	}

	return result
}
