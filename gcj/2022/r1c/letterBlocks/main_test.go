package main

import "testing"

func Test_hasNonConsecutiveRepetitiveLetters(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				word: "ab",
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				word: "baaaabbb",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				word: "aaaacccbbb",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasNonConsecutiveRepetitiveLetters(tt.args.word); got != tt.want {
				t.Errorf("hasNonConsecutiveRepetitiveLetters() = %v, want %v", got, tt.want)
				t.FailNow()
			}
		})
	}
}
