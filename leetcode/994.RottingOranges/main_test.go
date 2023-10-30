package main

import "testing"

func Test_orangesRotting(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[[2,1,1],[1,1,0],[0,1,1]]",
			args: args{
				grid: [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			},
			want: 4,
		},
		{
			name: "[[2,1,1],[0,1,1],[1,0,1]]",
			args: args{
				grid: [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			},
			want: -1,
		},
		{
			name: "[[0,2]]",
			args: args{
				grid: [][]int{{0, 2}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := orangesRotting(tt.args.grid); got != tt.want {
				t.Errorf("orangesRotting(%v) = %v, want %v", tt.args.grid, got, tt.want)
			}
		})
	}
}
