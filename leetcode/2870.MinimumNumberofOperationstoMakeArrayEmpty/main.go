package main

func minOperations(nums []int) int {
	counters := make(map[int]int)
	for _, v := range nums {
		counters[v]++
	}
	result := 0
	for k, v := range counters {
		switch v {
		case 1:
			return -1
		case 2, 3:
			result++
			delete(counters, k)
		case 4:
			result++
			result++
			delete(counters, k)

		default:
			atLeastOneOperationWithThree := false
			r := v / 3
			if r > 0 {
				atLeastOneOperationWithThree = true
			}
			result += r
			counters[k] -= r * 3

			r2 := (v - r*3) / 2
			result += r2
			counters[k] -= r2 * 2

			if counters[k] == 1 && atLeastOneOperationWithThree {
				counters[k] = 0
				result++
			}
		}

		if counters[k] == 1 {
			return -1
		}
	}

	return result
}
