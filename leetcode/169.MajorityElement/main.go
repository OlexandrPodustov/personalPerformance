package main

// time O(n)
// space O(1)
// 0 < N < 11
// [2, 2, 3, 3, 3]
// k=2 amount=2
// k=3 amount=3

// space O(N)
// sort complexity n*log(n)

func majorityElement(nums []int) int {
	var majority, currentAmount int
	for i := 0; i < len(nums); i++ {
		if currentAmount == 0 {
			majority = nums[i]
		}

		if nums[i] == majority {
			currentAmount++
		} else {
			currentAmount--
		}

	}

	return majority
}
