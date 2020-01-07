package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var iterations int
	_, err := fmt.Scan(&iterations)
	if err != nil {
		log.Println("failed to parse testCasesAmount:", err)
	}

	for ii := 1; ii <= iterations; ii++ {

		var tc string
		_, err := fmt.Scan(&tc)
		if err != nil {
			log.Println(ii, " failed to parse input:", err)
		}

		slint := make([]int, len(tc))
		for i, v := range tc {
			intVal, err := strconv.Atoi(string(v))
			if err != nil {
				log.Println(i, " failed to convert string into int:", err)
			}

			slint[i] = intVal
		}

		n1, n2 := check(slint)
		s1, s2 := convertToString(n1, n2)
		fmt.Printf("Case #%v: %v %v \n", ii, s1, s2)
	}
}

func check(slint []int) ([]int, []int) {
	var n1, n2 = make([]int, len(slint)), make([]int, len(slint))
	for i, v := range slint {
		n1[i] = v
		if v == 4 {
			n1[i] = v - 1
			n2[i] = 1
		}
	}

	return n1, n2
}

func convertToString(n1, n2 []int) (string, string) {
	var s1, s2 string
	for _, v := range n1 {
		s1 += strconv.Itoa(v)
	}

	for _, v := range n2 {
		s2 += strconv.Itoa(v)
	}
	s2 = strings.TrimLeft(s2, "0")

	return s1, s2
}
