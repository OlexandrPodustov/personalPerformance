package main

import (
	"fmt"
	"log"
)

func main() {
	var testCases int
	if _, err := fmt.Scanln(&testCases); err != nil {
		log.Fatal("testCases err", err)
	}

	for caseNumber := 1; caseNumber <= testCases; caseNumber++ {

		var bags, kids int
		if _, err := fmt.Scan(&bags, &kids); err != nil {
			log.Fatal("fmt.Scan bags", err)
		}
		var buckets []int
		for i := 0; i < bags; i++ {
			var bucket int
			if _, err := fmt.Scan(&bucket); err != nil {
				log.Fatal("fmt.Scan bucket", err)
			}
			buckets = append(buckets, bucket)
		}
		res := solution(kids, buckets)
		fmt.Printf("Case #%d: %v\n", caseNumber, res)
	}
}

func solution(kids int, buckets []int) int {
	candies := 0
	for _, v := range buckets {
		candies += v
	}
	return candies % kids
}
