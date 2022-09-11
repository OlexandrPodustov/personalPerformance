package main

import (
	"fmt"
	"log"
	"math/big"
	"runtime"
	"sync"
)

func main() {
	var testCases int
	if _, err := fmt.Scanln(&testCases); err != nil {
		log.Fatal("fmt.Scanln(&t) testCases err", err)
	}
	fmt.Printf("GOMAXPROCS  %v\n", runtime.GOMAXPROCS(100))
	var wg sync.WaitGroup
	resChan := make(chan int64, testCases)

	for caseNumber := 1; caseNumber <= testCases; caseNumber++ {
		wg.Add(1)
		var treesAmount int
		if _, err := fmt.Scan(&treesAmount); err != nil {
			log.Fatal("fmt.Scanln - err: ", err)
		}

		treesLocations := make([][]int64, 0, treesAmount)
		for i := 0; i < treesAmount; i++ {
			var itemA, itemB int64
			if _, err := fmt.Scan(&itemA, &itemB); err != nil {
				log.Fatal("Scan: ", err)
			}

			treesLocations = append(treesLocations, []int64{itemA, itemB})
		}
		// fmt.Printf("treesLocations %v\n", treesLocations)

		var wellLocationsAmount int
		if _, err := fmt.Scan(&wellLocationsAmount); err != nil {
			log.Fatal("fmt.Scanln - err: ", err)
		}
		fmt.Printf("testcaseNumber treesAmount, wellLocationsAmount %v %v %v\n", caseNumber, treesAmount, wellLocationsAmount)

		wellLocations := make([][]int64, 0, wellLocationsAmount)
		for i := 0; i < wellLocationsAmount; i++ {
			var itemA, itemB int64
			if _, err := fmt.Scan(&itemA, &itemB); err != nil {
				log.Fatal("Scan: ", err)
			}

			wellLocations = append(wellLocations, []int64{itemA, itemB})
		}
		// fmt.Printf("wellLocations %v\n", wellLocations)

		go solution(treesLocations, wellLocations, resChan)
	}
	wg.Wait()
	close(resChan)

	caseNumber := 1
	for res := range resChan {
		fmt.Printf("Case #%d: %v\n", caseNumber, res)
		caseNumber++
	}
}

const (
	modlo   = 1000000007
	maxABXY = 999999999
)

func solution(treesLocations, wellLocations [][]int64, resssss chan int64) int64 {
	bigResult := big.NewInt(0)

	for i := range wellLocations {
		// fmt.Printf("\twell  %v\n", well)
		for j := range treesLocations {
			// fmt.Printf("\t\t tree %v\n", tree)

			a := wellLocations[i][0] - treesLocations[j][0]
			b := wellLocations[i][1] - treesLocations[j][1]

			abig := big.NewInt(a * a)
			bb := big.NewInt(b * b)

			bigResult.Add(bigResult, abig)
			bigResult.Add(bigResult, bb)

			// fmt.Printf("\t\t abig %v\n", abig)
			// fmt.Printf("\t\t bb %v\n", bb)
			// fmt.Printf("\t\t bigResult %v\n", bigResult)
		}
		// fmt.Printf("\twell result %v %v\n", well, result)
	}

	modBigInt := big.NewInt(modlo)
	br := bigResult.Mod(bigResult, modBigInt)
	// fmt.Printf("bigResult %v\n", br)
	// fmt.Printf("br.Int64() %v\n", br.Int64())

	brr := br.Int64()
	resssss <- brr
	// select {
	// case resssss <- brr:
	// default:
	// }

	return brr
}
