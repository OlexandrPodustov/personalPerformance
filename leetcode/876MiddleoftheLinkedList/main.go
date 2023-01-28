package main

import (
	"fmt"
)

func main() {
	sl := []int{10, 20, 30}

	ll := composeLl(sl)
	// spew.Dump(ll)

	fmt.Println(middleNode(ll))
}

func composeLl(list []int) *ListNode {
	linkedList := &ListNode{}
	linkedList = nil

	for i := 0; i < len(list)-1; i++ {
		fmt.Println("", list[i], list[i+1])
		if i == 0 {
			linkedList = &ListNode{
				Val: list[i],
				Next: &ListNode{
					Val:  list[i+1],
					Next: nil,
				},
			}
			continue
		}
		fmt.Println("linkedList", linkedList)
		fmt.Println("linkedList.Next", linkedList.Next)

		linkedList.Next = &ListNode{
			Val:  list[i],
			Next: nil,
		}
		linkedList = linkedList.Next
		fmt.Println("linkedList.Next", linkedList.Next)
	}
	// spew.Dump(linkedList)

	return linkedList
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fmt.Println("middleNode slow", slow)

		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
