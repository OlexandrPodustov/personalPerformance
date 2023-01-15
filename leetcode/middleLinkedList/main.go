package main

import (
	"fmt"
)

func main() {
	sl := []int{10, 20, 30}
	// sl := []int{1, 2, 3, 4, 5}

	ll := composeLl(sl)
	// spew.Dump(ll)

	fmt.Println(middleNode(&ll))
}

func composeLl(sl []int) ListNode {
	var linkedList ListNode

	for i, v := range sl {
		// fmt.Println("i,v", i, v)
		if i == 0 {
			linkedList.Val = v
			linkedList.Next = &ListNode{}
		} else {
			curLast := linkedList.Next
			// fmt.Println("curLast", v, curLast)

			for {
				if curLast.Next != nil {
					// fmt.Println("curLast.Next != nil", v, curLast)
					curLast = curLast.Next
				} else {
					// fmt.Println("curLast.Next == nil", v, curLast)
					curLast.Val = v
					curLast.Next = &ListNode{}
					break
				}
			}

		}
		// spew.Dump(i, v, linkedList)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	// spew.Dump(linkedList)

	return linkedList
}

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	return head
}
