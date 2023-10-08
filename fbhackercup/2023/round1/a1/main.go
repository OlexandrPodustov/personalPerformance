// Here Comes Santa Claus
package main

import (
	"fmt"
	"log"
	"sort"
)

func main() {
	var testCases int
	if _, err := fmt.Scanln(&testCases); err != nil {
		log.Fatal("fmt.Scanln(&t) testCases err", err)
	}

	for caseNumber := 1; caseNumber <= testCases; caseNumber++ {
		var elvesAmount int
		if _, err := fmt.Scan(&elvesAmount); err != nil {
			log.Fatal("fmt.Scanln - err: ", err)
		}

		elves := make([]int, 0, elvesAmount)
		for i := 0; i < elvesAmount; i++ {
			var elvePlace int
			if _, err := fmt.Scan(&elvePlace); err != nil {
				log.Fatal("fmt.Scanln - err: ", err)
			}

			elves = append(elves, elvePlace)
		}
		sort.Ints(elves)

		var res int
		first := elves[0]
		second := elves[1]
		preLast := elves[len(elves)-2]
		last := elves[len(elves)-1]

		// var haveRemainder bool

		if elvesAmount == 5 {
			// take three elves on the side where they are further
			third := elves[2]

			fd := third - first
			ld := last - third

			if fd > ld {
				var hr1, hr2 bool

				if (third-first)%2 != 0 {
					hr1 = true
				}
				if (last-preLast)%2 != 0 {
					hr2 = true
				}
				if hr1 != hr2 {
					// haveRemainder = true
					// fmt.Println("hr1, hr2, haveRemainder", hr1, hr2, haveRemainder)
				}

				fp := (third-first)/2 + first
				lp := (last-preLast)/2 + preLast

				// fmt.Println("fd > ld first three", third, first, fp)
				// fmt.Println("fd > ld last two", last, preLast, lp)
				res = lp - fp
			} else {
				var hr1, hr2 bool
				if (second-first)%2 != 0 {
					hr1 = true
				}
				if (last-third)%2 != 0 {
					hr2 = true
				}
				if hr1 != hr2 {
					// haveRemainder = true
					// fmt.Println("else hr1, hr2, haveRemainder", hr1, hr2, haveRemainder)
				}

				fp := (second-first)/2 + first
				lp := (last-third)/2 + third

				// fmt.Println("fd <= ld first two", fp, elves)
				// fmt.Println("fd <= ld last three", lp, elves)
				res = lp - fp
			}
		} else {
			fp := (second-first)/2 + first
			lp := (last-preLast)/2 + preLast
			// fmt.Println("else fp", first, second, fp)
			// fmt.Println("else lp", last, preLast, lp)

			res = lp - fp
		}

		fmt.Printf("Case #%d: %v\n", caseNumber, res)
	}
}
