package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var (
		//testCasesAmount int
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
			result := calc(uint(maxWithstand), slsl[1])
			fmt.Println("Case #", caseNumber, ":", result)
		}
		if len(slsl) > 2 {
			log.Fatal("\t len(slsl) > 2 - unexxpected string length")
		}
		//fmt.Printf("%#v \n", slsl)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("1<<2: %v \n", 1<<1)
}

func calc(maxWithstand uint, input []byte) string {
	var (
		minTotalDamage uint
		damage         uint = 1
	)
	//fmt.Println("input b", input)
	input = bytes.TrimRight(input, `C`)
	//fmt.Println("input a", input)
	if len(input) == 0 {
		//fmt.Println("\t\t\trobot will never shoot")
		return "0"
	}

	fmt.Printf("\t maxWithstand:%v [%q]\n", maxWithstand, input)
	for i, v := range input {
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
	//fmt.Println("Min total damage without swap: ", minTotalDamage)

	if minTotalDamage <= maxWithstand {
		//fmt.Println("\t\t minTotalDamage <= maxWithstand. we don't need to swap")
		return "0"
	}

	// swap
	var amountOfiterationsToDeactivate int
	for j := len(input) - 1; j > 0; j-- {
		if minTotalDamage <= maxWithstand {
			//fmt.Println("\t\t minTotalDamage <= maxWithstand. stop swapping. minTotalDamage:", strconv.Itoa(int(minTotalDamage)))
			return strconv.Itoa(amountOfiterationsToDeactivate)
		}

		if input[j-1] < input[j] {
			amountOfiterationsToDeactivate++
			//fmt.Println("input j:", j-1, j, string(input[j-1]), string(input[j]), input[j-1], input[j])
			//input[j], input[j-1] = input[j-1], input[j]
			//fmt.Println("\tinput:", string(input))
			damage -= damage / 2
			minTotalDamage -= damage
		}
	}

	//swap the whole array
	l := len(input)
	middle := l / 2
	for i := 0; i < middle; i++ {
		l--
		input[i], input[l] = input[l], input[i]
	}

	// find the first occurrence of SC though our slice is swapped
	fmt.Println("\t string(input)", string(input))
	combinationPresent := bytes.IndexAny(input, "SC")
	fmt.Println("\t combinationPresent", combinationPresent)
	//for {
	//if found substr:
	// {
	//		amountOfiterationsToDeactivate++
	//		//input[j], input[j-1] = input[j-1], input[j]
	//		fmt.Println("input:", string(input))
	//		damage -= damage / 2
	//		minTotalDamage -= damage
	//
	//	}
	//if substr is absent - break
	//{
	// break
	// }
	//}

	if minTotalDamage > maxWithstand {
		fmt.Println("IMPOSSIBLE input:", string(input))
		return "IMPOSSIBLE"
	}

	return strconv.Itoa(amountOfiterationsToDeactivate)
}
