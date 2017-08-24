//Package scrabble helps us to calculate a scrabble score
package scrabble

import "strings"

const testVersion = 5

var scores map[int][]string = map[int][]string{
	1:  []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
	2:  []string{"D", "G"},
	3:  []string{"B", "C", "M", "P"},
	4:  []string{"F", "H", "V", "W", "Y"},
	5:  []string{"K"},
	8:  []string{"J", "X"},
	10: []string{"Q", "Z"},
}

//Score calculates scrabble score
func Score(word string) (res int) {
	newScoresMap := refactorScoreboard(scores)
	for _, inputLetter := range strings.ToLower(word) {
		if val, ok := newScoresMap[string(inputLetter)]; ok {
			res += val
		}
	}

	return res
}

func refactorScoreboard(in map[int][]string) (res map[string]int) {
	res = make(map[string]int)
	for key, sliceLetters := range in {
		for _, letter := range sliceLetters {
			letter = strings.ToLower(letter)
			res[letter] = key
		}
	}
	return res
}
