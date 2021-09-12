package main

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
