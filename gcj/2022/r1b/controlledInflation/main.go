package main

import (
	"fmt"
	"log"
)

func main() {
	var testCasesAmount int
	if _, err := fmt.Scan(&testCasesAmount); err != nil {
		log.Fatal("parse testCasesAmount", err)
	}
	for testCaseNumber := 1; testCaseNumber <= testCasesAmount; testCaseNumber++ {
		var customers, productsAmount int
		if _, err := fmt.Scan(&customers, &productsAmount); err != nil {
			log.Fatal(testCaseNumber, "", err)
		}

		customersWithProducts := make([][]int, 0, customers)
		for i := 0; i < customers; i++ {
			exerc := make([]int, 0, productsAmount)
			for j := 0; j < productsAmount; j++ {
				var numb int
				if _, err := fmt.Scan(&numb); err != nil {
					log.Fatal(testCaseNumber, err)
				}
				exerc = append(exerc, numb)
			}
			customersWithProducts = append(customersWithProducts, exerc)
		}

		result := solve(customersWithProducts)
		fmt.Printf("Case #%v: %v\n", testCaseNumber, result)
	}
}

func solve(customersWithProducts [][]int) int64 {
	fmt.Println(customersWithProducts)
	var buttonsPressed int64

	_, maxVal := findMinMax(customersWithProducts[0])
	// fmt.Println("maxVal", maxVal)
	buttonsPressed += maxVal

	newProductsList := getNewList(maxVal, customersWithProducts[1:])
	fmt.Println("newProductsList", newProductsList)

	for _, v := range newProductsList {
		df := diff(maxVal, v)
		buttonsPressed += df
		maxVal = v
	}
	// fmt.Println("buttonsPressed", buttonsPressed)

	return buttonsPressed
}

func findMinMax(numbers []int) (int64, int64) {
	min := 4294967295
	max := 0

	for _, v := range numbers {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return int64(min), int64(max)
}

func getNewList(lastVal int64, customersWithProducts [][]int) []int64 {
	newProductsList := make([]int64, 0, len(customersWithProducts))
	for i := range customersWithProducts {
		products := customersWithProducts[i]
		currMinVal, currMaxVal := findMinMax(products)
		var nextMinVal, nextMaxVal int64
		if i+1 < len(customersWithProducts) {
			nextMinVal, nextMaxVal = findMinMax(customersWithProducts[i+1])
		}

		first, second := decideOrder(lastVal, currMinVal, currMaxVal, nextMinVal, nextMaxVal)
		newProductsList = append(newProductsList, first, second)

		lastVal = second
	}

	return newProductsList
}

func decideOrder(lastVal, min, max, nextMinVal, nextMaxVal int64) (int64, int64) {
	fmt.Println("lastVal, min, max", lastVal, min, max)

	var first, second int64
	switch {
	case diff(lastVal, min) > diff(lastVal, max):
		first, second = max, min
	case diff(lastVal, min) < diff(lastVal, max):
		first, second = min, max
	case diff(lastVal, min) == diff(lastVal, max):
		// nextFirst, nextSecond = decideOrder(lastVal, nextMinVal, nextMaxVal, 0, 0)
	}

	fmt.Println("first, second", first, second)

	return first, second
}

func diff(a, b int64) int64 {
	if a >= b {
		return a - b
	}
	return b - a
}
