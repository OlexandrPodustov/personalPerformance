package main

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < len(matrix); i++ {
		start := 0
		end := len(matrix[i]) - 1
		for start < end {
			matrix[i][start], matrix[i][end] = matrix[i][end], matrix[i][start]
			start++
			end--
		}
	}

	return
}
