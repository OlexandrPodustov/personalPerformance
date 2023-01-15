package accumulate

const testVersion = 1

func Accumulate(s []string, f func(string) string) []string {
	slc := make([]string, len(s))

	for i, r := range s {
		slc[i] = f(r)
	}
	return slc
}
