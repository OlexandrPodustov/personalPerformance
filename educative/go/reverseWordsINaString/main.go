package main

func reverseWords(sentence string) string {
	resultSentence := reverse([]byte(sentence), 0, len(sentence)-1)

	{
		start := 0
		end := 0

		for i := range resultSentence {
			if resultSentence[i] == ` `[0] {
				end = i - 1
				resultSentence = reverse(resultSentence, start, end)
				start = i + 1
			}
		}

		end = len(sentence) - 1
		resultSentence = reverse(resultSentence, start, end)
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
