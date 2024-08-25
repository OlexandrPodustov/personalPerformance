package main

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := range board[i] {
			if possible(board, i, j, word) {
				return true
			}
		}
	}

	return false
}

type cell struct {
	i, j int
}

func possible(board [][]byte, i, j int, word string) bool {
	visited := make(map[cell]struct{})
	return checkAdjacentForLetter(board, i, j, word, visited, 0)
}

func checkAdjacentForLetter(board [][]byte, i, j int, word string, visited map[cell]struct{}, wi int) bool {
	if _, ok := visited[cell{i, j}]; ok {
		return false
	}
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
		return false
	}

	if board[i][j] != word[wi] {
		return false
	}
	if wi >= len(word)-1 {
		return true
	}
	visited[cell{i, j}] = struct{}{}

	someTrue := checkAdjacentForLetter(board, i-1, j, word, visited, wi+1) ||
		checkAdjacentForLetter(board, i, j-1, word, visited, wi+1) ||
		checkAdjacentForLetter(board, i+1, j, word, visited, wi+1) ||
		checkAdjacentForLetter(board, i, j+1, word, visited, wi+1)

	if !someTrue {
		delete(visited, cell{i, j})
	}

	return someTrue
}
