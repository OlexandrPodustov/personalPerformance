package main

import (
	"math"
)

func minWindow(s, t string) string {
	if len(t) > len(s) || len(t) == 0 {
		return ""
	}
	targetSubstringMap := fillTheMap(t)
	start, end := 0, 0
	substringMatch := make(map[string]struct{})
	// fmt.Println("targetSubstringMap", targetSubstringMap)

	for start >= 0 && end < len(s) && start < len(s) {
		currentSymbol := s[start]
		if _, ok := targetSubstringMap[currentSymbol]; ok {
			targetSubstringMap[currentSymbol]--
			if targetSubstringMap[currentSymbol] == 0 {
				delete(targetSubstringMap, currentSymbol)
			}
		}
		// fmt.Println("start, end, targetSubstringMap", start, end, targetSubstringMap)

		if len(targetSubstringMap) == 0 {
			targetSubstringMap = fillTheMap(t)
			// fmt.Println("inner targetSubstringMap", targetSubstringMap)

			end = start + 1
			for ; start >= 0; start-- {
				// fmt.Println("inner maplen, start, end, targetSubstringMap", len(targetSubstringMap), start, end, targetSubstringMap)
				currentSymbol := s[start]

				if _, ok := targetSubstringMap[currentSymbol]; ok {
					targetSubstringMap[currentSymbol]--
					if targetSubstringMap[currentSymbol] == 0 {
						delete(targetSubstringMap, currentSymbol)
					}
				}
				if len(targetSubstringMap) == 0 {
					targetSubstringMap = fillTheMap(t)

					substringMatch[s[start:end]] = struct{}{}
					// start = end
					// start++
					break
				}
			}
			// fmt.Println("substringMatch", substringMatch)

		}

		start++
	}

	minLen := math.MaxInt64
	minSubstring := ""
	for key := range substringMatch {
		// fmt.Println("substringMatch key", key)

		if len(key) < minLen {
			minLen = len(key)
			minSubstring = key
		}
	}

	// fmt.Println("")

	return minSubstring
}

func fillTheMap(t string) map[byte]int {
	targetSubstringMap := make(map[byte]int)
	for _, v := range t {
		targetSubstringMap[byte(v)]++
	}
	return targetSubstringMap
}
