package main

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "3,2,2,3",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "1",
			args: args{
				nums: []int{1},
				val:  1,
			},
			want: 0,
		},
		{
			name: "3,3",
			args: args{
				nums: []int{3, 3},
				val:  3,
			},
			want: 0,
		},
		{
			name: "4,5",
			args: args{
				nums: []int{4, 5},
				val:  4,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
