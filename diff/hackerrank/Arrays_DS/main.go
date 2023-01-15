package main

import (
	"container/list"
	"fmt"
	"log"
)

func main() {
	var numbersAmount int
	if _, err := fmt.Scan(&numbersAmount); err != nil {
		log.Fatal("parse", err)
	}

	inputData := make([]int32, 0, numbersAmount)
	for i := 1; i <= numbersAmount; i++ {
		var number int32
		if _, err := fmt.Scan(&number); err != nil {
			log.Fatal(i, err)
		}
		inputData = append(inputData, number)
	}

	// fmt.Println(inputData)
	rd := reverseArray(inputData)
	fmt.Println(rd)
}

func reverseArray(a []int32) []int32 {
	// container/list
	lst := list.New()
	lst.Init()

	mid := len(a) / 2
	i := 0
	for j := len(a) - 1; j >= mid; j-- {
		a[i], a[j] = a[j], a[i]
		i++
	}

	return a
}
