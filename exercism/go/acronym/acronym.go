package acronym

import (
	"strings"
)

const testVersion = 3

func Abbreviate(in string) string {
	var s, newString string
	splitByDash := strings.Split(in, "-")

	for _, e := range splitByDash {
		newString += e + " "
	}
	splitBySpace := strings.Split(newString, " ")
	for _, v := range splitBySpace {
		for ind, nv := range v {
			if ind == 0 {
				s += strings.ToUpper(string(nv))
			}
		}
	}

	return s
}
