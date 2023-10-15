package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		result = list1
		list1 = list1.Next
	} else {
		result = list2
		list2 = list2.Next
	}
	result.Next = mergeTwoLists(list1, list2)

	return result
}
