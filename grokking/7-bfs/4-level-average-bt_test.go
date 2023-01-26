package grokking

import (
	"reflect"
	"testing"

	"github.com/vnscriptkid/go-fluency/grokking/grokking/shared"
)

func Test_averageLevels(t *testing.T) {
	node4 := shared.TreeNode{Value: 4}
	node5 := shared.TreeNode{Value: 5}
	node6 := shared.TreeNode{Value: 6}
	node7 := shared.TreeNode{Value: 7}
	node2 := shared.TreeNode{Value: 2, Left: &node4, Right: &node5}
	node3 := shared.TreeNode{Value: 3, Left: &node6, Right: &node7}
	node1 := shared.TreeNode{Value: 1, Left: &node2, Right: &node3}

	actual := averageLevels(&node1)

	expected := []float32{1, 2.5, 5.5}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
