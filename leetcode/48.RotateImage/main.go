package main

import "fmt"

func rotate(matrix [][]int) {

	for i := 0; i < len(matrix); i++ {
		start := 0
		end := len(matrix[i]) - 1
		for start < end {
			matrix[i][start], matrix[i][end] = matrix[i][end], matrix[i][start]
			start++
			end--
		}
	}
	fmt.Println("matrix", matrix)

	j := len(matrix) - 1
	for i := 0; i < len(matrix); i++ {
		fmt.Println("swap", matrix[i][j-1], matrix[i+1][j])
		j--

		matrix[i][j-1], matrix[i+1][j] = matrix[i+1][j], matrix[i][j-1]
	}

	return
}

// matrix [
// [11 	9 1 5]
// [10 	8 4 2]
// [7 	6 3 13]
// [16 12 14 15]]

// want[
// [15 13 2 5]
// [14 3 4 1]
// [12 6 8 9]
// [16 7 10 11]]

// after first turn
// got: [
// [3 2 1]
// [6 5 4]
// [9 8 7]
// ]

// want: [
// [7,4,1]
// [8,5,2]
// [9,6,3]
// ]
