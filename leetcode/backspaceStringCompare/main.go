package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(backspaceCompare("###abc##d", "ad"))
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))

	S, T := "a##c", "#a#c"
	fmt.Println(backspaceCompare(S, T))
	S = "a#c"
	T = "b"
	fmt.Println(backspaceCompare(S, T))
}

func backspaceCompare(S string, T string) bool {
	if getPureString(S) == getPureString(T) {
		return true
	}

	return false
}

func getPureString(str string) string {
	s := strings.Split(str, "")
	// fmt.Println("before", s)

	var bAtTheBeginning int
	for i := 0; i < len(s); i++ {
		if s[i] != "#" {
			break
		}
		bAtTheBeginning++
	}
	s = s[bAtTheBeginning:]
	// fmt.Println("after", s)

	cleanString := ""

	var trimEl int
	for i := len(s) - 1; i >= 0; i-- {
		curEl := s[i]
		if curEl == "#" {
			trimEl++
		} else {
			if trimEl > 0 {
				trimEl--
			} else {
				cleanString = fmt.Sprintf("%v%v", curEl, cleanString)
			}
		}
	}
	// fmt.Println("cleanString", cleanString)

	return cleanString
}
