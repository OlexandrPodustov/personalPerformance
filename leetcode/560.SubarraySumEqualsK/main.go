package sse

func SubarraySum(nums []int, k int) int {
	amountOfSubarrays := 0

	ln := len(nums)
	sum := 0
	previousStartIndex := 0
	i := 0
	for previousStartIndex < ln {
		sum += nums[i]
		// fmt.Println("previousStartIndex, index, element, sum ", previousStartIndex, i, nums[i], sum)
		if sum == k {
			amountOfSubarrays++
		}
		i++
		if i == ln {
			previousStartIndex++
			i = previousStartIndex
			sum = 0
		}
	}

	return amountOfSubarrays
}
