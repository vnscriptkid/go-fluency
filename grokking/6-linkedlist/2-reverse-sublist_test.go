package grokking

import (
	"testing"

	"github.com/vnscriptkid/go-fluency/grokking/shared"
)

func TestReverseSublist(t *testing.T) {
	node5 := shared.Node{Value: 5}
	node4 := shared.Node{Value: 4, Next: &node5}
	node3 := shared.Node{Value: 3, Next: &node4}
	node2 := shared.Node{Value: 2, Next: &node3}
	node1 := shared.Node{Value: 1, Next: &node2}

	r := reverseSublist(&node1, 2, 4)

	if r.GetPrinted() != "1->4->3->2->5" {
		panic("Expected `1->4->3->2->5` but got: " + r.GetPrinted())
	}
}
