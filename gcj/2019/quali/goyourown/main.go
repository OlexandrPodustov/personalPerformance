package main

import (
	"fmt"
	"log"
)

const (
	e = "E"
	s = "S"
)

func main() {
	var iterations int
	_, err := fmt.Scan(&iterations)
	if err != nil {
		log.Println("failed to parse testCasesAmount:", err)
	}

	for ii := 1; ii <= iterations; ii++ {
		var dimensions int
		_, err := fmt.Scan(&dimensions)
		if err != nil {
			log.Println(ii, " failed to parse input:", err)
		}

		var lydaPath string
		_, err = fmt.Scan(&lydaPath)
		if err != nil {
			log.Println(ii, " failed to parse input:", err)
		}

		// if len(lydaPath) != 2*dimensions-2 {
		// 	log.Fatalln("incorrect input", len(lydaPath), 2*dimensions-2)
		// }

		result := move(dimensions, lydaPath)
		fmt.Printf("Case #%v: %v\n", ii, result)
	}
}

func move(dimensions int, lydaPath string) string {
	result := ""

	var eee, sss int
	for _, v := range lydaPath {
		// if i == 0 {
		// 	if string(v) == s {
		// 		eee++
		// 		result += e
		// 	} else {
		// 		sss++
		// 		result += s
		// 	}
		// 	continue
		// }
		if eee < sss {
			if string(v) != e {
				eee++
				result += e
			} else {
				sss++
				result += s
			}
		} else {
			if string(v) != s {
				sss++
				result += s
			} else {
				eee++
				result += e
			}
		}

		// if eee > dimensions || sss > dimensions {
		// 	fmt.Println("AAA:", eee, sss)
		// }
		// if len(result) > len(lydaPath) {
		// 	fmt.Println("AAA len(result) > len(lydaPath):", len(result), len(lydaPath))
		// }
	}
	// fmt.Println("eee:", eee)
	// fmt.Println("sss:", sss)

	return result
}
