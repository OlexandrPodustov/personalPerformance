package cryptosquare

import (
	"log"
	"math"
	"strings"
	"unicode"
)

const testVersion int = 2

func Encode(m string) string {
	m = normalize(m)
	log.Println(math.Sqrt(float64(len(m))))
	c := int(math.Sqrt(float64(len(m))))
	log.Println(c)
	//calculate r




	//var slsl [][]string
	//for i ,v := range m	{
	//	slsl =
	//}

	return m
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
	}
	m = string(mnew)
	log.Println(m)

	return m
}
