package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)
		keyStr := fmt.Sprint(s)

		groups[keyStr] = append(groups[keyStr], str)
	}

	result := make([][]string, 0)
	for _, grp := range groups {
		result = append(result, grp)
	}

	return result
}
