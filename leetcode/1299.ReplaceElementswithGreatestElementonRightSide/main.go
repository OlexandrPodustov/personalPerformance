package main

func replaceElements(arr []int) []int {
	if len(arr) == 1 {
		arr[0] = -1
		return arr
	}

	maxEl := -1
	for i := len(arr) - 1; i >= 0; i-- {
		prevMax := maxEl

		if arr[i] > maxEl {
			maxEl = arr[i]
		}

		arr[i] = prevMax
	}

	return arr
}
