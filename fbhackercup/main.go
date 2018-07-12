package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var t int

	_, err := fmt.Scanln(&t)
	if err != nil {
		fmt.Println("fmt.Scanln(&t)", err)
	}
	fmt.Println("campuses amount:", t)

	for caseNumber := 1; caseNumber <= t; caseNumber++ {
		var n, k, v int

		_, err := fmt.Scanln(&n, &k, &v)
		if err != nil {
			fmt.Println("fmt.Scanln - n,k,v - err: ", err)

		}
		fmt.Printf("#"+strconv.Itoa(caseNumber)+" N:%v, K:%v, V:%v \n", n, k, v)

		campusesOriginalOrder := make([]string, n)
		campus := ""
		for i := 0; i < n; i++ {
			_, err := fmt.Scanln(&campus)
			if err != nil {
				fmt.Println("fmt.Scanln - n,k,v - err: ", err)
			}
			campusesOriginalOrder[i] = campus
		}
		//fmt.Println("final result campusesOriginalOrder - ", campusesOriginalOrder)

		if n == k {
			fmt.Println("Case #"+strconv.Itoa(caseNumber)+":", strings.Join(campusesOriginalOrder, " "))
			continue
		}

		countCampus := make(map[string]int, n)

		for ii := 1; ii <= v; ii++ {
			kk := k
			fmt.Println("visit #", ii)
			for i, v := range campusesOriginalOrder {
				// sort slice from previous iteration in increasing order of visits
				if i < 2 {
					countCampus[v]++
					kk--
				}
				if kk == 0 {
					break
				}
				fmt.Println("countCampus", countCampus)

			}

		}

		//final result
		fmt.Println("Case #"+strconv.Itoa(caseNumber)+":", strings.Join(campusesOriginalOrder, " "))
	}

}
