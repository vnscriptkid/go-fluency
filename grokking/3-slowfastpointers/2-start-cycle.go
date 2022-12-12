package grokking

func startOfCycle(head *Node) *Node {
	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			break
		}
	}

	slow2 := head

	for {
		slow = slow.next
		slow2 = slow2.next

		if slow == slow2 {
			break
		}
	}

	return slow2
}
