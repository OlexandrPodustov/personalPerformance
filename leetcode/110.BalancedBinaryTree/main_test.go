package main

import (
	"testing"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "3,9,20,null,null,15,7",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val: 15,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "1,2,2,3,3,null,null,4,4",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 4,
							},
							Right: &TreeNode{
								Val: 4,
							},
						},
						Right: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: false,
		},
		{
			name: "[]",
			args: args{
				root: nil,
			},
			want: true,
		},
		{
			name: "[1,2,2,3,null,null,3,4,null,null,4]",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}},
					Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
				},
			},
			want: false,
		},
		{
			name: "1,null,2,null,3",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 3,
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced(%v) = %v, want %v", tt.args.root, got, tt.want)
			}
		})
	}
}
