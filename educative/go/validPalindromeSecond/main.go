package main

func validPalindrome(inputString string) bool {
	left := 0
	right := len(inputString) - 1

	var skipOneLetterIfNeeded bool
	for left < right {
		if inputString[left] != inputString[right] {
			if !skipOneLetterIfNeeded {
				if inputString[left] == inputString[right-1] {
					right--
				} else if inputString[left+1] == inputString[right] {
					left++
				}

				skipOneLetterIfNeeded = true
				continue
			}

			return false
		}

		left++
		right--
	}

	return true
}
