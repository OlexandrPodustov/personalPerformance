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
			name: "abcdebdde",
			args: args{
				str1: "abcdebdde",
				str2: "bde",
			},
			want: "bcde",
		},

		// {
		// 	name: "fgrqsqsnodwmxzkzxwqegkndaa",
		// 	args: args{
		// 		str1: "fgrqsqsnodwmxzkzxwqegkndaa",
		// 		str2: "kzed",
		// 	},
		// 	want: "",
		// },

		// {
		// 	name: "michmznaitnjdnjkdsnmichmznait",
		// 	args: args{
		// 		str1: "michmznaitnjdnjkdsnmichmznait",
		// 		str2: "michmznait",
		// 	},
		// 	want: "",
		// },

		{
			name: "afgegrwgwga",
			args: args{
				str1: "afgegrwgwga",
				str2: "aa",
			},
			want: "afgegrwgwga",
		},

		{
			name: "abababa",
			args: args{
				str1: "abababa",
				str2: "ba",
			},
			want: "ba",
		},

		{
			name: "abcdefg",
			args: args{
				str1: "abcdefg",
				str2: "bc",
			},
			want: "bc",
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
