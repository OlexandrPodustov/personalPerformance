package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		inputString string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				inputString: "abba",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				inputString: "",
			},
			want: true,
		},
		{
			name: "",
			args: args{
				inputString: "a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.inputString); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
				t.FailNow()
			}
		})
	}
}
