package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var newHead *ListNode
	for head != nil {
		newHead = &ListNode{Val: head.Val, Next: newHead}
		head = head.Next
	}

	return newHead
}
