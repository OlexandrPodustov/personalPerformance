package main

const back byte = '#'

func backspaceCompare(s string, t string) bool {
	s = replaceBackspaces(s)
	t = replaceBackspaces(t)

	return s == t
}

func replaceBackspaces(s string) string {
	i := 0
	for len(s) >= 0 && i < len(s) {
		if s[i] == back {

			s = s[:i-1] + s[i+1:]

			i--
			continue
		}
		i++
	}
	// fmt.Println("new s", s)

	return s
}
