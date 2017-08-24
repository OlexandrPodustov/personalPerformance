// Package twofer should have a package comment that summarizes what it's about.
package twofer

// ShareWith needs a comment documenting it.
func ShareWith(in string) string {
	if len(in) > 0 {
		st := "One for " + in + ", one for me."
		return st
	}
	return "One for you, one for me."
}
