package main

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1,2,3,1",
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: true,
		},
		{
			name: "1,2,3,4",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "minimum_amount_of_elements",
			args: args{
				nums: []int{1},
			},
			want: false,
		},
		{
			name: "1,1,1,3,3,4,3,2,4,2",
			args: args{
				nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
