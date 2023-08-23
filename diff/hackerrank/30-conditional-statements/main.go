package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	const (
		weird    = "Weird"
		notWeird = "Not Weird"
	)

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	result := ""
	if even := nTemp%2 == 0; !even {
		result = weird
	} else if even && nTemp >= 2 && nTemp <= 5 {
		result = notWeird
	} else if even && nTemp >= 6 && nTemp <= 20 {
		result = weird
	} else if even && nTemp >= 20 {
		result = notWeird
	}

	fmt.Println(result)
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
