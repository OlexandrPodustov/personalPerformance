package etl

import "strings"

const testVersion = 1

func Transform(in map[int][]string) (res map[string]int) {
	res = make(map[string]int)

	for i := 1; i <= 10; i++ {
		letters, ok := in[i]
		if ok {
			for _, vv := range letters {
				newV := strings.ToLower(vv)
				res[newV] = i
			}
		}
	}

	return res
}
