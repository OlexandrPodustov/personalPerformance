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
