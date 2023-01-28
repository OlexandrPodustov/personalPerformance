package main

import "testing"

func Test_circularArrayLoop(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example [1 3 -2 -4 1]",
			args: args{
				arr: []int{1, 3, -2, -4, 1},
			},
			want: true,
		},
		{
			name: "1 [3 -3 2 -2]",
			args: args{
				arr: []int{3, -3, 2, -2}, // The loop direction is changed for every possible loop starting at each index. So, the loop cannot be detected.
			},
			want: false,
		},
		{
			name: "2 [1 4 3 2 1]",
			args: args{
				arr: []int{1, 4, 3, 2, 1}, // Loop detected starting from the first index of the array.
			},
			want: true,
		},
		{
			name: "3 [-2 -3 1 -3 2]",
			args: args{
				arr: []int{-2, -3, 1, -3, 2}, // Loop detected starting from the first index of the array.
			},
			want: true,
		},
		{
			name: "4 [5 -1 1 1 -7 -9]",
			args: args{
				arr: []int{5, -1, 1, 1, -7, -9}, // The loop direction is changed for every possible loop starting at each index. So, the loop cannot be detected.
			},
			want: false,
		},
		{
			name: "5 [2 5 -4 3 -1 4]",
			args: args{
				arr: []int{2, 5, -4, 3, -1, 4}, // The loop direction is changed for every possible loop starting at each index. So, the loop cannot be detected.
			},
			want: false,
		},

		{
			name: "6 [3 3 1 -1 2]",
			args: args{
				arr: []int{3, 3, 1, -1, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := circularArrayLoop(tt.args.arr); got != tt.want {
				t.Errorf("circularArrayLoop() = %v, want %v", got, tt.want)
				t.FailNow()
			}
		})
	}
}
