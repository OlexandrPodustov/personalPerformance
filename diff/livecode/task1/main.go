package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	currentDirectoryName, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	maxFilesForDir := 2
	if err := walk(currentDirectoryName, 0, maxFilesForDir); err != nil {
		log.Fatal(err)
	}
}

var (
	regular   = "├──"
	lastEntry = "└──"
)

func walk(path string, level, limit int) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	var counter int
	for i := 0; i < len(files); i++ {
		v := files[i]
		symbol := regular
		if i == len(files)-1 {
			symbol = lastEntry
		}

		fmt.Println(strings.Repeat("\t", level), symbol, v.Name())
		if v.IsDir() {
			if err := walk(path+"/"+v.Name(), level+1, limit); err != nil {
				fmt.Println(err)
				// return err
			}
		}
		counter++
		if counter > limit {
			return nil
		}
	}

	return nil
}
