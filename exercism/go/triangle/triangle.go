package triangle

const testVersion = 3

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

// Pick values for the following identifiers used by the test program.
var (
	naT Kind = "NaT" // not a triangle
	equ Kind = "Equ" // equilateral
	iso Kind = "Iso" // isosceles
	sca Kind = "Sca" // scalene
)

func KindFromSides(a, b, c float64) Kind {
	if a*0 != 0 || b*0 != 0 || c*0 != 0 {
		return naT
	} else if a+b+c <= 0 {
		return naT
	} else if a+b < c || a+c < b || b+c < a {
		return naT
	} else if a == b && a == c {
		return equ
	} else if a == b || a == c || b == c {
		return iso
	} else if a != b && a != c && b != c {
		return sca
	}

	return "not defined yet"
}
