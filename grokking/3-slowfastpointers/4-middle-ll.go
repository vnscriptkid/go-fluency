package grokking

import "github.com/vnscriptkid/go-fluency/grokking/grokking/shared"

func middleOfLinkedList(head *shared.Node) int {
	// 1 -> 2 -> 3 -> 4 -> 5 -> nil
	//           ^
	//                     $

	// 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> null
	//                ^
	//                                $

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow.Value
}
