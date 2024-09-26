package main

import "testing"

func Test_validMountainArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				arr: []int{2, 1},
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				arr: []int{3, 5, 5},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				arr: []int{0, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				arr: []int{0, 3, 2, 3},
			},
			want: false,
		},
		{
			name: "5",
			args: args{
				arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			want: false,
		},
		{
			name: "6",
			args: args{
				arr: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			},
			want: false,
		},
		// [9,8,7,6,5,4,3,2,1,0]
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
