package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	//"strings"
	"log"
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
	_, err := fmt.Fprintf(writer, "received path - %+v\t, Base - %+v\n", s, filepath.Base(s))
	checkErr(err)

	file, err := os.Open(s)
	checkErr(err)

	w := new(filepath.WalkFunc)
	err = filepath.Walk(file.Name(), *w)
	checkErr(err)

	//if v.IsDir() {
	//	_, err = fmt.Fprintf(writer, "├─── %+v\n", v.Name())
	//	checkErr(err)
	//}
	err = file.Close()
	checkErr(err)
	//dirTree(writer, file.Name(), b)
	return nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
