package main

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	increaseStagePresent := false
	decreaseStagePresent := false
	for i := 1; i < len(arr); i++ {
		if arr[i-1] == arr[i] {
			return false
		}
		if arr[i-1] < arr[i] {
			increaseStagePresent = true
			if decreaseStagePresent {
				return false
			}
			continue
		}
		decreaseStagePresent = true
	}

	if !decreaseStagePresent || !increaseStagePresent {
		return false
	}

	return true
}
