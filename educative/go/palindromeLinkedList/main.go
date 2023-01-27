package main

// palindrome function checks if the given linked list is a palindrome or not
func palindrome(head *EduLinkedListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	reversed := ReverseList(slow)
	slow, fast = head, head
	isPalindrome := true

	for reversed != nil && slow.next != nil {
		// fmt.Println()
		// fmt.Println("slow and reversed.data", slow, reversed)
		if slow.data != reversed.data {
			isPalindrome = false
			break
		}

		slow = slow.next
		reversed = reversed.next
	}
	// fmt.Println("isPalindrome", isPalindrome)

	return isPalindrome
}
