package main

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 1, 0, 1, 1, 1},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 0, 1, 1, 0, 1},
			},
			want: 2,
		},
		{
			name: "3",
			args: args{
				nums: []int{0},
			},
			want: 0,
		},
		{
			name: "4",
			args: args{
				nums: []int{0, 0},
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				nums: []int{1, 0},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
