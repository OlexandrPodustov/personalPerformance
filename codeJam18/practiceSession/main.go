package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// initial read of test cases amount
	var (
		t    int
		corr = true
	)
	fmt.Print("testCasesAmount:")
	_, err := fmt.Scanln(&t)
	if err != nil {
		log.Println("fmt.Scanln(&t)", err)
	}
	//log.Println("testCasesAmount:", t)

	for i := 0; i < t; i++ {
		var a, b, n int

		// After N tries, if the test case is not solved,
		// the judge will print WRONG_ANSWER and
		// then stop sending output to your input stream.
		if !corr {
			var wrongAnswer string
			_, err = fmt.Scanln(&wrongAnswer)
			if err != nil {
				log.Println("fmt.Scanln(&answer)", err)
			}
			if wrongAnswer == "WRONG_ANSWER" {
				os.Exit(2)
			}
		}

		// scan line with a and b
		fmt.Print("input a and b:")
		_, err := fmt.Scan(&a, &b)
		if err != nil {
			log.Println("fmt.Scanln(&anb)", err)
		}
		//log.Println("a and b given: ", a, b)

		// scan line with n
		fmt.Print("amount of attempts to guess:")
		_, err = fmt.Scanln(&n)
		if err != nil {
			log.Println("fmt.Scanln(&n):", err)
		}
		//log.Println("amount of attempts to guess given: ", n)

		// try to guess

		//for i := a; i <= b; i++ {
		corr = false
		for i := 0; i < n; i++ {
			if corr {
				break
			}
			log.Println(i, "th attempt ", a, b)
			var val = (a + 1 + b) / 2 // exclusive lower bound
			fmt.Println(val)

			// coordinate following flow
			var answer string
			fmt.Print("server answer:")
			_, err = fmt.Scanln(&answer)
			if err != nil {
				log.Println("fmt.Scanln(&answer)", err)
			}
			//log.Println("answer for coordination:", answer)

			switch {
			case answer == "WRONG_ANSWER":
				os.Exit(1)
			case answer == "TOO_SMALL":
				a = val + 1
			case answer == "TOO_BIG":
				b = val - 1
			case answer == "CORRECT":
				corr = true
				break
			default:
				log.Println("default case")
			}
		}
	}
}
