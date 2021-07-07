package main

import "fmt"

// To execute Go code, please declare a func main() in a package "main"

/*

What is an anagram? Well, two words are anagrams of each other if they both contain the same letters. For example:
"abba" & "baab" == true
"abba" & "bbaa" == true
"abba" & "abbba" == false
"abba" & "abca" == false
"abcd" & "bdca" == true

Write a function that will find all the anagrams of a word from a list. You will be given two inputs a word and an array with words. You should return an array of all the anagrams or an empty array if there are none.

For example:
anagrams("abba", ["aabb", "abcd", "bbaa", "dada"]) => ["aabb", "bbaa"]
anagrams("racer", ["crazer", "carer", "racar", "caers", "racer"]) => ["carer", "racer"]
anagrams("laser", ["lazing", "lazy",  "lacer"]) => []

*/

func main() {
	// result := check("abba", []string{"aabb", "abcd", "bbaa", "dada"})
	// result := check("racer", []string{"crazer", "carer", "racar", "caers", "racer"})
	result := check("laser", []string{"lazing", "lazy", "lacer"})

	fmt.Println(result)
}

func check(input string, args []string) []string {
	result := []string{}

	inMap := make(map[rune]int)
	for _, v := range input {
		inMap[v]++
	}

	for _, a := range args {
		elementMap := make(map[rune]int)
		for _, v := range a {
			elementMap[v]++
		}

		if mapsEqual(inMap, elementMap) {
			result = append(result, a)
		}
	}

	return result
}

func mapsEqual(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}

	for ka, va := range a {
		vb, ok := b[ka]
		if !ok {
			return false
		}
		if va != vb {
			return false
		}
	}

	return true
}
