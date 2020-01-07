package pythagorean

const testVersion = 1

type Triplet [3]int

func Range(min, max int) (res []Triplet) {

	for i := 1; i <= max; i++ {
		for ii := i + 1; ii <= max; ii++ {
			a, b, c := check(ii, i)
			if a > min && c <= max {
				res = append(res, Triplet{a, b, c})
			}
		}
	}
	return res
}

func Sum(in int) (resultTriplet []Triplet) {
	var tripletMap = make(map[[3]int]bool)
	for i := 1; i <= 10; i++ {
		for ii := i + 1; ii <= 10; ii++ {
			a, b, c := check(ii, i)
			if a+b+c == in {
				tripletMap[Triplet{a, b, c}] = true
			}
			for iii := 1; iii < 100; iii++ {
				if (a+b+c)*iii == in {
					tripletMap[Triplet{a * iii, b * iii, c * iii}] = true
				}
			}
		}
	}

	for key := range tripletMap {
		resultTriplet = append(resultTriplet, key)
	}

	for {
		var cnt int
		for i := 0; i < len(resultTriplet)-1; i++ {
			cnt = 0
			if resultTriplet[i][0] > resultTriplet[i+1][0] {
				resultTriplet[i+1], resultTriplet[i] = resultTriplet[i], resultTriplet[i+1]
				cnt++
			}

		}
		if cnt == 0 {
			break
		}
	}

	return resultTriplet
}

func check(m, n int) (a, b, c int) {
	a = m*m - n*n
	b = 2 * m * n
	c = m*m + n*n

	if a > b {
		a, b = b, a
	}

	return
}
