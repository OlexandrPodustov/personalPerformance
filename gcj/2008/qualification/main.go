package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "input.txt"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln("open file "+filename+"failed, err:", err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
