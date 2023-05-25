package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var testCases int
	if _, err := fmt.Scanln(&testCases); err != nil {
		log.Fatal("fmt.Scanln(&t) testCases err", err)
	}

	for caseNumber := 1; caseNumber <= testCases; caseNumber++ {
		var greetingString string
		if _, err := fmt.Scanln(&greetingString); err != nil {
			log.Fatal("fmt.Scanln - err: ", err)
		}
		res := solution(greetingString)
		fmt.Printf("Case #%d: %d\n", caseNumber, res)
	}
}

func solution(input string) int {
	// getAllLetters()
	if len(input) == 1 {
		return 0
	}

	letters := map[string]int{}
	for _, v := range strings.Split(input, "") {
		letters[v]++
	}
	priceToChange := calc(letters)
	min := 1000_000_000
	for _, v := range priceToChange {
		if min > v {
			min = v
		}
	}

	return min
}

func calc(letters map[string]int) map[string]int {
	alphabet := getAllLetters()

	for singleLetter := range alphabet {
		res := calcFor(letters, singleLetter)
		alphabet[singleLetter] = res
	}

	return alphabet
}

func calcFor(letters map[string]int, let string) int {
	localLetters := map[string]int{}
	for k, v := range letters {
		if k == let {
			continue
		}
		localLetters[k] = v
	}

	cc := countConsonants(localLetters)
	cv := countVowels(localLetters)

	result := 0
	if isVowel(let) {
		result = cc + cv*2
	} else {
		result = cv + cc*2
	}

	return result
}

func isVowel(letter string) bool {
	_, ok := getVowels()[letter]
	// fmt.Println("isVowel letter", letter, ok)

	return ok
}

func countVowels(letters map[string]int) int {
	sum := 0
	vows := getVowels()
	for v, amnt := range letters {
		if _, ok := vows[v]; ok {
			// fmt.Println("vows letter amount", v, amnt)
			sum += amnt
		}
	}

	return sum
}

func countConsonants(letters map[string]int) int {
	sum := 0
	cons := getConsonants()
	for v, amnt := range letters {
		if _, ok := cons[v]; ok {
			// fmt.Println("cons letter amount", v, amnt)
			sum += amnt
		}
	}

	return sum
}

func getVowels() map[string]struct{} {
	return map[string]struct{}{"A": {}, "E": {}, "I": {}, "O": {}, "U": {}} // "A", "E", "I", "O", and "U"
}

func getConsonants() map[string]struct{} {
	vowels := getVowels()

	consonants := map[string]struct{}{}
	for letter := 'A'; letter < 'A'+26; letter++ {
		// fmt.Println("getConsonants", string(letter))
		if _, ok := vowels[string(letter)]; !ok {
			consonants[string(letter)] = struct{}{}
		}
	}

	return consonants
}

func getAllLetters() map[string]int {
	letters := map[string]int{}
	for letter := 'A'; letter < 'A'+26; letter++ {
		// print(string(letter))
		letters[string(letter)] = 0
	}
	// print("\n")

	return letters
}
