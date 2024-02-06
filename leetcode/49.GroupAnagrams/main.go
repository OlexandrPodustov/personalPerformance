package main

func groupAnagrams(strs []string) [][]string {
	data := map[[26]uint16][]string{}

	for _, w := range strs {
		stats := wordStats(w)
		data[stats] = append(data[stats], w)
	}

	res := make([][]string, len(data))
	var i int
	for _, val := range data {
		res[i] = val
		i++
	}
	return res
}

func wordStats(w string) [26]uint16 {
	res := [26]uint16{}
	for _, c := range w {
		res[c-'a']++
	}

	return res
}
