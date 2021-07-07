// To execute Go code, please declare a func main() in a package "main"
package main

import "fmt"

type interval struct {
	start int
	end   int
}

func main() {
	calendar := []interval{
		{10, 12},
		{4, 8},
		{4, 88},
		{9, 10},
		{3, 5},
		{0, 1},
	}
	fmt.Println(calendar)
	res := mergeRanges(calendar)

	for {
		lenOld := len(res)
		res = mergeRanges(res)
		lenNew := len(res)

		if lenOld == lenNew {
			break
		}
	}

	sortRanges(res)
	fmt.Println(res)
}

func mergeRanges(a []interval) []interval {
	if len(a) == 0 {
		return nil
	}
	result := []interval{a[0]}

	for i := 1; i < len(a); i++ {
		v := a[i]

		var intersectOnce bool
		for j, rv := range result {
			// fmt.Println("result len", len(result), result)
			// fmt.Println("result j", j)

			if intersect(rv, v) {
				intersectOnce = true
				result[j] = merge(rv, v)
			}
		}
		if !intersectOnce {
			result = append(result, v)
		}
		// fmt.Println()
	}

	return result
}

func sortRanges(a []interval) {
	if len(a) == 0 {
		return
	}
	for {
		var moved bool
		for i := 0; i < len(a)-1; i++ {
			if a[i].start > a[i+1].start {
				moved = true
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
		if !moved {
			break
		}
	}
}

func merge(l, r interval) interval {
	res := l
	if res.start > r.start {
		res.start = r.start
	}
	if res.end < r.end {
		res.end = r.end
	}

	return res
}

func intersect(l, r interval) bool {
	// fmt.Println("	intersect l, r", l, r)

	if l.start > r.start {
		l, r = r, l
	}
	if l.end >= r.start {
		// fmt.Println("	intersect true")
		return true
	}
	// fmt.Println("	intersect false")

	return false
}
