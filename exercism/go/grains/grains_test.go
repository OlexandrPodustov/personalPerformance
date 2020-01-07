package grains

import (
	"testing"
)

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func TestSquare(t *testing.T) {
	for _, test := range squareTests {
		actualVal, actualErr := Square(test.input)

		// check actualVal only if no error expected
		if !test.expectError && actualVal != test.expectedVal {
			t.Errorf("Square(%d) expected %d, Actual %d", test.input, test.expectedVal, actualVal)
		}

		// if we expect an error and there isn't one
		if test.expectError && actualErr == nil {
			t.Errorf("Square(%d) expected an error, but error is nil", test.input)
		}
		// if we don't expect an error and there is one
		if !test.expectError && actualErr != nil {
			var _ error = actualErr
			t.Errorf("Square(%d) expected no error, but error is: %s", test.input, actualErr)
		}
	}
}

func TestTotal(t *testing.T) {
	var expected uint64 = 18446744073709551615
	if actual := Total(); actual != expected {
		t.Errorf("Total() expected %d, Actual %d", expected, actual)
	}
}

func BenchmarkSquare(b *testing.B) {
	b.StopTimer()

	for _, test := range squareTests {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			Square(test.input)
		}

		b.StopTimer()
	}
}

func BenchmarkTotal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Total()
	}
}
