package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	printList(head)

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fmt.Println("middleNode slow", slow)

		slow = slow.Next
		fast = fast.Next.Next
	}

	fmt.Println("middleNode slow result", slow)

	return slow
}

func composeLl(list []int) *ListNode {
	linkedList := &ListNode{}
	linkedList = nil

	for i := len(list) - 1; i >= 0; i-- {
		linkedList = &ListNode{Val: list[i], Next: linkedList}
	}

	return linkedList
}

func printList(head *ListNode) {
	place := head
	for place != nil {
		fmt.Println("printList ", place)
		place = place.Next
	}
}
