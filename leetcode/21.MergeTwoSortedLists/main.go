package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// func main() {
// 	var l *ListNode
// 	res := l.appendToEnd(1).appendToEnd(2).appendToEnd(3)

// 	val, _ := json.Marshal(res)
// 	fmt.Println("added main", string(val))
// }

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode
	for {
		if list1 == nil && list2 == nil {
			break
		}

		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				result = appendToEnd(result, list1.Val)
				list1 = list1.Next
			} else {
				result = appendToEnd(result, list2.Val)
				list2 = list2.Next
			}
		}

		if list1 == nil {
			for list2 != nil {
				result = appendToEnd(result, list2.Val)
				list2 = list2.Next
			}
		}
		if list2 == nil {
			for list1 != nil {
				result = appendToEnd(result, list1.Val)
				list1 = list1.Next
			}
		}
	}

	return result
}

func appendToEnd(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val, Next: nil}
	if head == nil { // initial step - used once
		return newNode
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode

	return head
}
