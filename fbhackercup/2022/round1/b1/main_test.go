package main

import "testing"

func Test_solution(t *testing.T) {
	type args struct {
		treesLocations [][]int64
		wellLocations  [][]int64
		resssss        chan int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "1",
			args: args{
				treesLocations: [][]int64{{2, 2}, {5, 5}},
				wellLocations:  [][]int64{{2, 5}, {6, 6}},
				resssss:        make(chan int64, 1),
			},
			want: 52,
		},
		{
			name: "2",
			args: args{
				treesLocations: [][]int64{
					{1, 1},
					{4, 3},
					{6, 3},
					{6, 5},
				},
				wellLocations: [][]int64{
					{3, 1},
					{5, 2},
					{6, 5},
				},
				resssss: make(chan int64, 1),
			},
			want: 131,
		},
		{
			name: "3",
			args: args{
				treesLocations: [][]int64{
					{2837, 745},
					{62, 1162},
					{2634, 1112},
					{1746, 2618},
					{847, 127},
					{986, 1993},
					{732, 1273},
					{2003, 1998},
				},
				wellLocations: [][]int64{
					{1276, 2231},
					{1234, 1234},
					{287, 2371},
					{3000, 3000},
				},
				resssss: make(chan int64, 1),
			},
			want: 110090622,
		},
		{
			name: "999999999",
			args: args{
				treesLocations: [][]int64{
					{2837, 745},
					{62, 1162},
					{2634, 1112},
					{1746, 2618},
					{847, 127},
					{986, 1993},
					{732, 1273},
					{2003, 1998},
					{2003, 1998},
					{2003, 1998},
					{2003, 1998},
					{2003, 1998},
				},
				wellLocations: [][]int64{
					{maxABXY, maxABXY},
					{maxABXY, maxABXY},
					{maxABXY, maxABXY},
					{maxABXY, maxABXY},
					{maxABXY, maxABXY},
					{maxABXY, maxABXY},
					{maxABXY, maxABXY},
					{maxABXY, maxABXY},
					{maxABXY, maxABXY},
					{maxABXY, maxABXY},
					{maxABXY, maxABXY},
					{maxABXY, maxABXY},
				},
				resssss: make(chan int64, 1),
			},
			want: 918149076,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.treesLocations, tt.args.wellLocations, tt.args.resssss); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
				t.FailNow()
			}
		})
	}
}
