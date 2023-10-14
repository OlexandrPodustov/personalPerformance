package main

import (
	"strings"
)

func isValid(s string) bool {
	r := strings.NewReplacer("()", "", "[]", "", "{}", "")

	oldLen := 0
	for oldLen != len(s) {
		oldLen = len(s)
		s = r.Replace(s)
	}

	return s == ""
}

// todo: implement with stack
