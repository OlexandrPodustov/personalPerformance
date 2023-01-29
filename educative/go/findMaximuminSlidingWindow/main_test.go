package main

import (
	"reflect"
	"testing"
)

func Test_findMaxSlidingWindow(t *testing.T) {
	type args struct {
		nums       []int
		windowSize int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// {
		// 	name: "testcase from the note: If the window size is greater than the array size, we consider the entire array as a single window.			",
		// 	args: args{
		// 		nums:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		// 		windowSize: 10,
		// 	},
		// 	want: []int{10},
		// },
		// {
		// 	name: "1",
		// 	args: args{
		// 		nums:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		// 		windowSize: 3,
		// 	},
		// 	want: []int{3, 4, 5, 6, 7, 8, 9, 10},
		// },

		{
			name: "2",
			args: args{
				nums:       []int{-4, 2, -5, 3, 6},
				windowSize: 3,
			},
			want: []int{2, 3, 6},
		},
		{
			name: "3",
			args: args{
				nums:       []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				windowSize: 4,
			},
			want: []int{4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "4",
			args: args{
				nums:       []int{1, 2, 3, 4, 5, 6},
				windowSize: 6,
			},
			want: []int{6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxSlidingWindow(tt.args.nums, tt.args.windowSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMaxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
