package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "abcabcbb",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "bbbbb",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "f",
			args: args{
				s: "f",
			},
			want: 1,
		},
		{
			name: "pwwkew",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "au",
			args: args{
				s: "au",
			},
			want: 2,
		},
		{
			name: "aab",
			args: args{
				s: "aab",
			},
			want: 2,
		},
		{
			name: "abcb",
			args: args{
				s: "abcb",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
