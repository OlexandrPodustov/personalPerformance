package main

func checkIfExist(arr []int) bool {
	doubles := make(map[int]int)
	for i, v := range arr {
		doubles[v*2] = i
	}
	for i, v := range arr {
		if ind, ok := doubles[v]; ok && i != ind {
			return true
		}
	}

	return false
}
