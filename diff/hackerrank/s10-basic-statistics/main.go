package main

import (
	"fmt"
	"log"
	"math"
	"sort"
)

func main() {
	var numbersAmount int
	if _, err := fmt.Scan(&numbersAmount); err != nil {
		log.Fatal("parse", err)
	}

	modes := make(map[float64]int, 0)
	var totalSum float64
	var totalAmount float64
	medianData := make([]float64, 0, numbersAmount)
	for i := 1; i <= numbersAmount; i++ {
		var number float64
		if _, err := fmt.Scan(&number); err != nil {
			log.Fatal(i, err)
		}
		totalSum += number
		totalAmount++
		modes[number]++
		medianData = append(medianData, number)
	}

	mean := totalSum / totalAmount
	fmt.Printf("%.1f\n", mean)

	median := findMedian(medianData)
	fmt.Printf("%.1f\n", median)

	mode := findMode(modes)
	fmt.Println(mode)
}

func findMode(list map[float64]int) float64 {
	mostFrequent := 0
	modeList := make(map[int][]float64, 0)
	for key, v := range list {
		if v > mostFrequent {
			mostFrequent = v
		}
		modeList[v] = append(modeList[v], key)
	}

	mode := findMin(modeList[mostFrequent])

	return mode
}

func findMin(modeList []float64) float64 {
	minIfAllAreEquallyFrequent := math.MaxFloat64
	for _, v := range modeList {
		if v < minIfAllAreEquallyFrequent {
			minIfAllAreEquallyFrequent = v
		}
	}

	return minIfAllAreEquallyFrequent
}

func findMedian(medianData []float64) float64 {
	sort.SliceStable(medianData, func(i, j int) bool {
		return medianData[i] < medianData[j]
	})
	middle := len(medianData) / 2
	if len(medianData)%2 != 0 {
		return medianData[middle]
	}
	median := (medianData[middle-1] + medianData[middle]) / 2

	return median
}
