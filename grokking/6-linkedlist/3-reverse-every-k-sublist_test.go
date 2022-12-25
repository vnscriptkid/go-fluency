package grokking

import (
	"testing"

	"github.com/vnscriptkid/go-fluency/grokking/shared"
)

func TestReverseEveryKEleSublist(t *testing.T) {
	node8 := shared.Node{Value: 8}
	node7 := shared.Node{Value: 7, Next: &node8}
	node6 := shared.Node{Value: 6, Next: &node7}
	node5 := shared.Node{Value: 5, Next: &node6}
	node4 := shared.Node{Value: 4, Next: &node5}
	node3 := shared.Node{Value: 3, Next: &node4}
	node2 := shared.Node{Value: 2, Next: &node3}
	node1 := shared.Node{Value: 1, Next: &node2}

	newHead := reverseEveryKEleSublist(&node1, 3)

	expected := "3->2->1->6->5->4->8->7"

	if newHead.GetPrinted() != expected {
		panic("Expected: " + expected + "\nActual: " + newHead.GetPrinted())
	}
}
