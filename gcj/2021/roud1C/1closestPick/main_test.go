package main

import "testing"

func Test_check(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("check() = %v, want %v", got, tt.want)
				t.FailNow()
			}
		})
	}
}
