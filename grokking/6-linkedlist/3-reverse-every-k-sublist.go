package grokking

import "github.com/vnscriptkid/go-fluency/grokking/shared"

func reverseEveryKEleSublist(head *shared.Node, k int) *shared.Node {
	// 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> nil

	// 1 <- 2 <- 3    4 <- 5 <- 6    7 <- 8 -> nil
	//                                    p    c    t
	// e1        s1   e2        s2

	// s1 is global head
	// e1 points to s2

	var prev *shared.Node
	var prevSub *shared.Node = nil
	cur := head

	for cur != nil {
		prev = nil
		k1 := k

		e1 := cur

		// even sublist
		for k1 > 0 && cur != nil {
			temp := cur.Next
			cur.Next = prev

			prev = cur
			cur = temp

			k1--
		}

		if prevSub == nil {
			head = prev
		} else {
			prevSub.Next = prev
		}

		k2 := k

		e2 := cur

		// odd sublist
		for k2 > 0 && cur != nil {
			temp := cur.Next
			cur.Next = prev

			prev = cur
			cur = temp

			k2--
		}

		if cur != nil {
			e1.Next = prev
			prevSub = e2
		}
	}

	return head
}
