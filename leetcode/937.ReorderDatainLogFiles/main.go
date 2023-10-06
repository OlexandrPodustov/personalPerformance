package main

import (
	"regexp"
	"sort"
	"strings"
)

// The letter-logs come before all digit-logs.
// The letter-logs are sorted lexicographically by their contents. If their contents are the same, then sort them lexicographically by their identifiers.
// The digit-logs maintain their relative ordering.

var str = regexp.MustCompile(`[0-9]`)

func reorderLogFiles(logs []string) []string {
	var letterPartOfTheresult []string
	var digitPartOfTheresult []string

	for _, log := range logs {
		sstrs := strings.Split(log, " ")
		if !str.MatchString(sstrs[1]) {
			letterPartOfTheresult = append(letterPartOfTheresult, log)
			continue
		}
		digitPartOfTheresult = append(digitPartOfTheresult, log)
	}

	sort.SliceStable(letterPartOfTheresult, func(i, j int) bool {
		is := strings.Split(letterPartOfTheresult[i], " ")
		js := strings.Split(letterPartOfTheresult[j], " ")
		iss := strings.Join(is[1:], " ")
		jss := strings.Join(js[1:], " ")
		if iss == jss {
			return letterPartOfTheresult[i] < letterPartOfTheresult[j]
		}

		return iss < jss
	})

	return append(letterPartOfTheresult, digitPartOfTheresult...)
}
