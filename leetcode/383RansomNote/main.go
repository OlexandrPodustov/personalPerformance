package main

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	magLetters := make(map[rune]int, 26)
	for _, v := range magazine {
		magLetters[v]++
	}

	for _, v := range ransomNote {
		if magLetters[v] <= 0 {
			return false
		}
		magLetters[v]--
	}

	return true
}
