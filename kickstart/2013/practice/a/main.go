package main // Bad Horse
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("A-small-practice-1.in")
	if err != nil {
		log.Fatal("open file", err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatal("close file", err)
		}
	}()

	s := bufio.NewScanner(f)
	var testCasesAmount int
	if s.Scan() {
		testCasesAmount, err = strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal("close file", err)
		}
	}

	for i := 1; i <= testCasesAmount; i++ {
		result := scanCase(s)
		fmt.Printf("Case #%v: %v\n", i, result)
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func scanCase(s *bufio.Scanner) string {
	var rows int
	var err error
	if s.Scan() {
		rows, err = strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal("strconv.Atoi", err)
		}
	}

	left := make(map[string]struct{})
	right := make(map[string]struct{})
	for i := 0; i < rows; i++ {
		if !s.Scan() {
			break
		}
		ss := strings.Split(s.Text(), " ")

		left[ss[0]] = struct{}{}
		right[ss[1]] = struct{}{}

		// fmt.Printf("case: %v, %v\n", i, s.Text())
	}

	// fmt.Printf("left: %v\n", left)
	// fmt.Printf("right: %v\n", right)

	res := "Yes"
	for l := range left {
		if _, ok := right[l]; ok {
			res = "No"
			break
		}
	}

	for l := range right {
		if _, ok := left[l]; ok {
			res = "No"
			break
		}
	}

	return res
}
