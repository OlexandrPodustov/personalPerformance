package main

import (
	"testing"

	"gotest.tools/assert/cmp"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
				want:  []int{1, 2, 2, 3, 5, 6},
			},
		},
		{
			name: "2",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
				want:  []int{1},
			},
		},
		{
			name: "3",
			args: args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
				want:  []int{1},
			},
		},
		{
			name: "4",
			args: args{
				nums1: []int{2, 0},
				m:     1,
				nums2: []int{1},
				n:     1,
				want:  []int{1, 2},
			},
		},
		{
			name: "5",
			args: args{
				nums1: []int{4, 5, 6, 0, 0, 0},
				m:     3,
				nums2: []int{1, 2, 3},
				n:     3,
				want:  []int{1, 2, 3, 4, 5, 6},
			},
		},
		{
			name: "6",
			args: args{
				nums1: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				m:     0,
				nums2: []int{-50, -50, -48, -47, -44, -44, -37, -35, -35, -32, -32, -31, -29, -29, -28, -26, -24, -23, -23, -21, -20, -19, -17, -15, -14, -12, -12, -11, -10, -9, -8, -5, -2, -2, 1, 1, 3, 4, 4, 7, 7, 7, 9, 10, 11, 12, 14, 16, 17, 18, 21, 21, 24, 31, 33, 34, 35, 36, 41, 41, 46, 48, 48},
				n:     63,
				want:  []int{-50, -50, -48, -47, -44, -44, -37, -35, -35, -32, -32, -31, -29, -29, -28, -26, -24, -23, -23, -21, -20, -19, -17, -15, -14, -12, -12, -11, -10, -9, -8, -5, -2, -2, 1, 1, 3, 4, 4, 7, 7, 7, 9, 10, 11, 12, 14, 16, 17, 18, 21, 21, 24, 31, 33, 34, 35, 36, 41, 41, 46, 48, 48},
			},
		},
		{
			name: "7",
			args: args{
				nums1: []int{0, 0, 0},
				m:     0,
				nums2: []int{-2, -2, 1},
				n:     3,
				want:  []int{-2, -2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)

			c := cmp.DeepEqual(tt.args.nums1, tt.args.want)
			res := c()
			if !res.Success() {
				t.Errorf("got, but expected %v %v ", tt.args.nums1, tt.args.want)
			}
		})
	}
}
