package strain

const testVersion = 1

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(in func(int) bool) (res Ints) {
	for _, value := range i {
		if in(value) {
			res = append(res, value)
		}
	}
	return res
}

func (i Ints) Discard(in func(int) bool) (res Ints) {
	//return i.Keep(func(x int) bool {
	//	return !in(x)
	//})
	for _, value := range i {
		if !in(value) {
			res = append(res, value)
		}
	}
	return res
}

func (i Lists) Keep(in func([]int) bool) (res Lists) {
	for _, value := range i {
		if in(value) {
			res = append(res, value)
		}
	}
	return res
}

func (i Strings) Keep(in func(string) bool) (res Strings) {
	for _, value := range i {
		if in(value) {
			res = append(res, value)
		}
	}
	return res
}
