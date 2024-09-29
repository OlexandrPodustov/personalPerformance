package main

import "testing"

func Test_thirdMax(t *testing.T) {
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
				nums: []int{3, 2, 1},
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2},
			},
			want: 2,
		},

		{
			name: "3",
			args: args{
				nums: []int{2, 2, 3, 1},
			},
			want: 1,
		},

		{
			name: "4",
			args: args{
				nums: []int{1, 2, 2, 5, 3, 5},
			},
			want: 2,
		},
		{
			name: "5",
			args: args{
				nums: []int{5, 2, 4, 1, 3, 6, 0},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
