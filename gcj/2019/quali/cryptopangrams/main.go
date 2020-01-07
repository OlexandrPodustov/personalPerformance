package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var iterations int
	_, err := fmt.Scan(&iterations)
	if err != nil {
		log.Println("failed to parse testCasesAmount:", err)
	}

	for ii := 1; ii <= iterations; ii++ {
		var n, l int
		_, err = fmt.Scan(&n, &l)
		if err != nil {
			log.Println(ii, " failed to parse input:", err)
		}

		cypherSlice := make([]int, l)
		for i := 0; i < l; i++ {
			val := 0
			_, err = fmt.Scan(&val)
			if err != nil {
				log.Println(ii, " failed to parse input:", err)
			}
			cypherSlice[i] = val
		}
		// fmt.Println("len cypherSlice: ", len(cypherSlice))

		alph := genAlph(n)
		// fmt.Println("map: ", alph)

		result := decipher(cypherSlice, alph)
		fmt.Printf("Case #%v: %v\n", ii, result)
	}
}

// 2
// 103 31
// 217 1891 4819 2291 2987 3811 1739 2491 4717 445 65 1079 8383 5353 901 187 649 1003 697 3239 7663 291 123 779 1007 3551 1943 2117 1679 989 3053
// 10000 25
// 3292937 175597 18779 50429 375469 1651121 2102 3722 2376497 611683 489059 2328901 3150061 829981 421301 76409 38477 291931 730241 959821 1664197 3057407 4267589 4729181 5335543

func decipher(input []int, alphabetMap map[int]string) string {
	result := ""

	var lastInt int
	if len(input) > 1 {
		var a, b, c, d int
		i := 0
		// fmt.Println("input: ", input[i], input[i+1])
		for v := range alphabetMap {
			// fmt.Println("try prime from map: ", v)

			if got := input[i] / v; (got)*v == input[i] {
				// fmt.Println(input[i], "divided by: ", v, " got ", got)
				a = v
				b = got
				break
			}
		}

		for v := range alphabetMap {
			if got := input[i+1] / v; (got)*v == input[i+1] {
				// fmt.Println(input[i+1], "divided by: ", v, " got ", got)
				c = v
				d = got
				break
			}
		}

		if a == c || a == d {
			midMan := a
			a = b
			b = midMan
			if a == c {
				c = d
			}
		} else if b == c || b == d {
			if b == c {
				c = d
			}
		} else {
			log.Fatal("fuck: ", a, b, c, d)
		}

		// fmt.Println("decipher a", len(input))
		input = input[2:]
		// fmt.Println("decipher b", len(input))
		// fmt.Println("res --- ", alphabetMap[a], alphabetMap[b], alphabetMap[c])

		result += fmt.Sprintf("%v%v%v", alphabetMap[a], alphabetMap[b], alphabetMap[c])
		lastInt = c
	}

	for i := 0; i < len(input); i++ {
		// fmt.Println("input: ", input[i])
		// fmt.Println("try prime from map: ", v)

		if got := input[i] / lastInt; (got)*lastInt == input[i] {
			// fmt.Println(input[i], "divided by: ", lastInt, " got ", got)

			result += fmt.Sprintf("%v", alphabetMap[got])
			lastInt = got
		}
	}

	return result
}

func genAlph(max int) map[int]string {
	alph := make(map[int]string, 26)
	aSlice := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	primeNumbers := SieveOfEratosthenes(max + 1)
	// fmt.Println("primeNumbers len: ", len(primeNumbers))
	for i := len(primeNumbers) - 1; i >= 0 && len(aSlice) > 0; i-- {
		// fmt.Println(aSlice[len(aSlice)-1], primeNumbers[i])

		alph[primeNumbers[i]] = aSlice[len(aSlice)-1]
		aSlice = aSlice[:len(aSlice)-1]
	}
	// fmt.Println("maplen: ", len(alph))

	return alph
}

func SieveOfEratosthenes(value int) []int {
	f := make([]bool, value)
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if !f[i] {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	result := []int{}
	for i := 3; i < value; i++ {
		if !f[i] {
			// fmt.Printf("%v ", i)
			result = append(result, i)
		}
	}
	// fmt.Println("")

	return result
}
