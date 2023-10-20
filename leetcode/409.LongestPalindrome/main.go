package main

func longestPalindrome(s string) int {
	if len(s) == 1 {
		return 1
	}
	lettersCount := make(map[rune]int)
	for _, v := range s {
		lettersCount[v]++
	}

	res := 0
	for _, v := range lettersCount {
		res += v / 2 * 2
	}
	if res != len(s) {
		res++
	}

	return res
}
