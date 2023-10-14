package main

func isAnagram(s string, t string) bool {
	smap := make(map[rune]int)
	tmap := make(map[rune]int)

	for _, v := range s {
		smap[v]++
	}
	for _, v := range t {
		tmap[v]++
	}
	if len(smap) != len(tmap) {
		return false
	}

	for letter, amnt := range tmap {
		if am := smap[letter]; am != amnt {
			return false
		}
	}

	return true
}
