package main

func findMaxSlidingWindow(nums []int, windowSize int) []int {
	result := make([]int, 0)

	deq := NewDeque()
	for _, v := range nums {
		if deq.Len() < windowSize {
			deq.PushBack(v)
			continue
		}
		result = append(result, deq.currentTop())
		deq.PopFront()
		deq.PushBack(v)
	}
	result = append(result, deq.currentTop())

	return result
}

func (s *Deque) currentTop() int {
	maxVal := -99_999
	for _, v := range s.items {
		if v > maxVal {
			maxVal = v
		}
	}

	return maxVal
}
