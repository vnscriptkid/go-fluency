package grokking

import "fmt"

type Node struct {
	value int
	next  *Node
}

func hasCycle(head *Node) bool {
	// 1 -> 2 -> 3 -> 4 -> 5 -> 6
	//           ^_____________|
	//                    fs

	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	node6 := Node{value: 6}
	node5 := Node{value: 5, next: &node6}
	node4 := Node{value: 4, next: &node5}
	node3 := Node{value: 3, next: &node4}
	node2 := Node{value: 2, next: &node3}
	node1 := Node{value: 1, next: &node2}

	node6.next = &node3

	fmt.Println(hasCycle(&node1))
}
