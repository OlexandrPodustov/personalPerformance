package main

import "testing"

func Test_furthestBuilding(t *testing.T) {
	type args struct {
		heights []int
		bricks  int
		ladders int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				heights: []int{4, 2, 7, 6, 9, 14, 12},
				bricks:  5,
				ladders: 1,
			},
			want: 4,
		},
		{
			name: "2",
			args: args{
				heights: []int{4, 12, 2, 7, 3, 18, 20, 3, 19},
				bricks:  10,
				ladders: 2,
			},
			want: 7,
		},
		{
			name: "3",
			args: args{
				heights: []int{14, 3, 19, 3},
				bricks:  17,
				ladders: 0,
			},
			want: 3,
		},
		{
			name: "4",
			args: args{
				heights: []int{3, 19},
				bricks:  87,
				ladders: 1,
			},
			want: 1,
		},
		{
			name: "5",
			args: args{
				heights: []int{2, 7, 9, 12},
				bricks:  5,
				ladders: 1,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := furthestBuilding(tt.args.heights, tt.args.bricks, tt.args.ladders); got != tt.want {
				t.Errorf("furthestBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}
