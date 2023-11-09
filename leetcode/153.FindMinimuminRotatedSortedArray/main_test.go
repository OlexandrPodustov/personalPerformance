package main

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4,5,6,7,0,1,2",
			args: args{
				nums: []int{4, 5, 6, 7, 0, 1, 2},
			},
			want: 0,
		},
		{
			name: "4,5,6,7,1,2",
			args: args{
				nums: []int{4, 5, 6, 7, 1, 2},
			},
			want: 1,
		},
		{
			name: "11,13,15,17",
			args: args{
				nums: []int{11, 13, 15, 17},
			},
			want: 11,
		},
		{
			name: "11,13,9,17",
			args: args{
				nums: []int{11, 13, 9, 17},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
