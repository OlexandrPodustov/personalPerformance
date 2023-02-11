package main

func findRepeatedSequences(s string, k int) []string {
	windowSize := k
	if len(s) <= windowSize {
		return make([]string, 0)
	}

	result := make([]string, 0)
	lenght := len(s)
	sequences := make(map[string]int, lenght)

	start, end := 0, windowSize
	for end < lenght {
		sequences[s[start:end]]++

		start++
		end++
	}
	for k, v := range sequences {
		if v >= 2 {
			result = append(result, k)
		}
	}

	return result
}
