package triangle

import (
	"math"
	"testing"
)

const targetTestVersion = 3

type testCase struct {
	want    Kind
	a, b, c float64
}

// basic test cases
var testData = []testCase{
	{equ, 2, 2, 2},    // same length
	{equ, 10, 10, 10}, // a little bigger
	{iso, 3, 4, 4},    // last two sides equal
	{iso, 4, 3, 4},    // first and last sides equal
	{iso, 4, 4, 3},    // first two sides equal
	{iso, 10, 10, 2},  // again
	{iso, 2, 4, 2},    // a "triangle" that is just a line is still OK
	{sca, 3, 4, 5},    // no sides equal
	{sca, 10, 11, 12}, // again
	{sca, 5, 4, 2},    // descending order
	{sca, .4, .6, .3}, // small sides
	{sca, 1, 4, 3},    // a "triangle" that is just a line is still OK
	{sca, 5, 4, 6},    // 2a == b+c looks like equilateral, but isn't always.
	{sca, 6, 4, 5},    // 2a == b+c looks like equilateral, but isn't always.
	{naT, 0, 0, 0},    // zero length
	{naT, 3, 4, -5},   // negative length
	{naT, 1, 1, 3},    // fails triangle inequality
	{naT, 2, 5, 2},    // another
	{naT, 7, 3, 2},    // another
}

// generate cases with NaN and Infs, append to basic cases
func init() {
	nan := math.NaN()
	pinf := math.Inf(1)
	ninf := math.Inf(-1)
	nf := make([]testCase, 4*4*4)
	i := 0
	for _, a := range []float64{3, nan, pinf, ninf} {
		for _, b := range []float64{4, nan, pinf, ninf} {
			for _, c := range []float64{5, nan, pinf, ninf} {
				nf[i] = testCase{naT, a, b, c}
				i++
			}
		}
	}
	testData = append(testData, nf[1:]...)
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

// Test that the kinds are not equal to each other.
// If they are equal, then TestKind will return false positives.
func TestKindsNotEqual(t *testing.T) {
	kindsAndNames := []struct {
		kind Kind
		name string
	}{
		{equ, "Equ"},
		{iso, "Iso"},
		{sca, "Sca"},
		{naT, "NaT"},
	}

	for i, pair1 := range kindsAndNames {
		for j := i + 1; j < len(kindsAndNames); j++ {
			pair2 := kindsAndNames[j]
			if pair1.kind == pair2.kind {
				t.Fatalf("%s should not be equal to %s", pair1.name, pair2.name)
			}
		}
	}
}

func TestKind(t *testing.T) {
	for _, test := range testData {
		got := KindFromSides(test.a, test.b, test.c)
		if got != test.want {
			t.Fatalf("Triangle with sides, %g, %g, %g = %v, want %v",
				test.a, test.b, test.c, got, test.want)
		}
	}
}

func BenchmarkKind(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testData {
			KindFromSides(test.a, test.b, test.c)
		}
	}
}
