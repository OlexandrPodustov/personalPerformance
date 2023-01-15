package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

const testVersion = 2

func Valid(s string) (b bool) {
	s = strings.Replace(s, " ", "", -1)

	if len(s) < 2 {
		return
	}
	for _, v := range s {
		if !unicode.IsDigit(v) {
			return
		}
	}
	sl := make([]int, len(s))

	var nv int
	for i, v := range s {
		ii, r := strconv.Atoi(string(v))
		if r != nil {
			panic(r)
		}

		if len(s)%2 == 0 {
			if i%2 == 0 {
				nv = ii * 2
				if nv > 9 {
					nv -= 9
				}
				sl[i] = nv
			} else {
				sl[i] = ii
			}
		} else {
			if i%2 != 0 {
				nv = ii * 2
				if nv > 9 {
					nv -= 9
				}
				sl[i] = nv
			} else {
				sl[i] = ii
			}
		}
	}

	var sum int
	for _, v := range sl {
		sum += v
	}
	if sum%10 == 0 {
		b = true
	}

	return
}
