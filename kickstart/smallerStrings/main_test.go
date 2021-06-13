package main

import "testing"

func Test_check(t *testing.T) {
	tests := []struct {
		name string
	}{
		// 		3
		// 2 3
		// bc
		// 5 5
		// abcdd
		// 1 5
		// d

	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			check()
		})
	}
}
