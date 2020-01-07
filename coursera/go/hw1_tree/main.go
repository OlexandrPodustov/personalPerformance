package main

import (
	"fmt"
	"io"
	"log"
	"os"
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
		//panic(err.Error())
		log.Println("aaa", err.Error())
	}
}

func dirTree(writer io.Writer, s string, b bool) error {
	file, err := os.Open(s)
	checkErr(err)
	fileDirs, err := file.Readdir(0)
	checkErr(err)

	err = file.Close()
	checkErr(err)

	for _, v := range fileDirs {
		if v.IsDir() {
			_, err = fmt.Fprintf(writer, "|")
			checkErr(err)
			_, err = fmt.Fprintf(writer, "├─── %+v\n", v.Name())
			checkErr(err)
			dirTree(writer, v.Name(), b)
			os.Chdir("..")
		}
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
