package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_solution(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums1: []int{4, 7, 11, 0, 0},
				m:     3,
				nums2: []int{1, 6},
				n:     2,
			},
			want: []int{1, 4, 6, 7, 11},
		},
		{
			name: "2",
			args: args{
				nums1: []int{1, 2, 3, 4, 7, 11, 0, 0},
				m:     6,
				nums2: []int{5, 6},
				n:     2,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 11},
		},
		{
			name: "3",
			args: args{
				nums1: []int{5, 6, 7, 0, 0, 0},
				m:     3,
				nums2: []int{1, 2, 3},
				n:     3,
			},
			want: []int{1, 2, 3, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		})
		require.Equal(t, tt.want, tt.args.nums1)
	}
}
