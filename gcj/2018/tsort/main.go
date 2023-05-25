package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var tc int
	_, err := fmt.Scanln(&tc)
	if err != nil {
		log.Println("fmt.Scanln(&t)", err)
	}

	for i := 0; i < tc; i++ {
		sliceToSort := loopTC()

		sliceToCheck := tSort(sliceToSort)

		res := checkAscending(sliceToCheck)

		fmt.Printf("Case #%v: %v\n", i+1, res)
	}
}

func tSort(res []int) []int {
	var done bool
	for !done {
		done = true
		for i := 0; i < len(res)-2; i++ {
			if res[i] > res[i+2] {
				res[i], res[i+2] = res[i+2], res[i]
				done = false
			}
		}
	}

	return res
}

func checkAscending(intSlice []int) string {
	var done bool
	for !done {
		done = true
		for i := 0; i < len(intSlice)-1; i++ {
			if intSlice[i] > intSlice[i+1] {
				done = false
				return strconv.Itoa(i + 1)
			}
		}
	}

	return "OK"
}

func loopTC() []int {
	var a int
	_, err := fmt.Scan(&a)
	if err != nil {
		log.Println("fmt.Scanln(&anb)", err)
	}
	var res []int
	for i := 0; i < a; i++ {
		var r int
		_, err = fmt.Scan(&r)
		if err != nil {
			log.Println("fmt.Scanln(&n):", err)
		}
		res = append(res, r)
	}

	return res
}
