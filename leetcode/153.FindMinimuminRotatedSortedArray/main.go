package main

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	start := 0
	end := len(nums) - 1
	for start < end-1 {
		middle := (end-start)/2 + start
		if middle == start {
			middle++
		}
		// fmt.Println("start, middle, end", start, middle, end)
		if nums[start] < nums[middle] && nums[middle] < nums[end] {
			return nums[start]
		} else if nums[start] < nums[middle] {
			start = middle
		} else {
			end = middle
		}
	}

	return min(nums[start], nums[end])
}

// [4,5,6,7,0,1,2]
// -5000 <= n <= 5000
// 1 <= len(nums) <= 5000

// O(N) time compl
// O(1) space compl

// [2,3,1]

// suggested solution with the binary search
// O(log(N)) time compl
// O(1) space compl

// [2,3,4,5,6,7,1]
// len == 7
// mid == len/2 =>3.5=>3

// iteration 2
// startIndex = 3
// end = 6
// mid = 1+ 3 => 4

// iteration 3
// startIndex = 4
// end = 6
// mid = (6-4)/2 + 4 => 5

// iteration 4
// startIndex = 5
// end = 6
// mid = (6-5)/2 + 5 => 5
// middle increment => 6

// [7,1,2,3,4,5,6]
// len == 7
// mid == len/2 =>3.5=>3
