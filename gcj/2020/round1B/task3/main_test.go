package main

import (
	"strings"
	"testing"
)

func TestMultiply(t *testing.T) {
	tcs := getCases()

	for _, tc := range tcs {
		steps, resp := check(tc.x, tc.y)
		if tc.steps != steps {
			t.Errorf("for x: %v, y: %v - wanted: %v, got: %v", tc.x, tc.y, tc.steps, steps)
			t.FailNow()
		}
		if tc.want != strings.Join(resp, ",") {
			t.Errorf("for x: %v, y: %v - wanted: %v, got: %v", tc.x, tc.y, tc.want, resp)
			t.FailNow()
		}

	}
}

type tc struct {
	x, y  int
	steps int
	want  string
}

func getCases() []tc {
	return []tc{
		{x: 2, y: 2, steps: 1, want: "2 1"},
		{x: 3, y: 2, steps: 2, want: "3 2,2 1"},
		// {x: 2, y: 3, steps: 2, want: "2 3,2 2"},
	}
}
