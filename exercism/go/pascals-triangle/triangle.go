package pascal

const testVersion = 1

func Triangle(in int) [][]int {
	res := make([][]int, in)
	res[0] = []int{1}

	for i := 1; i < in; i++ {
		res[i] = make([]int, i+1)
		res[i] = fillSlice(res[i-1])
	}

	return res
}

func fillSlice(in []int) []int {
	newSlice := make([]int, len(in)+1)
	copy(newSlice, in)

	for i := 1; i < len(in); i++ {
		newSlice[i] = in[i-1] + in[i]
	}

	newSlice[len(in)] = 1
	return newSlice
}
