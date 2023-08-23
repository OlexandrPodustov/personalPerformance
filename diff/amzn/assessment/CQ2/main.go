package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const modlo = 1000000007

func main() {
	oneKibibyte := 1024
	reader := bufio.NewReaderSize(os.Stdin, 16*oneKibibyte*oneKibibyte)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer checkError(stdout.Close())

	powerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 0)
	checkError(err)

	var power []int32
	for i := 0; i < int(powerCount); i++ {
		powerItem, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 32)
		checkError(err)
		power = append(power, int32(powerItem))
	}

	writer := bufio.NewWriterSize(stdout, 16*oneKibibyte*oneKibibyte)
	fmt.Fprintf(writer, "%d\n", findTotalPower(power))

	checkError(writer.Flush())
}

func findTotalPower(power []int32) int32 {
	fmt.Println("len power", len(power))
	lenPower := len(power)

	results := make([]uint64, 0, lenPower)
	for i := 0; i < lenPower; i++ {
		for j := 0; j < lenPower; j++ {
			if i <= j {
				results = append(results, getMinSum(power[i:j+1]))
			}
		}
	}

	var overallSum uint64
	for _, v := range results {
		overallSum += v
	}

	osm := overallSum % modlo
	overallSum32 := int32(osm)
	fmt.Println("overallSum32", overallSum, osm, overallSum32)

	return overallSum32
}

func getMinSum(powerChunk []int32) uint64 {
	min := uint64(math.MaxUint64)
	var localSum uint64
	for i := 0; i < len(powerChunk); i++ {
		cur := uint64(powerChunk[i])
		localSum += cur
		if cur < min {
			min = cur
		}
	}

	return min * localSum
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
