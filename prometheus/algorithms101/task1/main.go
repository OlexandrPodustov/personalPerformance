package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

}

var iteration int

func multiply(x, y uint64) uint64 {

	iteration++
	a, b, n1 := divide(x)
	c, d, n2 := divide(y)

	var ac, bd, adbc uint64
	if a > 9 || b > 9 || c > 9 || d > 9 {
		ac = multiply(a, c)
		bd = multiply(b, d)
	}
	ab := (a + b)
	cd := (c + d)

	fmt.Println(iteration, "a", a, "b", b, "c", c, "d", d)

	if ab > 9 || cd > 9 {
		fmt.Println(iteration, "internal ab, cd", ab, cd)
		ad := multiply(a, d)
		bc := multiply(b, c)

		adbc = ad + bc
		fmt.Println(iteration, "adbc", adbc)
	}

	ac = a * c
	bd = b * d
	fmt.Println(iteration, "ac", ac, "bd", bd)

	adbc = ab*cd - ac - bd
	fmt.Println(iteration, "adbc", adbc)

	n10 := math.Pow10(n1)
	n102 := math.Pow10(n2 / 2)
	fmt.Println(iteration, "math.Pow10(n1)", n10)
	fmt.Println(iteration, "math.Pow10(n2)", n102)

	result := uint64(n10)*ac + uint64(n102)*adbc + bd
	fmt.Println(iteration, "ac, adbc, bd", ac, adbc, bd)
	fmt.Println(iteration, "ac, adbc, bd", ac, adbc, bd)
	fmt.Println(iteration, "result", result)

	return result
}

// 287769407308846640970310151509826255482575362419155842891311909556878670000425352112987881085839680

func divide(x uint64) (uint64, uint64, int) {
	// str := strconv.Itoa(x)
	str := strconv.FormatUint(x, 10)

	if len(str) == 1 {
		return 0, x, 1
	}

	div := len(str) / 2
	var yy, zz string

	for i, v := range strings.Split(str, "") {
		if i < div {
			yy += v
		} else {
			zz += v
		}
	}

	// fmt.Println("yy", yy)
	// fmt.Println("zz", zz)
	// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.

	y, err := strconv.ParseUint(yy, 10, 0)
	if err != nil {
		panic(err)
	}
	z, err := strconv.ParseUint(zz, 10, 0)
	if err != nil {
		panic(err)
	}

	return y, z, len(str)
}
