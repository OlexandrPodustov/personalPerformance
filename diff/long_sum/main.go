package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	r, err := sum("15", "999986")
	if err != nil {
		log.Fatal("failed to calculate sum", err)
	}

	fmt.Println(r)
}

func sum(a, b string) (string, error) {
	var aint []int
	for _, v := range strings.Split(a, "") {
		in, err := strconv.Atoi(v)
		if err != nil {
			return "", err
		}
		aint = append(aint, in)
	}

	var bint []int
	for _, v := range strings.Split(b, "") {
		in, err := strconv.Atoi(v)
		if err != nil {
			return "", err
		}
		bint = append(bint, in)
	}

	var remainder int
	var result string
	// for i:= len(aint); i>=0; i--{
	for len(aint) != 0 && len(bint) != 0 {
		aa := aint[len(aint)-1]
		bb := bint[len(bint)-1]

		x := aa + bb + remainder
		if x >= 10 {
			remainder = 1
		} else {
			remainder = 0
		}
		mod := x % 10

		result = fmt.Sprintf("%v%v", mod, result)

		aint = aint[:len(aint)-1]
		bint = bint[:len(bint)-1]
	}

	for i := len(aint); i > 0; i-- {
		v := aint[i-1]

		if remainder > 0 {
			v += remainder
			if v >= 10 {
				v -= 10
				remainder = 1
			} else {
				remainder = 0
			}

		}

		el := strconv.Itoa(v)
		result = fmt.Sprintf("%v%v", el, result)
	}

	for i := len(bint); i > 0; i-- {
		v := bint[i-1]

		if remainder > 0 {
			v += remainder
			if v >= 10 {
				v -= 10
				remainder = 1
			} else {
				remainder = 0
			}

		}

		el := strconv.Itoa(v)
		result = fmt.Sprintf("%v%v", el, result)
	}

	if remainder > 0 {
		result = fmt.Sprintf("%v%v", remainder, result)
	}

	return result, nil
}
