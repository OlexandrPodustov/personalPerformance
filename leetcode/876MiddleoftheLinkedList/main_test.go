package main

import (
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{

		{
			name: "1,2,3,4,5",
			args: args{
				head: composeLl([]int{1, 2, 3, 4, 5}),
			},
			want: &ListNode{Val: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); got.Val != tt.want.Val {
				t.Errorf("middleNode() = %v, want %v", got.Val, tt.want.Val)
			}
		})
	}
}
