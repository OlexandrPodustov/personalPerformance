package main

import "testing"

func Test_findSumOfThree(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				nums:   []int{1, 3, 5},
				target: 9,
			},
			want: true,
		},

		{
			name: "2",
			args: args{
				nums:   []int{1, 3, 4, 6, 8, 20},
				target: 31,
			},
			want: true,
		},

		{
			name: "3",
			args: args{
				nums:   []int{3, 7, 1, 2, 8, 4, 5},
				target: 20,
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				nums:   []int{-1, 2, 1, -4, 5, 5, 5, -3},
				target: -8,
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				nums:   []int{1, -1, 0},
				target: -1,
			},
			want: false,
		},

		{
			name: "1_neg",
			args: args{
				nums:   []int{3, 7, 1, 2, 8, 4, 5},
				target: 21,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSumOfThree(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findSumOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
