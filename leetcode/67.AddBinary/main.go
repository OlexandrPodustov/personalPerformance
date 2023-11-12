package main

import (
	"strings"
)

func addBinary(a string, b string) string {
	maxlen := max(len(a), len(b))
	if diff := len(a) - len(b); diff > 0 {
		b = strings.Repeat("0", diff) + b
	}
	if diff := len(b) - len(a); diff > 0 {
		a = strings.Repeat("0", diff) + a
	}

	remainder := 0
	result := ""
	for i := maxlen - 1; i >= 0; i-- {
		nv := int(a[i]-'0') + int(b[i]-'0') + remainder

		if nv%2 == 0 {
			result = "0" + result
		} else {
			result = "1" + result
		}
		if nv > 1 {
			remainder = 1
		} else {
			remainder = 0
		}
		// fmt.Println("i, result", i, result, a[i], b[i], remainder, nv)
	}
	if remainder > 0 {
		result = "1" + result
	}

	return result
}
