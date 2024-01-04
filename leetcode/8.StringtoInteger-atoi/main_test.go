package main

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "42",
			},
			want: 42,
		},
		{
			name: "2",
			args: args{
				s: "   -42",
			},
			want: -42,
		},
		{
			name: "3",
			args: args{
				s: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "4",
			args: args{
				s: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "5",
			args: args{
				s: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "6",
			args: args{
				s: "+-12",
			},
			want: 0,
		},
		{
			name: "7",
			args: args{
				s: "00000-42a1234",
			},
			want: 0,
		},
		{
			name: "8",
			args: args{
				s: "9223372036854775808",
			},
			want: 2147483647,
		},
		{
			name: "9",
			args: args{
				s: "-2147483647",
			},
			want: -2147483647,
		},
		{
			name: "10",
			args: args{
				s: "-5-",
			},
			want: -5,
		},
		{
			name: "11",
			args: args{
				s: "  +  413",
			},
			want: 0,
		},
		{
			name: "12",
			args: args{
				s: "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459",
			},
			want: 2147483647,
		},
		{
			name: "13",
			args: args{
				s: "  0000000000012345678",
			},
			want: 12345678,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
