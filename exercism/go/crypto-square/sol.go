package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

const testVersion int = 2

func Encode(m string) string {
	var le, r, c, root int
	m = normalize(m)
	le = len(m)

	root = round(math.Sqrt(float64(le)))
	intpart, div := math.Modf(float64(le) / float64(root))
	if int(intpart) < 0 {
		return ""
	} else if div == 0 {
		c = int(intpart)
		r = root
		//log.Println("intpart1 - c, r", c, r, c*r, le)
	} else {
		c = int(intpart) + 1
		r = int(intpart)
		if c*r < le {
			r++
		}
		//log.Println("intpart2 - c, r", c, r, c*r, le)
	}

	var slsl = make([]string, r)
	//m = "splun"
	for i := 0; i < r; i++ {

		if c < len(m) {
			//log.Println(i, ii, c, m[:ii])
			slsl[i] = m[:c]
			m = m[c:]
		} else {
			slsl[i] = m
		}

	}
	//slsl = []string{"spl", "un"}

	//log.Println("slsl ", slsl)
	var mm string
	for ic := 0; ic < c; ic++ {
		for ir := 0; ir < r; ir++ {
			if ic >= len(slsl[ir]) {
			} else {
				//log.Println("ic ", ic, "len ", len(slsl[ir]), slsl[ir], string(slsl[ir][ic]))
				mm += string(slsl[ir][ic])
			}
		}
		mm += " "
	}

	mm = mm[:len(mm)-1]
	//log.Println("mm ", mm)

	//log.Println()
	//log.Println()

	return mm
}

func round(val float64) int {
	//log.Println("round ", val)
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}

func normalize(m string) string {
	m = strings.ToLower(m)
	//log.Println(m)

	m = strings.Replace(m, " ", "", -1)
	//log.Println(m)

	var mnew []rune
	for _, v := range m {

		if unicode.IsLetter(v) {
			mnew = append(mnew, v)
		}
		if unicode.IsNumber(v) {
			mnew = append(mnew, v)
		}
	}
	m = string(mnew)
	//log.Println(m)

	return m
}
