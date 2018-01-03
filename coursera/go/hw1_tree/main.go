package main

import (
	"fmt"
	"io"
	"os"
	//"path/filepath"
	//"strings"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(writer io.Writer, s string, b bool) error {
	file, err := os.Open(s)
	if err != nil {
		return err
	}
	defer file.Close()

	fstat, err := file.Stat()
	if err != nil {
		return err
	}

	fmt.Printf("fstat.IsDir() %#v\n", fstat.IsDir())
	fmt.Printf("fstat.Size() (%+vb)\n", fstat.Size())
	return nil
}
