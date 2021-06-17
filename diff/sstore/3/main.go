package main

import "fmt"

func main() {
	// input := []int{3, 4, 3, -3, 7, 4, -6}
	// input := []int{1, -2, 2, 2, 2, 2, 4, -4}
	// input := []int{0, 1, 2, -3}
	// input := []int{1, 2, 3, 4}
	input := []int{1, 1, -2, 2}
	// input := []int{0, 0, 0, 0}
	res := check(input)
	fmt.Println(res)
}

func check(input []int) bool {
	if len(input) < 4 {
		return false
	}

	type indexVal struct {
		oi int
		ii int
	}

	m := make(map[int]indexVal, len(input))
	for oi, ov := range input {
		for ii, iv := range input {
			// if ii <= oi {
			// 	continue
			// }

			val := indexVal{
				oi: oi,
				ii: ii,
			}
			key := ov + iv

			m[key] = val
		}
	}

	for oi, ov := range input {
		for ii, iv := range input {
			if ii <= oi {
				continue
			}

			vv := -1 * (ov + iv)
			mv, ok := m[vv]
			_ = mv

			// if vv == 0 && mv == 2 {
			// 	continue
			// }
			// if mv == 1 {
			// 	// vv = -2
			// 	// ov = -2
			// 	// iv =  4
			// 	if vv == ov || vv == iv {
			// 		continue
			// 	}
			// }

			if ok {
				return true
			}
		}
	}

	return false
}
