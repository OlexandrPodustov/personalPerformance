package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "abccccdd",
			args: args{
				s: "abccccdd",
			},
			want: 7,
		},
		{
			name: "a",
			args: args{
				s: "a",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
