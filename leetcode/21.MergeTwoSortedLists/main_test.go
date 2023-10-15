package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Input: list1 = [1,2,4], list2 = [1,3,4]",
			args: args{
				list1: func() *ListNode {
					var li1 *ListNode
					li1 = appendToEnd(li1, 1)
					li1 = appendToEnd(li1, 2)
					li1 = appendToEnd(li1, 4)

					return li1
				}(),
				list2: func() *ListNode {
					var li2 *ListNode

					li2 = appendToEnd(li2, 1)
					li2 = appendToEnd(li2, 3)
					li2 = appendToEnd(li2, 4)
					return li2
				}(),
			},
			want: func() *ListNode {
				var liRes *ListNode
				// Output: [1,1,2,3,4,4]

				liRes = appendToEnd(liRes, 1)
				liRes = appendToEnd(liRes, 1)
				liRes = appendToEnd(liRes, 2)
				liRes = appendToEnd(liRes, 3)
				liRes = appendToEnd(liRes, 4)
				liRes = appendToEnd(liRes, 4)

				return liRes
			}(),
		},
		{
			name: "second",
			args: args{
				list1: func() *ListNode {
					var liRes *ListNode

					liRes = appendToEnd(liRes, 1)
					liRes = appendToEnd(liRes, 2)
					liRes = appendToEnd(liRes, 3)

					return liRes
				}(),
				list2: func() *ListNode {
					var liRes *ListNode

					liRes = appendToEnd(liRes, 4)
					liRes = appendToEnd(liRes, 5)

					return liRes
				}(),
			},
			want: func() *ListNode {
				var liRes *ListNode

				liRes = appendToEnd(liRes, 1)
				liRes = appendToEnd(liRes, 2)
				liRes = appendToEnd(liRes, 3)
				liRes = appendToEnd(liRes, 4)
				liRes = appendToEnd(liRes, 5)

				return liRes
			}(),
		},
		{
			name: "empty",
			args: args{
				list1: func() *ListNode {
					var liRes *ListNode

					return liRes
				}(),
				list2: func() *ListNode {
					var liRes *ListNode

					return liRes
				}(),
			},
			want: func() *ListNode {
				var liRes *ListNode

				return liRes
			}(),
		},
		{
			name: "l1 empty",
			args: args{
				list1: func() *ListNode {
					var liRes *ListNode

					return liRes
				}(),
				list2: func() *ListNode {
					var liRes *ListNode
					liRes = appendToEnd(liRes, 22)

					return liRes
				}(),
			},
			want: func() *ListNode {
				var liRes *ListNode
				liRes = appendToEnd(liRes, 22)

				return liRes
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				got1, _ := json.Marshal(got)
				// fmt.Println("got1", string(got1))
				waaa, _ := json.Marshal(tt.want)
				// fmt.Println("waaa", string(waaa))

				t.Errorf("mergeTwoLists() got\n%v\nwant\n%v", string(got1), string(waaa))
			}
		})
	}
}

func appendToEnd(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val, Next: nil}
	if head == nil { // initial step - used once
		return newNode
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode

	return head
}
