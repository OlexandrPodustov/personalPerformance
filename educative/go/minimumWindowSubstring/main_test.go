package main

import (
	"testing"
)

func Test_minWindow(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ABCD",
			args: args{
				str1: "ABCD",
				str2: "ABC",
			},
			want: "ABC",
		},

		{
			name: "XYZYX",
			args: args{
				str1: "XYZYUX",
				str2: "XYZ",
			},
			want: "XYZ",
		},

		{
			name: "ABXYZJKLSNFC",
			args: args{
				str1: "ABXYZJKLSNFC",
				str2: "ABC",
			},
			want: "ABXYZJKLSNFC",
		},

		{
			name: "AAAAAAAAAAA",
			args: args{
				str1: "AAAAAAAAAAA",
				str2: "A",
			},
			want: "A",
		},

		{
			name: "ABDFGDCKAB",
			args: args{
				str1: "ABDFGDCKAB",
				str2: "ABCD",
			},
			want: "DCKAB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
