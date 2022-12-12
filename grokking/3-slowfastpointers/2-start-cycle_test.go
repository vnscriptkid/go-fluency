package grokking

import "testing"

func TestStartOfCycle(t *testing.T) {
	node6 := Node{value: 6}
	node5 := Node{value: 5, next: &node6}
	node4 := Node{value: 4, next: &node5}
	node3 := Node{value: 3, next: &node4}
	node2 := Node{value: 2, next: &node3}
	node1 := Node{value: 1, next: &node2}

	node6.next = &node3

	r := startOfCycle(&node1)

	if r != &node3 {
		panic("Wrong start of cycle")
	}
}
