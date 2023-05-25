package summultiples

const testVersion = 2

func SumMultiples(in int, list ...int) (res int) {
	ma := make(map[int]bool)
	for iter := 1; iter < in; iter++ {
		for _, digit := range list {
			number := in - iter
			if number%digit == 0 {
				ma[number] = true
			}
		}
	}
	for ind := range ma {
		res += ind
	}

	return res
}
