package main

func duplicateZeros(arr []int) {
	la := len(arr)
	if la < 2 {
		return
	}

	for i := 0; i < la; i++ {
		if arr[i] == 0 {

			newRemainder := append([]int{0}, arr[i:la-1]...)

			arr = append(arr[:i], newRemainder...)
			i++

		}
	}

	return
}
