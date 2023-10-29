package main

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1, 5, 0, 3, 5",
			args: args{
				nums: []int{1, 5, 0, 3, 5},
			},
			want: 3,
		},
		{
			name: "1, 6, 0, 3, 5",
			args: args{
				nums: []int{1, 6, 0, 3, 5},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := minimumOperations(tt.args.nums); got != tt.want {
				t.Errorf("minimumOperations(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
