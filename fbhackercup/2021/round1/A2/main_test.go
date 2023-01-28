package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "",
		// 	args: args{
		// 		input: "O",
		// 	},
		// 	want: 0,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		input: "XFO",
		// 	},
		// 	want: 1,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		input: "FFOFF",
		// 	},
		// 	want: 0,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		input: "FXXFXFOOXF",
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		input: "XFOFXFOFXFOFX",
		// 	},
		// 	want: 6,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		input: "O",
		// 	},
		// 	want: 0,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		input: "O",
		// 	},
		// 	want: 0,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		input: "O",
		// 	},
		// 	want: 0,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		input: "O",
		// 	},
		// 	want: 0,
		// },
		// {
		// 	name: "",
		// 	args: args{
		// 		input: "XXXXFXXXOXXX",
		// 	},
		// 	want: 2,
		// },

		{
			name: "",
			args: args{
				input: "FFFFFFFFFF",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.input); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
				t.FailNow()
			}
		})
	}
}
