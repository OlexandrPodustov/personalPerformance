package main

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// {
		// 	name: "[1,2,3,4,5,6,7]",
		// 	args: args{
		// 		root: &Node{
		// 			Val: 1,
		// 			Left: &Node{
		// 				Val: 2,
		// 				Left: &Node{
		// 					Val: 4,
		// 				},
		// 				Right: &Node{
		// 					Val: 5,
		// 				},
		// 			},
		// 			Right: &Node{
		// 				Val: 3,
		// 				Left: &Node{
		// 					Val: 6,
		// 				},
		// 				Right: &Node{
		// 					Val: 7,
		// 				},
		// 			},
		// 		},
		// 	},
		// 	want: &Node{
		// 		Val: 1,
		// 		Left: &Node{
		// 			Val:   2,
		// 			Left:  &Node{Val: 4},
		// 			Right: &Node{Val: 5},
		// 			Next:  &Node{},
		// 		},
		// 		Right: &Node{
		// 			Val:   3,
		// 			Left:  &Node{Val: 6},
		// 			Right: &Node{Val: 7},
		// 			Next:  &Node{},
		// 		},
		// 		Next: nil,
		// 	},
		// },
		{
			name: "empty",
			args: args{
				root: &Node{},
			},
			want: &Node{},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect(%v) =\n%v,\nwant\n%v", tt.args.root, got, tt.want)
			}
		})
	}
}
