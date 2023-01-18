package main

// getMiddleNode function to find the middle node of the linked list
func getMiddleNode(head *EduLinkedListNode) *EduLinkedListNode {
	slow, fast := head, head

	for {
		if fast == nil || fast.next == nil {
			break
		}

		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			break
		}
	}

	return slow
}
