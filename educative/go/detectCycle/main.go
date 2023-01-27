package main

import (
	"fmt"
)

// func detectCycle(head *EduLinkedListNode) bool {
// 	slow, fast := head, head

// 	for fast != nil && fast.next != nil {
// 		slow = slow.next
// 		fast = fast.next.next

// 		if slow == fast {
// 			break
// 		}
// 	}

// 	return false
// }

func detectCycle(head *EduLinkedListNode) bool {
	slow := head
	fast := head
	i, j := 0, 0

	// Run the loop until we reach the end of the linked list or find a cycle
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		i += 1
		j += 2

		// Break the loop if the fast pointer reach the end of the linked list
		if fast == nil {
			break
		}
		if slow == fast {
			return true
		}
	}
	return false
}

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

func NewLinkedListNode(data int, next *EduLinkedListNode) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = next
	return node
}

func InitLinkedListNode(data int) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = nil
	return node
}

type EduLinkedList struct {
	head *EduLinkedListNode
}

/*
InsertNodeAtHead method will insert a LinkedListNode at head

	of a linked list.
*/
func (l *EduLinkedList) InsertNodeAtHead(node *EduLinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

/*
	CreateLinkedList method will create the linked list using

the given integer array with the help of InsertAthead method.
*/
func (l *EduLinkedList) CreateLinkedList(lst []int) {
	for i := len(lst) - 1; i >= 0; i-- {
		newNode := InitLinkedListNode(lst[i])
		l.InsertNodeAtHead(newNode)
	}
}

// DisplayLinkedList method will display the elements of linked list.
func (l *EduLinkedList) DisplayLinkedList() {
	temp := l.head
	fmt.Print("[")
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
}
