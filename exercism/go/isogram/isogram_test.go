package isogram

import "testing"

const targetTestVersion = 1

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func TestIsIsogram(t *testing.T) {
	for _, c := range testCases {
		if IsIsogram(c.input) != c.expected {
			t.Fatalf("FAIL: Word %q, expected %v, got %v", c.input, c.expected, !c.expected)
		}

		t.Logf("PASS: Word %q", c.input)
	}
}

func BenchmarkIsIsogram(b *testing.B) {
	b.StopTimer()
	for _, c := range testCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			IsIsogram(c.input)
		}

		b.StopTimer()
	}
}
