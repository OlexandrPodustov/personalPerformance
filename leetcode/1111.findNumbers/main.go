package main

func findNumbers(nums []int) int {
	even := 0

	for _, v := range nums {
		count := 0
		for v != 0 {
			v /= 10
			count++
		}
		if count%2 == 0 {
			even++
		}
	}

	return even
}
