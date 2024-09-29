package main

func sortArrayByParity(nums []int) []int {
	res := make([]int, len(nums))

	insertEven := 0
	insertOdd := len(nums) - 1
	for _, v := range nums {
		if v%2 == 0 {
			res[insertEven] = v
			insertEven++
			continue

		}
		res[insertOdd] = v
		insertOdd--
	}

	return res
}
