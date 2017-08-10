package pythagorean

const testVersion = 1

type Triplet struct {
	a int
	b int
	c int
}

func Range(a, b int) (res int) {

	return 0
}

func Sum(in int) (res []Triplet) {
	for a := 1; a < 100; a++ {
		for b := 1; b < 100; b++ {
			for c := 1; c < 10; c++ {
				if sum := form(a, b, c); sum == in {
					//res = Triplet{a, b, int(math.Sqrt(float64(c)))}
					res = append([]Triplet{}, Triplet{a * a, b * b, c})
				}
			}
		}

	}

	return res
}

func form(a, b, c int) int {
	sum := a*a + b*b + c*c
	return sum
}
