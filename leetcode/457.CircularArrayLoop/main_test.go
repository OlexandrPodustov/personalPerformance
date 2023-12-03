package main

import "testing"

func Test_circularArrayLoop(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				nums: []int{2, -1, 1, 2, 2},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				nums: []int{-1, -2, -3, -4, -5, 6},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				nums: []int{1, -1, 5, 1, 4},
			},
			want: true,
		},
		// {
		// 	name: "4",
		// 	args: args{
		// 		nums: []int{-2, 1, -1, -2, -2},
		// 	},
		// 	want: false,
		// },
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := circularArrayLoop(tt.args.nums); got != tt.want {
				t.Errorf("circularArrayLoop(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
