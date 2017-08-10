package hamming

import (
	"fmt"
)

const testVersion = 6

func Distance(a, b string) (int, error) {
	var cnt int
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				cnt++
			}
		}
	} else {
		return 0, fmt.Errorf("length of strings is different")
	}

	return cnt, nil
}
