package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	sumOp = "SUM"
	difOp = "DIF"
	intOp = "INT"
)

func main() {
	var input string
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	_, err := fmt.Scan(&input)
	if err != nil {
		log.Println("parse testCasesAmount:", err)
		return
	}

	log.Println("input", input)

	reg := regexp.MustCompile(`[\[]]`)
	sliceSubmatches := reg.FindAllStringSubmatch(input, -1)
	for i, v := range sliceSubmatches {
		fmt.Println("sliceSubmatche", i, v)
	}

	// res, err := calc(input)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, v := range res {
	// 	fmt.Println(v)
	// }
}

func calc(input string) ([]int, error) {
	log.Println("input", input)

	reg := regexp.MustCompile(`\[(.*)\]`)
	sliceSubmatches := reg.FindAllStringSubmatch(input, -1)

	operand := sliceSubmatches[0][1]
	operand = strings.TrimSpace(operand)
	fmt.Printf("result %q \n", operand)

	switch {
	case strings.HasPrefix(operand, sumOp):
		log.Println("find ", sumOp, input)

		newInput := strings.TrimPrefix(operand, sumOp)
		newInput = strings.TrimSpace(newInput)

		return sum(newInput)

	case strings.HasPrefix(operand, difOp):
		log.Println("find ", difOp, input)

		newInput := strings.TrimPrefix(operand, difOp)
		newInput = strings.TrimSpace(newInput)

		return diff(newInput)

	case strings.HasPrefix(operand, intOp):
		log.Println("find ", intOp, input)

		newInput := strings.TrimPrefix(operand, intOp)
		newInput = strings.TrimSpace(newInput)

		return intersect(newInput)
	}

	return nil, nil
}

func calcOld(input string) ([]int, error) {
	// parse highest level input
	// call calc again untill this is the lowest level operation

	if strings.ContainsAny(input, "[]") {
		log.Println("perform recursive call", input)

		result, err := calc(input)
		if err != nil {
			return nil, err
		}
		log.Println("recursive call result", result)
	}

	log.Println("perform simple operation", input)

	// var operation, args string
	// switch operation {
	// case "SUM":
	// // log.Printf("find sum", input)
	// // sum(args)
	// case "DIF":
	// 	// diff(args)
	// case "INT":
	// 	// intersect(args)
	// }

	// result, err := sum(input)
	// if err != nil {
	// 	return nil, err
	// }

	// return result, nil
	return nil, nil
}

func sum(input string) ([]int, error) {
	args := strings.Split(input, " ")
	parts, err := read(args...)
	if err != nil {
		return nil, fmt.Errorf("input: %q, err: %w", input, err)
	}
	if len(parts) == 0 {
		return nil, fmt.Errorf("nothing to compare in: %q, err: %w", input, err)
	}
	if len(parts) == 1 {
		return parts[0], nil
	}

	distinct := make(map[int]struct{})
	for _, part := range parts {
		for _, v := range part {
			distinct[v] = struct{}{}
		}
	}

	var result []int
	for key := range distinct {
		result = append(result, key)
	}

	sort.Ints(result)

	return result, nil
}

func diff(input string) ([]int, error) {
	args := strings.Split(input, " ")
	parts, err := read(args...)
	if err != nil {
		return nil, fmt.Errorf("input: %q, err: %w", input, err)
	}
	if len(parts) == 0 {
		return nil, fmt.Errorf("nothing to compare in: %q, err: %w", input, err)
	}
	if len(parts) == 1 {
		return parts[0], nil
	}

	etalon := make(map[int]struct{})
	for _, v := range parts[0] {
		etalon[v] = struct{}{}
	}

	for _, p := range parts[1:] {
		for _, v := range p {
			delete(etalon, v)
		}
	}

	var result []int
	for key := range etalon {
		result = append(result, key)
	}

	sort.Ints(result)

	return result, nil
}

func intersect(input string) ([]int, error) {
	args := strings.Split(input, " ")
	parts, err := read(args...)
	if err != nil {
		return nil, fmt.Errorf("input: %q, err: %w", input, err)
	}
	if len(parts) == 0 {
		return nil, fmt.Errorf("nothing to compare in: %q, err: %w", input, err)
	}
	if len(parts) == 1 {
		return nil, fmt.Errorf("only one part for comparison isn't enough: %q", input)
	}

	left := make(map[int]struct{})
	for _, v := range parts[0] {
		left[v] = struct{}{}
	}

	right := make(map[int]struct{})
	for _, part := range parts[1:] {
		for _, v := range part {
			right[v] = struct{}{}
		}
	}

	var result []int
	for key := range right {
		if _, ok := left[key]; ok {
			result = append(result, key)
		}
	}

	sort.Ints(result)

	return result, nil
}

func read(files ...string) ([][]int, error) {
	var result [][]int

	for _, v := range files {
		content, err := ioutil.ReadFile(v)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %q, err: %w", v, err)
		}

		// log.Println("[]byte content", content)
		sbb := bytes.ReplaceAll(content, []byte("\r"), []byte(""))
		sbb = bytes.ReplaceAll(sbb, []byte("\n"), []byte(""))
		// log.Printf("ints %T %+v \n", sbb, sbb)

		var res []int
		for _, v := range sbb {
			sv := string(v)
			if sv == "\n" || sv == "\r" {
				// log.Println("AAAAAAA", sv)
				continue
			}

			i, err := strconv.Atoi(sv)
			if err != nil {
				log.Printf("input: %T %+v. err: %T %+v \n", sv, sv, err, err)
				return nil, err
			}
			// log.Println("final int", i)

			res = append(res, i)
		}

		log.Printf("res %T %+v \n", res, res)

		result = append(result, res)
	}

	// log.Printf("result %T %+v \n", result, result)

	return result, nil
}
