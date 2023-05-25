package main

func twoSum(nums []int, target int) []int {
	numberIndex := make(map[int]int, len(nums))
	for i, v := range nums {
		numberIndex[v] = i
	}

	result := make([]int, 0, 2)
	for i, v := range nums {
		j, ok := numberIndex[target-v]
		if ok && i != j {
			result = append(result, i)
			result = append(result, j)

			break
		}
	}

	return result
}
