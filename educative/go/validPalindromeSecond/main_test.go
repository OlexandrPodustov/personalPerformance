package main

import "testing"

func Test_validPalindrome(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				inputString: "RACEACAR",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				inputString: "abbababa",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				inputString: "abcdeca",
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				inputString: "",
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				inputString: "a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.inputString); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
