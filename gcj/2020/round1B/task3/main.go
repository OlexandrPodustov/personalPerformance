package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func main() {
	var iterations int
	_, err := fmt.Scan(&iterations)
	if err != nil {
		log.Println("parse testCasesAmount:", err)
		return
	}

	for testCaseNumber := 1; testCaseNumber <= iterations; testCaseNumber++ {
		var rank, suit int
		_, err := fmt.Scan(&rank, &suit)
		if err != nil {
			log.Println(testCaseNumber, "  numberOfPatternsToMatch:", err)
			return
		}

		result, ms := check(rank, suit)
		fmt.Printf("Case #%v: %v\n%v\n", testCaseNumber, result, strings.Join(ms, "\n"))
	}
}

func check(rank, suit int) (int, []string) {
	ussl := []int{}
	// ussl := [][]int{}
	for s := 1; s <= suit; s++ {
		for r := 1; r <= rank; r++ {
			// ussl = append(ussl, []int{r, s})
			ussl = append(ussl, r)
		}

	}
	fmt.Println()
	fmt.Println("ussl", ussl)

	movesAmount := 0
	moves := []string{}
	for !sort.IntsAreSorted(ussl) {
		if rank >= suit && rank > 1 {
			a, b := rank, rank-1
			moves = append(moves, fmt.Sprintf("%v %v", a, b))

			partA := make([]int, len(ussl[:a]))
			partB := make([]int, len(ussl[a:a+b]))
			partC := make([]int, len(ussl[a+b:]))
			copy(partA, ussl[:a])
			copy(partB, ussl[a:a+b])
			copy(partC, ussl[a+b:])

			ussl = nil
			ussl = append(ussl, partB...)
			fmt.Println("ussl+b", ussl)
			ussl = append(partB, partA...)
			fmt.Println("ussl+a", ussl)
			ussl = append(ussl, partC...)
			fmt.Println("ussl+c", ussl)

			rank--
			suit--
		} else if suit > rank && rank > 1 {
			a, b := rank, rank+1
			moves = append(moves, fmt.Sprintf("%v %v", a, b))

			partA := make([]int, len(ussl[:a]))
			partB := make([]int, len(ussl[a:a+b]))
			partC := make([]int, len(ussl[a+b:]))
			copy(partA, ussl[:a])
			copy(partB, ussl[a:a+b])
			copy(partC, ussl[a+b:])

			ussl = nil
			ussl = append(ussl, partB...)
			// fmt.Println("e ussl+b", ussl)
			ussl = append(partB, partA...)
			// fmt.Println("e ussl+a", ussl)
			ussl = append(ussl, partC...)
			// fmt.Println("e ussl+c", ussl)

			rank--
			suit--
		} else if suit > rank && rank == 1 {
			a, b := rank+1, rank+1
			moves = append(moves, fmt.Sprintf("%v %v", a, b))

			partA := make([]int, len(ussl[:rank]))
			partB := make([]int, len(ussl[rank:rank+rank]))
			partC := make([]int, len(ussl[rank+rank:]))

			copy(partA, ussl[:rank])
			copy(partB, ussl[rank:rank+rank])
			copy(partC, ussl[rank+rank:])

			ussl = nil
			ussl = append(ussl, partB...)
			// fmt.Println("= ussl+b", ussl)
			ussl = append(partB, partA...)
			// fmt.Println("= ussl+a", ussl)
			ussl = append(ussl, partC...)
			// fmt.Println("= ussl+c", ussl)

			rank--
			suit--
		}

		if movesAmount > 100 {
			break
		}
		movesAmount++
	}

	fmt.Println("final ussl", ussl)
	fmt.Println("movesAmount, moves", movesAmount, moves)

	return movesAmount, moves
}
