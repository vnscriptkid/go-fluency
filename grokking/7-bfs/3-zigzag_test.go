package grokking

import (
	"reflect"
	"testing"

	"github.com/vnscriptkid/go-fluency/grokking/shared"
)

func TestZigzagTraversal(t *testing.T) {
	//     1
	//   /   \
	//  2     3
	// /\     /\
	//4  5   6  7

	node4 := shared.TreeNode{Value: 4}
	node5 := shared.TreeNode{Value: 5}
	node6 := shared.TreeNode{Value: 6}
	node7 := shared.TreeNode{Value: 7}
	node2 := shared.TreeNode{Value: 2, Left: &node4, Right: &node5}
	node3 := shared.TreeNode{Value: 3, Left: &node6, Right: &node7}
	node1 := shared.TreeNode{Value: 1, Left: &node2, Right: &node3}

	result := zigzagTraversal(&node1)

	expected := [][]int{{4, 5, 6, 7}, {2, 3}, {1}}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

}
