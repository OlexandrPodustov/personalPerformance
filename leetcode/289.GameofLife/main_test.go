package main

import (
	"slices"
	"testing"
)

func Test_gameOfLife(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				board: [][]int{
					{0, 1, 0},
					{0, 0, 1},
					{1, 1, 1},
					{0, 0, 0},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		{
			name: "2",
			args: args{
				board: [][]int{
					{1, 1},
					{1, 0},
				},
			},
			want: [][]int{
				{1, 1},
				{1, 1},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			gameOfLife(tt.args.board)

			for i := range tt.want {
				if slices.Compare(tt.args.board[i], tt.want[i]) != 0 {
					t.Errorf("gameOfLife %v got = %v, want %v", i, tt.args.board[i], tt.want[i])
				}
			}
		})
	}
}
