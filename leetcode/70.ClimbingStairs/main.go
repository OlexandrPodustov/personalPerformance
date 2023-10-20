package main

func climbStairs(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	}

	first, second := 1, 2
	res := first + second
	for i := 3; i < n; i++ {
		first = second
		second = res

		res = first + second
	}

	return res
}
