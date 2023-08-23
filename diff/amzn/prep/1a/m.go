package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type inv struct {
	start  int32
	amount int32
}

func numberOfItems(s string, startIndices []int32, endIndices []int32) []int32 {
	result := make([]inv, 0)
	var start int32 = 1
	var amount int32
	var firstOpen bool
	fmt.Println("s", s)

	for i, v := range s {
		fmt.Println("i, v", i, v, amount)

		if !firstOpen {
			if v != 124 {
				continue
			}
			firstOpen = true
			start = int32(i) + 1
			fmt.Println("firstOpen", start, firstOpen, i)
			continue
		}

		if v == 42 {
			amount++
		}

		if v == 124 { // closed compartment. save result
			result = append(result, inv{start: start, amount: amount})

			start = int32(i) + 1
			amount = 0
		}
	}
	fmt.Println(result)

	grandRes := make([]int32, 0)
	for i := 0; i < len(startIndices); i++ {
		startInd := startIndices[i]
		endInd := endIndices[i]
		fmt.Println("startInd", startInd)
		fmt.Println("endInd", endInd)

		for _, v := range result {
			if endInd >= v.start+v.amount {
				break
			}
			if startInd >= v.start && endInd >= v.start+v.amount {
				grandRes = append(grandRes, v.amount)
			}
		}
	}

	return grandRes
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer checkError(stdout.Close())

	s := readLine(reader)

	startIndicesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 0)
	checkError(err)

	var startIndices []int32
	for i := 0; i < int(startIndicesCount); i++ {
		startIndicesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 32)
		checkError(err)
		startIndicesItem := int32(startIndicesItemTemp)
		startIndices = append(startIndices, startIndicesItem)
	}

	endIndicesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 0)
	checkError(err)

	var endIndices []int32
	for i := 0; i < int(endIndicesCount); i++ {
		endIndicesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 32)
		checkError(err)
		endIndicesItem := int32(endIndicesItemTemp)
		endIndices = append(endIndices, endIndicesItem)
	}

	result := numberOfItems(s, startIndices, endIndices)
	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}
	fmt.Fprintf(writer, "\n")

	checkError(writer.Flush())
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
