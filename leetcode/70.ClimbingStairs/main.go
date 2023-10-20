package main

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	if n == 4 {
		return 5
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
