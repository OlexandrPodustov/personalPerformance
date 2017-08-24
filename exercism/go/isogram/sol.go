// Package isogram contains only one function and does what that function does
package isogram

import "strings"

const testVersion = 1

// IsIsogram helps us to check whether the word is an isogram or not
func IsIsogram(in string) (res bool) {
	var countmap map[rune]int = make(map[rune]int)
	ns := strings.Replace(strings.ToLower(in), "-", "", -1)
	ns = strings.Replace(strings.ToLower(ns), " ", "", -1)
	for _, v := range ns {
		countmap[v]++
	}
	for _, v := range countmap {
		if v > 1 {
			return false
		}
	}
	return true
}
