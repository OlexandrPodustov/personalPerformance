package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = removeNonAlphanumeric(s)
	s = strings.ToLower(s)

	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if j >= len(s)/2 {
			if s[i] != s[j] {
				return false
			}
		}
		j--
	}

	return true
}

func removeNonAlphanumeric(input string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(input, "")
}
