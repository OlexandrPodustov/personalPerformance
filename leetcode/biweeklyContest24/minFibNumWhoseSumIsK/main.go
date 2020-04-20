package main

import "fmt"

func main() {
	// in := 4
	in := 19
	res := findMinFibonacciNumbers(in)
	_ = res
	fmt.Println(res)
}

func findMinFibonacciNumbers(k int) int {
	fibSeq := []int{1, 1}

	n := 0
	for n <= k {
		n = fibSeq[len(fibSeq)-1] + fibSeq[len(fibSeq)-2]
		fibSeq = append(fibSeq, n)
	}

	res := 0
	for i := len(fibSeq) - 1; i >= 0; i-- {
		curEl := fibSeq[i]
		if curEl == k {
			res++
			break
		}

		if k-curEl > 0 {
			res++
			k -= curEl
		}

	}

	return res
}
