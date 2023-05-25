package isogram

import "strings"

func IsIsogram(word string) bool {
	runeCount := make(map[string]int)
	for _, v := range word {
		if sv := string(v); sv != " " && sv != "-" {
			lowers := strings.ToLower(sv)
			runeCount[lowers]++
		}
	}
	for _, v := range runeCount {
		if v > 1 {
			return false
		}
	}

	return true
}
