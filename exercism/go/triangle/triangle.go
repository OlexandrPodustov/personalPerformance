package triangle

const testVersion = 3

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

// Pick values for the following identifiers used by the test program.
var NaT Kind = "NaT" // not a triangle
var Equ Kind = "Equ" // equilateral
var Iso Kind = "Iso" // isosceles
var Sca Kind = "Sca" // scalene

func KindFromSides(a, b, c float64) Kind {
	if a*0 != 0 || b*0 != 0 || c*0 != 0 {
		return NaT
	} else if a+b+c <= 0 {
		return NaT
	} else if a+b < c || a+c < b || b+c < a {
		return NaT
	} else if a == b && a == c {
		return Equ
	} else if a == b || a == c || b == c {
		return Iso
	} else if a != b && a != c && b != c {
		return Sca
	}

	return "not defined yet"
}
