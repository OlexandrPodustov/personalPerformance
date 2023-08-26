package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func getMinimumDays(parcels []int32) int32 {
	differentAmounts := make(map[int32]struct{}, 0)
	fmt.Println(len(parcels))
	fmt.Println("parcels", parcels)

	for _, v := range parcels {
		if v == 0 {
			continue
		}
		differentAmounts[v] = struct{}{}
	}
	fmt.Println(len(differentAmounts))

	return int32(len(differentAmounts))
}

func main() {
	oneKibibyte := 1024
	reader := bufio.NewReaderSize(os.Stdin, 16*oneKibibyte*oneKibibyte)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer checkError(stdout.Close())

	parcelsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 0)
	checkError(err)

	var parcels []int32
	for i := 0; i < int(parcelsCount); i++ {
		parcelsItem, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 32)
		checkError(err)
		parcels = append(parcels, int32(parcelsItem))
	}

	writer := bufio.NewWriterSize(stdout, 16*oneKibibyte*oneKibibyte)
	fmt.Fprintf(writer, "%d\n", getMinimumDays(parcels))
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
