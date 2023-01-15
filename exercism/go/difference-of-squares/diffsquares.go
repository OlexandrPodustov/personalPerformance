package diffsquares

const testVersion = 1

func SquareOfSums(in int) int {
	var sum int
	for i := 0; i <= in; i++ {
		sum += i
	}
	res := sum * sum
	return res
}

func SumOfSquares(in int) int {
	var sum int
	for i := 0; i <= in; i++ {
		sum += i * i
	}

	return sum
}

func Difference(in ...int) int {
	var res int
	for _, v := range in {
		res = SquareOfSums(v) - SumOfSquares(v)
	}
	return res
}
