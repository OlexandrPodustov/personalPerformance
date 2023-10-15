package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	i := len(nums) / 2
	rightPointer := len(nums)

	var oldi int
	for i != oldi {
		oldi = i
		if nums[i] == target {
			return i
		} else if nums[i] < target { // search in rigth part
			i = (i + rightPointer) / 2
		} else { // search in left part
			rightPointer = i
			i = i / 2
		}
	}

	return -1
}

// i = 3. 5 < 2? - enter else block rightPointer=3 =>
// i = 1  0 < 2? - enter else if rightPointer=3 =>
// i = 2  3 < 2? - enter else rightPointer=2 =>
// i = 1  0 < 2? - enter else if rightPointer=1 =>
// i = 0  -1 <2? - enter else if rightPointer=1
