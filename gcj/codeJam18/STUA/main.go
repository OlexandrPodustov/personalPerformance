package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var tCasesAmount int
	var caseNumber int

	//fmt.Print("enter tc am:")
	_, err := fmt.Scan(&tCasesAmount)
	if err != nil {
		log.Fatalln("fmt.Scanln(&anb)", err)
	}

	for i := 0; i < tCasesAmount; i++ {
		//fmt.Print("n and str:")
		var r int
		_, err = fmt.Scan(&r)
		if err != nil {
			log.Fatalln("fmt.Scanln(&n):", err)
		}
		//fmt.Println("n", r)

		var st string
		_, err = fmt.Scan(&st)
		if err != nil {
			log.Fatalln("fmt.Scanln(&n):", err)
		}
		//fmt.Println("st", st)
		caseNumber++
		var sliceB []byte
		for _, v := range st {
			sliceB = append(sliceB, byte(v))
		}
		result := calc(uint(r), &sliceB)
		fmt.Printf("Case #%v: %v\n", caseNumber, result)
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
	for _, v := range *input {
		switch v {
		case 'C':
			damage *= 2
			//fmt.Println("C - charge: ", damage)
		case 'S':
			minTotalDamage += damage
			//fmt.Println("S - shoot: ", damage)
		default:
			//fmt.Println("\t default case: ", i, v)
		}
	}

	return minTotalDamage
}
