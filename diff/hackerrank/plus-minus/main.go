package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func plusMinus(arr []int) {
	divider := float64(len(arr))

	var positive, negative, zero float64
	for _, v := range arr {
		switch {
		case v > 0:
			positive++
		case v < 0:
			negative++
		case v == 0:
			zero++
		}
	}

	fmt.Printf("%.6f\n", positive/divider)
	fmt.Printf("%.6f\n", negative/divider)
	fmt.Printf("%.6f\n", zero/divider)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	nTemp, err := strconv.Atoi(strings.TrimSpace(readLine(reader)))
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int
	for i := 0; i < nTemp; i++ {
		arrItem, err := strconv.Atoi(arrTemp[i])
		checkError(err)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
