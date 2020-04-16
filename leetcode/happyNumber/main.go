package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(isHappy(19))
	// fmt.Println(isHappy(2))
	fmt.Println(isHappy(7))
	// fmt.Println(isHappy(8))
	// fmt.Println(isHappy(145))
}

var limit int

func isHappy(n int) bool {
	// limit++
	// if limit > 20 {
	// 	return false
	// }
	if n%145 == 0 {
		return false
	}

	nts := strings.Split(strconv.Itoa(n), "")
	if len(nts) == 1 {
		if nts[0] == "1" {
			return true
		}
	}

	res := 0
	for _, v := range nts {
		// fmt.Println(v)
		nn, _ := strconv.Atoi(v)
		res += nn * nn
	}
	// fmt.Println("res", limit, res)

	return isHappy(res)
}
