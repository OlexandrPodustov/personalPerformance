package bob

import (
	"strings"
)

const testVersion = 3

func Hey(inn string) string {
	var res string
	var resArrs []int32
	var allUpper bool
	var silence bool
	var lastSymb int32
	var spacesAmount int
	var amountLower int

	in := strings.Trim(inn, " ")
	for _, v := range in {
		lastSymb = v
		uc := strings.ToUpper(string(lastSymb))
		if uc >= "A" && uc <= "Z" {
			resArrs = append(resArrs, v)
		}
		sym := string(v)
		if sym == "" || sym == " " || sym == "\t" || sym == "\n" || sym == "\r" {
			spacesAmount++
		}

	}
	for _, val := range resArrs {
		if val < 65 || val > 90 {
			amountLower++
		}
		if amountLower == 0 {
			allUpper = true
		} else {
			allUpper = false
		}
	}
	if spacesAmount == len(in) || in == "" {
		silence = true
	}

	if allUpper {
		res = "Whoa, chill out!"
	} else if lastSymb == '?' {
		res = "Sure."
	} else if silence {
		res = "Fine. Be that way!"
	} else {
		res = "Whatever."
	}

	return res
}
