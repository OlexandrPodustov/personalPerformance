package main

import "testing"

func Test_checkIfExist(t *testing.T) {
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
				arr: []int{3, 1, 7, 11},
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				arr: []int{-2, 0, 10, -19, 4, 6, -8},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				arr: []int{0, 0},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfExist(tt.args.arr); got != tt.want {
				t.Errorf("checkIfExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
