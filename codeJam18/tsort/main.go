package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner, tc := openF()

	for i := 0; i < tc; i++ {
		sliceToSort := loopTC(scanner)

		sliceToCheck := tSort(sliceToSort)

		res := checkAscending(sliceToCheck)

		fmt.Printf("Case #%v: %v\n", i+1, res)
	}
}

func checkAscending(inSlice []byte) string {

	return "OK"
}

func openF() (*bufio.Scanner, int) {
	var tcAmnt int

	file, err := os.Open("tst.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		aNumber := scanner.Text()
		tcAmnt, err = strconv.Atoi(aNumber)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("tc amount", tcAmnt)
	}

	return scanner, tcAmnt
}

func loopTC(scanner *bufio.Scanner) []byte {
	scanner.Scan()
	ds := scanner.Text()
	size, err := strconv.Atoi(ds)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("dataSet size", size)
	scanner.Scan()
	slsl := scanner.Bytes()
	fmt.Printf("data to sort:%q \n", slsl)

	return slsl
}

func tSort(L []byte) []byte {
	var res []byte

	for i := 0; i < len(L)-2; i++ {
		if L[i] > L[i+2] {

		}
	}

	return res
}
