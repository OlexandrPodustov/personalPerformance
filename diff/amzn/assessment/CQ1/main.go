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
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	parcelsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var parcels []int32

	for i := 0; i < int(parcelsCount); i++ {
		parcelsItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		parcelsItem := int32(parcelsItemTemp)
		parcels = append(parcels, parcelsItem)
	}

	result := getMinimumDays(parcels)

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
		fmt.Println(err)
	}
}
