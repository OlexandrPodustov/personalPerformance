package main

import (
	"fmt"
	"strings"
)

func main() {
	var result strings.Builder
	for {
		var input []byte
		if _, err := fmt.Scan(&input); err != nil {
			// fmt.Println("err", err)
			break
		}
		result.Write(input)
		result.Write([]byte(" "))
	}

	fmt.Println("Hello, World.")
	fmt.Println(result.String())
}
