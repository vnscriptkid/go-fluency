package grokking

import "github.com/vnscriptkid/go-fluency/grokking/grokking/shared"

func reverseAlternatingSublist(head *shared.Node, k int) *shared.Node {
	// 1->2->3->4->5->6->7->8->nil
	//(____  ____)
	// 1<-2  3->4

	cur := head
	var prev *shared.Node

	var newHead *shared.Node
	var batchEnd *shared.Node

	for cur != nil {
		// reverse first k nodes (sub1)

		sub1End := cur

		i := k
		for cur != nil && i > 0 {
			// store next
			temp := cur.Next

			// point cur to prev
			cur.Next = prev

			// reassign cur, prev
			prev = cur
			cur = temp

			i--
		}

		if batchEnd == nil {
			newHead = prev
		} else {
			batchEnd.Next = prev
		}

		// connect sub1 and sub2
		sub1End.Next = cur

		// traverse next k nodes (sub2)
		i = k
		for cur != nil && i > 0 {
			prev = cur
			cur = cur.Next

			i--
		}

		batchEnd = prev
	}

	return newHead
}
