package grokking

import (
	"reflect"
	"testing"

	"github.com/vnscriptkid/go-fluency/grokking/shared"
)

func TestReverseLevelOrderTraverse(t *testing.T) {
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

	result := reverseLevelOrderTraverse(&node1)

	if !reflect.DeepEqual(result, [][]int{{1}, {3, 2}, {4, 5, 6, 7}}) {
		panic("Does not match")
	}
}
