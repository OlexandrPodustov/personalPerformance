package main

import "testing"

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1,3,4,2,2",
			args: args{
				nums: []int{1, 3, 4, 2, 2},
			},
			want: 2,
		},
		{
			name: "3,1,3,4,2",
			args: args{
				nums: []int{3, 1, 3, 4, 2},
			},
			want: 3,
		},
		{
			name: "2,2,1",
			args: args{
				nums: []int{2, 2, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
