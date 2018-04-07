package main

//
//import (
//	"fmt"
//	"log"
//	"os"
//)
//
//func main() {
//	// A - prepared cells. eg = 11
//	// T -
//	var matrix = make([][]int, 1000)
//	fmt.Println("m:", matrix)
//
//	fmt.Print("dig here:")
//	var t int
//	_, err := fmt.Scanln(&t)
//	if err != nil {
//		log.Println("fmt.Scanln(&t)", err)
//	}
//	// scan line with n
//	var a, b, n int
//	fmt.Print("amount of attempts to guess:")
//	_, err = fmt.Scanln(&n)
//	if err != nil {
//		log.Println("fmt.Scanln(&n):", err)
//	}
//	//log.Println("amount of attempts to guess given: ", n)
//
//	// try to guess
//
//	//for i := a; i <= b; i++ {
//	for i := 0; i < n; i++ {
//		log.Println(i, "th attempt ", a, b)
//		var val = (a + 1 + b) / 2 // exclusive lower bound
//		fmt.Println(val)
//
//		// coordinate following flow
//		var answer string
//		fmt.Print("server answer:")
//		_, err = fmt.Scanln(&answer)
//		if err != nil {
//			log.Println("fmt.Scanln(&answer)", err)
//		}
//		//log.Println("answer for coordination:", answer)
//
//		switch {
//		case answer == "WRONG_ANSWER":
//			os.Exit(1)
//		case answer == "TOO_SMALL":
//			a = val + 1
//		case answer == "TOO_BIG":
//			b = val - 1
//		case answer == "CORRECT":
//			break
//		default:
//			log.Println("default case")
//		}
//	}
//}
