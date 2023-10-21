package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

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
