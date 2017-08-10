package raindrops

import "strconv"

const testVersion = 3

func Convert(i int) string {
	var s string

	if i%(3*5*7) == 0 {
		s = "PlingPlangPlong"
	} else if i%(3*5) == 0 {
		s = "PlingPlang"
	} else if i%(5*7) == 0 {
		s = "PlangPlong"
	} else if i%(3*7) == 0 {
		s = "PlingPlong"
	} else if i%3 == 0 {
		s = "Pling"
	} else if i%5 == 0 {
		s = "Plang"
	} else if i%7 == 0 {
		s = "Plong"
	} else {
		s = strconv.Itoa(i)
	}
	return s
}
