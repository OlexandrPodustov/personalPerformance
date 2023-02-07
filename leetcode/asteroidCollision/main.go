package main

func asteroidCollision(asteroids []int) []int {
	j := len(asteroids) - 1
	i := j - 1

	for i >= 0 && j < len(asteroids) {
		ai := abs(asteroids[i])
		aj := abs(asteroids[j])

		// fmt.Println(asteroids, "i, j, len(asteroids)", i, j, "values:", asteroids[i], asteroids[j], ". len", len(asteroids), "- ai, aj:", ai, aj)

		sumAbs := ai + aj
		if abs(asteroids[i]+asteroids[j]) == sumAbs {
			// fmt.Println(asteroids, "will not reach out each other. i, j:", i, j, "abs()=", abs(asteroids[i]+asteroids[j]), "sumAbs", sumAbs, "=", ai, "+", aj)

			i--
			j--
			continue
		}
		if asteroids[i] < 0 && asteroids[j] > 0 {
			// fmt.Println(asteroids, "looks into opposite direction. i, j:", i, j, "asteroids[i]<0:", asteroids[i], "asteroids[j]>0:", asteroids[j])
			i--
			j--
			continue
		}
		switch {
		case ai == aj:
			// fmt.Println("case 1: ", i, j, len(asteroids), ai, aj)

			asteroids = append(asteroids[:i], asteroids[j+1:]...)
			// j = len(asteroids) - 1
			// i = j - 1
			i--
			j--
		case ai > aj:
			// fmt.Println("case 2: ", i, j, len(asteroids), ai, aj)

			asteroids = append(asteroids[:j], asteroids[j+1:]...)
			j = len(asteroids) - 1
			i = j - 1

		case ai < aj:
			// fmt.Println("case 3: ", i, j, len(asteroids), ai, aj)

			// fmt.Println("asteroids 1", asteroids)
			asteroids = append(asteroids[:i], asteroids[j:]...)
			// fmt.Println("asteroids 2", asteroids)
			j = len(asteroids) - 1
			i = j - 1
		}

		// fmt.Println(asteroids, "end loop conditions... i, j, len(asteroids)", i, j, len(asteroids))
	}

	// fmt.Println()
	return asteroids
}

func abs(param int) int {
	if param < 0 {
		return -1 * param
	}
	return param
}
