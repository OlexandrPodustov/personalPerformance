package pangram

import (
	"strings"
)

const testVersion = 1

func IsPangram(in string) bool {
	// create initial array with full alphabet
	var res bool
	var alf []int32
	for i := 'a'; i < 'a'+26; i++ {
		alf = append(alf, i)
	}

	// lowercase
	inLow := strings.ToLower(in)

	// range loop
	for _, v := range inLow {
		for i, tv := range alf {
			if v == tv {
				alf = append(alf[:i], alf[i+1:]...)
			}
		}
		if len(alf) == 0 {
			res = true
		}
	}

	return res
}
