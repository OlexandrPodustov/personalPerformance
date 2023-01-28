package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_palindrome(t *testing.T) {
	type args struct {
		head *EduLinkedListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[1,2,3,2,1]",
			args: args{
				head: func() *EduLinkedListNode {
					ell := EduLinkedList{}

					data := "1,2,3,2,1"
					ss := strings.Split(data, ",")
					input := []int{}
					for _, v := range ss {
						in, err := strconv.Atoi(v)
						if err != nil {
							t.Error(err)
							return nil
						}

						input = append(input, in)
					}
					ell.CreateLinkedList(input)
					ell.DisplayLinkedList()
					fmt.Println()

					return ell.head
				}(),
			},
			want: true,
		},

		{
			name: "4,7,9,5,4",
			args: args{
				head: func() *EduLinkedListNode {
					ell := EduLinkedList{}

					data := "4,7,9,5,4"
					ss := strings.Split(data, ",")
					input := []int{}
					for _, v := range ss {
						in, err := strconv.Atoi(v)
						if err != nil {
							t.Error(err)
							return nil
						}

						input = append(input, in)
					}
					ell.CreateLinkedList(input)
					ell.DisplayLinkedList()
					fmt.Println()

					return ell.head
				}(),
			},
			want: false,
		},

		{
			name: "2,3,5,5,3,2",
			args: args{
				head: func() *EduLinkedListNode {
					ell := EduLinkedList{}

					data := "2,3,5,5,3,2"
					ss := strings.Split(data, ",")
					input := []int{}
					for _, v := range ss {
						in, err := strconv.Atoi(v)
						if err != nil {
							t.Error(err)
							return nil
						}

						input = append(input, in)
					}
					ell.CreateLinkedList(input)
					ell.DisplayLinkedList()
					fmt.Println()

					return ell.head
				}(),
			},
			want: true,
		},

		{
			name: "6,1,0,5,1,6",
			args: args{
				head: func() *EduLinkedListNode {
					ell := EduLinkedList{}

					data := "6,1,0,5,1,6"
					ss := strings.Split(data, ",")
					input := []int{}
					for _, v := range ss {
						in, err := strconv.Atoi(v)
						if err != nil {
							t.Error(err)
							return nil
						}

						input = append(input, in)
					}
					ell.CreateLinkedList(input)
					ell.DisplayLinkedList()
					fmt.Println()

					return ell.head
				}(),
			},
			want: false,
		},

		{
			name: "3,6,9,8,4,8,9,6,3",
			args: args{
				head: func() *EduLinkedListNode {
					ell := EduLinkedList{}

					data := "3,6,9,8,4,8,9,6,3"
					ss := strings.Split(data, ",")
					input := []int{}
					for _, v := range ss {
						in, err := strconv.Atoi(v)
						if err != nil {
							t.Error(err)
							return nil
						}

						input = append(input, in)
					}
					ell.CreateLinkedList(input)
					ell.DisplayLinkedList()
					fmt.Println()

					return ell.head
				}(),
			},
			want: true,
		},

		{
			name: "0,0,0,0,0,0,0,0",
			args: args{
				head: func() *EduLinkedListNode {
					ell := EduLinkedList{}

					data := "0,0,0,0,0,0,0,0"
					ss := strings.Split(data, ",")
					input := []int{}
					for _, v := range ss {
						in, err := strconv.Atoi(v)
						if err != nil {
							t.Error(err)
							return nil
						}

						input = append(input, in)
					}
					ell.CreateLinkedList(input)
					ell.DisplayLinkedList()
					fmt.Println()

					return ell.head
				}(),
			},
			want: true,
		},

		{
			name: "7,5,3,1,3,5,3",
			args: args{
				head: func() *EduLinkedListNode {
					ell := EduLinkedList{}

					data := "7,5,3,1,3,5,3"
					ss := strings.Split(data, ",")
					input := []int{}
					for _, v := range ss {
						in, err := strconv.Atoi(v)
						if err != nil {
							t.Error(err)
							return nil
						}

						input = append(input, in)
					}
					ell.CreateLinkedList(input)
					ell.DisplayLinkedList()
					fmt.Println()

					return ell.head
				}(),
			},
			want: false,
		},

		{
			name: "1,2,3,2,1,2,3",
			args: args{
				head: func() *EduLinkedListNode {
					ell := EduLinkedList{}

					data := "1,2,3,2,1,2,3"
					ss := strings.Split(data, ",")
					input := []int{}
					for _, v := range ss {
						in, err := strconv.Atoi(v)
						if err != nil {
							t.Error(err)
							return nil
						}

						input = append(input, in)
					}
					ell.CreateLinkedList(input)
					ell.DisplayLinkedList()
					fmt.Println()

					return ell.head
				}(),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindrome(tt.args.head); got != tt.want {
				t.Errorf("palindrome() = %v, want %v", got, tt.want)
				t.FailNow()
			}
		})
	}
}
