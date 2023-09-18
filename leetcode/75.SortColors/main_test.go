package main

import (
	"testing"

	"gotest.tools/v3/assert/cmp"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name           string
		args           args
		expectedResult []int
	}{
		{
			name: "Example_1",
			args: args{
				nums: []int{2, 0, 2, 1, 1, 0},
			},
			expectedResult: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "Example_2",
			args: args{
				nums: []int{2, 0, 1},
			},
			expectedResult: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
			c := cmp.DeepEqual(tt.args.nums, tt.expectedResult)
			res := c()
			if !res.Success() {
				t.Errorf("got, but expected %v %v ", tt.args.nums, tt.expectedResult)
			}
		})
	}
}
