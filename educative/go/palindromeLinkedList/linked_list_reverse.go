// Template to reverse the linked list
package main

func ReverseList(slowPtr *EduLinkedListNode) *EduLinkedListNode {
	reverse := new(EduLinkedListNode)
	reverse = nil
	for slowPtr != nil {
		next := slowPtr.next
		slowPtr.next = reverse
		reverse = slowPtr
		slowPtr = next
	}
	return reverse
}
