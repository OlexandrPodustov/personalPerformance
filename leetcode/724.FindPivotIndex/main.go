package main

func pivotIndex(nums []int) int {
	lnums := len(nums)
	leftSum, rightSum := make([]int, 0, lnums), make([]int, lnums)

	{
		lsum := 0
		leftSum = append(leftSum, lsum)

		for i := 1; i < lnums; i++ {
			lsum += nums[i-1]
			leftSum = append(leftSum, lsum)
		}
	}
	// fmt.Println("leftSum", leftSum)
	{
		rsum := 0
		rightSum[lnums-1] = 0

		for i := lnums - 2; i >= 0; i-- {
			rsum += nums[i+1]
			rightSum[i] = rsum
		}
	}
	// fmt.Println("rightSum", rightSum)

	for i := 0; i < lnums; i++ {
		if leftSum[i] == rightSum[i] {
			return i
		}
	}

	return -1
}
