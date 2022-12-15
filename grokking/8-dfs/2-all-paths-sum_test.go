package grokking

import (
	"testing"

	"github.com/vnscriptkid/go-fluency/grokking/shared"
)

func TestAllPathsForSum(t *testing.T) {
	node4 := shared.TreeNode{Value: 4}
	node5 := shared.TreeNode{Value: 5}
	node6 := shared.TreeNode{Value: 6}
	node3b := shared.TreeNode{Value: 3}
	node2 := shared.TreeNode{Value: 2, Left: &node4, Right: &node5}
	node3 := shared.TreeNode{Value: 3, Left: &node6, Right: &node3b}
	node1 := shared.TreeNode{Value: 1, Left: &node2, Right: &node3}

	r := allPathsForSum(&node1, 7)

	if r != 2 {
		panic("Not equal to 2")
	}

	//     1
	//   /   \
	//  2     3
	// /\     /\
	//4  5   6  3
}
