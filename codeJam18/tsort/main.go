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
				//fmt.Println("bef", res)
				res[i], res[i+2] = res[i+2], res[i]
				//fmt.Println("after: ", res)
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
				//intSlice[i], intSlice[i+1] = intSlice[i+1], intSlice[i]
				//fmt.Println("after 2: ", intSlice)
				done = false
				//fmt.Println("wrong index", i)
				return strconv.Itoa(i)
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
