package main

func sortedSquares(nums []int) []int { // retry with two pointers
	if len(nums) == 1 {
		nums[0] *= nums[0]
		return nums
	}
	var (
		stk        = New()
		insertHere = 0
		i          = 1
	)

	for ; i < len(nums); i++ {
		prev := abs(nums[i-1])
		curr := abs(nums[i])
		if prev >= curr {
			stk.Push(prev) // save to stack and continue
			continue
		}

		if prev <= stk.Peek() || (stk.Len() == 0 && prev < curr) {
			nums[insertHere] = prev * prev
		} else if prev > stk.Peek() {
			top := stk.Pop()
			nums[insertHere] = top * top
		}
		insertHere++

		for stk.Len() > 0 && stk.Peek() < curr {
			top := stk.Pop()
			nums[insertHere] = top * top
			insertHere++
		}
	}
	nums[insertHere] = nums[i-1] * nums[i-1]
	insertHere++

	for stk.Len() > 0 {
		top := stk.Pop()
		nums[insertHere] = top * top
		insertHere++
	}

	return nums
}

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}

type stack struct {
	data []int
}

func New() stack {
	return stack{data: make([]int, 0)}
}

func (s *stack) Len() int {
	return len(s.data)
}

func (s *stack) Peek() int {
	if s.Len() == 0 {
		return 0
	}
	return s.data[s.Len()-1]
}

func (s *stack) Pop() int {
	if s.Len() == 0 {
		return 0
	}
	el := s.data[s.Len()-1]
	s.data = s.data[:s.Len()-1]
	return el
}

func (s *stack) Push(num int) {
	s.data = append(s.data, num)
	return
}
