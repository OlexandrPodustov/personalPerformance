package main

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	m1s := map[string]int{}
	for _, v := range strings.Split(s1, " ") {
		m1s[v]++
	}

	m2s := map[string]int{}
	for _, v := range strings.Split(s2, " ") {
		m2s[v]++
	}

	uncommon := []string{}
	for k, v := range m1s {
		if v == 1 {
			if _, ok := m2s[k]; !ok {
				uncommon = append(uncommon, k)
			}
		}
	}
	for k, v := range m2s {
		if v == 1 {
			if _, ok := m1s[k]; !ok {
				uncommon = append(uncommon, k)
			}
		}
	}

	return uncommon
}
