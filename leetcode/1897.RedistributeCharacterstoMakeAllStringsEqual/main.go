package main

func makeEqual(words []string) bool {
	count := make(map[rune]int)
	for _, word := range words {
		for _, v := range word {
			count[v]++
		}
	}

	lw := len(words)
	for _, v := range count {
		if v%lw != 0 {
			return false
		}
	}

	return true
}
