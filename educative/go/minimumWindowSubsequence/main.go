package main

import "fmt"

func minWindow(str1 string, str2 string) string {
	fmt.Println("str1, str2", str1, str2)

	startIndex1 := 0
	startIndex2 := 0
	minLen := len(str1)
	minSubstring := ""
	start, end := 0, 1

	for ; startIndex1 < len(str1); startIndex1++ {
		if str1[startIndex1] == str2[0] {
			start = startIndex1
		}
		fmt.Println(startIndex2, "len(str2)", len(str2))

		if startIndex2 >= len(str2) {
			end = startIndex1

			if nlen := len(str1[start:end]); nlen < minLen {
				minLen = nlen
				minSubstring = str1[start:end]
			}
			continue
		}
		fmt.Println(startIndex1, startIndex2, "match", str1[startIndex1] == str2[startIndex2], str1[startIndex1], str2[startIndex2])
		if str1[startIndex1] == str2[startIndex2] {
			startIndex2++
		}
	}
	fmt.Println()

	return minSubstring
}
