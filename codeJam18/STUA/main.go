package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		caseNumber int
	)

	file, err := os.Open("tst.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		readL := scanner.Bytes()
		slsl := bytes.Split(readL, []byte(" "))
		//fmt.Printf("%#v \n", len(slsl))
		//if len(slsl) == 1 {
		//	fmt.Println(string(slsl[0]))
		//	testCasesAmount, err := strconv.Atoi(string(slsl[0]))
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//}
		if len(slsl) == 2 {
			caseNumber++
			maxWithstand, err := strconv.Atoi(string(slsl[0]))
			if err != nil {
				log.Fatal(err)
			}
			result := calc(uint(maxWithstand), &slsl[1])
			fmt.Printf("Case #%v: %v\n", caseNumber, result)
		}
		if len(slsl) > 2 {
			log.Fatal("\t len(slsl) > 2 - unexxpected string length")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func calc(maxWithstand uint, input *[]byte) string {
	var (
		minTotalDamage uint
	)
	*input = bytes.TrimRight(*input, `C`)
	if len(*input) == 0 {
		//fmt.Println("\t\t\t robot will never shoot")
		return "0"
	}

	minTotalDamage = countDamage(input)

	if minTotalDamage <= maxWithstand {
		//fmt.Println("\t\t minTotalDamage <= maxWithstand. we don't need to swap")
		return "0"
	}

	var iterationsToDeactivate int
	for {
		if minTotalDamage <= maxWithstand {
			return strconv.Itoa(iterationsToDeactivate)
		}

		combinationIndex := strings.LastIndex(string(*input), "CS")
		if combinationIndex == -1 {
			break
		}

		iterationsToDeactivate++
		(*input)[combinationIndex], (*input)[combinationIndex+1] = (*input)[combinationIndex+1], (*input)[combinationIndex]
		*input = bytes.TrimRight(*input, `C`)
		//fmt.Println("\t new string(input):", string(input))
		minTotalDamage = countDamage(input)
	}

	if minTotalDamage > maxWithstand {
		//fmt.Println("\t IMPOSSIBLE input:", string(input))
		return "IMPOSSIBLE"
	}

	return strconv.Itoa(iterationsToDeactivate)
}

func countDamage(input *[]byte) uint {
	var damage uint = 1
	var minTotalDamage uint
	for i, v := range *input {
		switch v {
		case 'C':
			damage = damage * 2
			//fmt.Println("C - charge: ", damage)
		case 'S':
			minTotalDamage += damage
			//fmt.Println("S - shoot: ", damage)
		default:
			fmt.Println("\t default case: ", i, v)
		}
	}

	return minTotalDamage
}
