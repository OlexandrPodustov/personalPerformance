package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

func findTotalPower(power []int32) int32 {
	fmt.Println("len power", len(power))
	// fmt.Println("power", power)
	lenPower := len(power)

	resChan := make(chan uint64, lenPower)
	// mem := make(map[string]uint64, 0)
	// var overallSum uint64
	wg := sync.WaitGroup{}
	for i := 0; i < lenPower; i++ {
		for j := 0; j < lenPower; j++ {
			if i <= j {
				// key := fmt.Sprint(power[i : j+1])
				// // fmt.Println("key", key)
				// if mm, ok := mem[key]; ok {
				//     overallSum += mm
				// } else {
				// }
				wg.Add(1)
				go getMinSum(&wg, power[i:j+1], resChan)
				// fmt.Println("i:j", i, j, ms, power[i:j])
			}
		}
	}

	ovsCh := make(chan uint64, 1)
	go newFunction(ovsCh, resChan)
	wg.Wait()
	close(resChan)

	overallSum := <-ovsCh
	osm := overallSum % 1000000007
	overallSum32 := int32(osm)
	fmt.Println("overallSum32", overallSum, osm, overallSum32)

	return overallSum32
}

func newFunction(ovsCh, resChan chan uint64) {
	var overallSum uint64
	for v := range resChan {
		overallSum += v
	}

	ovsCh <- overallSum
}

func getMinSum(wg *sync.WaitGroup, powerChunk []int32, resChan chan uint64) {
	defer wg.Done()
	min := uint64(2147483647)

	var localSum uint64
	for i := 0; i < len(powerChunk); i++ {
		cur := uint64(powerChunk[i])
		localSum += cur
		if cur < min {
			min = cur
		}
	}

	// fmt.Println("min", min)
	// fmt.Println("localSum", localSum)
	// fmt.Println("msum", msum)
	// fmt.Println("msum % 1000000007", msum%1000000007)

	resChan <- min * localSum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	powerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var power []int32

	for i := 0; i < int(powerCount); i++ {
		powerItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		powerItem := int32(powerItemTemp)
		power = append(power, powerItem)
	}

	result := findTotalPower(power)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
		panic(err)
	}
}
