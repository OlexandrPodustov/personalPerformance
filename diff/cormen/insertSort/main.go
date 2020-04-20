package main

import "fmt"

func main() {
	in := []int{31, 41, 59, 26, 41, 58}
	res := insertSortASC(in)
	fmt.Println(res)

	res = insertSortDESC(in)
	fmt.Println(res)
}

func insertSortASC(in []int) []int {
	for i := 0; i < len(in)-1; i++ {
		for i >= 0 && in[i+1] < in[i] {
			in[i+1], in[i] = in[i], in[i+1]
			i--
		}
	}

	return in
}

func insertSortDESC(in []int) []int {
	for i := 0; i < len(in)-1; i++ {
		for i >= 0 && in[i+1] > in[i] {
			in[i+1], in[i] = in[i], in[i+1]
			i--
		}
	}

	return in
}
