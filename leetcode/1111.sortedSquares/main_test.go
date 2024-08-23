package main

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{0, 1, 2},
			},
			want: []int{0, 1, 4},
		},
		{
			name: "2",
			args: args{
				nums: []int{-4, -3, -2, -1, 0, 3, 10},
			},
			want: []int{0, 1, 4, 9, 9, 16, 100},
		},
		{
			name: "3",
			args: args{
				nums: []int{-1},
			},
			want: []int{1},
		},
		{
			name: "4",
			args: args{
				nums: []int{-5, -3, -2, -1},
			},
			want: []int{1, 4, 9, 25},
		},
		{
			name: "5",
			args: args{
				nums: []int{-4, -4, -3},
			},
			want: []int{9, 16, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
