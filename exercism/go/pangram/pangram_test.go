package pangram

import (
	"testing"
)

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func TestPangram(t *testing.T) {
	for _, test := range testCases {
		actual := IsPangram(test.input)
		if actual != test.expected {
			t.Errorf("Pangram test [%s], expected [%t], actual [%t]", test.input, test.expected, actual)
			if !test.expected {
				t.Logf("[%s] should not be a pangram because : %s\n", test.input, test.description)
			}
		}
	}
}

func BenchmarkPangram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			IsPangram(test.input)
		}
	}
}
