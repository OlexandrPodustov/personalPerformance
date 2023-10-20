package main

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "4",
			args: args{
				n: 4,
			},
			want: 5,
		},
		{
			name: "5",
			args: args{
				n: 5,
			},
			want: 8,
		},
		{
			name: "6",
			args: args{
				n: 6,
			},
			want: 13,
		},
		{
			name: "7",
			args: args{
				n: 7,
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
