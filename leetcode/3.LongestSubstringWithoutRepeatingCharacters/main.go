package main

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	metChars := make(map[byte]int)
	left := 0
	right := 0

	maxLen := 0
	for right < len(s) {
		metChars[s[right]]++
		for metChars[s[right]] > 1 {
			metChars[s[left]]--
			left++
		}

		right++
		maxLen = max(maxLen, right-left)
	}

	return maxLen
}
