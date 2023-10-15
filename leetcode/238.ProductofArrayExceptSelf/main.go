package main

func productExceptSelf(nums []int) []int {
	{
		amount, position := countZeroes(nums)
		if amount > 1 {
			makeAllZeroes(nums)
			return nums
		} else if amount == 1 {
			num := multiply(nums[:position])
			if res := multiply(nums[position+1:]); res != 0 {
				if num == 0 {
					num = res
				} else {
					num *= res
				}
			}

			makeAllZeroes(nums)
			nums[position] = num
			return nums
		}
	}

	var precur int
	var cur int
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			cur = nums[i]
			nums[i] = multiply(nums[i+1:])
			precur = cur
			continue
		}
		if i != 0 && i != len(nums)-1 {
			cur *= nums[i]
		}

		// fmt.Println("", i, len(nums), nums[i+1:], cur, multiply(nums[i+1:]))

		nums[i] = multiply(nums[i+1:]) * precur
		if i == len(nums)-1 {
			nums[i] = cur
		}
		precur = cur
	}

	return nums
}

func multiply(nums []int) int {
	var result int
	for i, v := range nums {
		if i == 0 {
			result = v
			continue
		}
		result *= v
	}

	return result
}

func countZeroes(nums []int) (int, int) { // return amount and position if 1 zero
	var result int
	position := -1
	for i, v := range nums {
		if v == 0 {
			result++
			position = i
		}
	}

	return result, position
}

func makeAllZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		nums[i] = 0
	}

	return
}
