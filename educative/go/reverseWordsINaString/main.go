package main

import (
	"strings"
)

func reverseWords(sentence string) string {
	sentence = removeRedundantSpaces(sentence)
	resultSentence := reverse([]byte(sentence), 0, len(sentence)-1)

	{
		start := 0

		for i := 0; i < len(resultSentence); i++ {
			if resultSentence[i] == ` `[0] {
				end := i - 1
				resultSentence = reverse(resultSentence, start, end)
				start = i + 1
			}
		}

		resultSentence = reverse(resultSentence, start, len(resultSentence)-1)
	}

	return string(resultSentence)
}

func reverse(input []byte, start, end int) []byte {
	for start < end {
		input[start], input[end] = input[end], input[start]
		start++
		end--
	}

	return input
}

func removeRedundantSpaces(sentence string) string {
	r := strings.NewReplacer("  ", " ")
	for strings.Contains(sentence, "  ") {
		sentence = r.Replace(sentence)
	}
	return sentence
}
