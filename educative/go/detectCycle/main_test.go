package main

import (
	"fmt"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *EduLinkedListNode
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				head: func() *EduLinkedListNode {
					el := &EduLinkedList{}
					el.CreateLinkedList([]int{2, 4, 6, 8, 10})
					el.DisplayLinkedList()
					fmt.Println()

					return el.head
				}(),
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				head: func() *EduLinkedListNode {
					el := &EduLinkedList{}
					el.CreateLinkedList([]int{1, 3, 5, 7, 9})
					el.DisplayLinkedList()
					fmt.Println()

					return el.head
				}(),
			},
			want: false,
		},

		{
			name: "3",
			args: args{
				head: func() *EduLinkedListNode {
					el := &EduLinkedList{}
					el.CreateLinkedList([]int{1, 3, 5, 7, 9})
					el.DisplayLinkedList()
					fmt.Println()

					el.InsertNodeAtHead(el.head.next.next.next.next)

					return el.head
				}(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); got != tt.want {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
