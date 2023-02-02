package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		fmt.Println("usage go run main.go . [-f]")
		return
	}

	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func dirTree(out io.Writer, pathStr string, printFiles bool) error {
	return walk(0, out, pathStr, printFiles)
}

func walk(depth int, out io.Writer, pathStr string, printFiles bool) error {
	depth++

	f, err := os.Open(pathStr)
	if err != nil {
		return err
	}
	// fmt.Println("depth name", depth, f.Name())
	// defer fmt.Println()

	files, err := f.Readdir(-1)
	if err != nil {
		return err
	}

	for _, v := range files {
		if v.IsDir() {
			// fmt.Println(pathStr, "dir name", v.Name())
			fmt.Println("├", strings.Repeat("───", depth), v.Name())

			p := path.Join(pathStr, string(os.PathSeparator), v.Name())
			if err := walk(depth, out, p, printFiles); err != nil {
				return err
			}
		}
	}

	return nil
}
