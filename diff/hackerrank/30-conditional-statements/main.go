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
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)

	const (
		weird    = "Weird"
		notWeird = "Not Weird"
	)

	result := ""
	even := N%2 == 0
	odd := !even
	if odd {
		result = weird
	} else if even && N >= 2 && N <= 5 {
		result = notWeird
	} else if even && N >= 6 && N <= 20 {
		result = weird
	} else if even && N >= 20 {
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
		panic(err)
	}
}
