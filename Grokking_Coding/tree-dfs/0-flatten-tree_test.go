package TreeDFS

import (
	"fmt"
	"strings"
	"testing"
)

func Test_flattenBinaryTree(t *testing.T) {
	node1 := NewBinaryTreeNode(1)
	node4 := NewBinaryTreeNode(4)
	node2 := NewBinaryTreeNode(2)
	node3 := NewBinaryTreeNode(3)
	node5 := NewBinaryTreeNode(5)
	node17 := NewBinaryTreeNode(17)
	node19 := NewBinaryTreeNode(19)

	node2.left = node1
	node2.right = node4

	node17.left = node19
	node17.right = node5

	node3.left = node2
	node3.right = node17

	newRoot := flattenBinaryTree(node3)

	visualL := strings.Builder{}

	p := newRoot

	for p != nil {
		visualL.WriteString(fmt.Sprint(p.data))
		visualL.WriteString("->")
		p = p.next
	}

	visualL.WriteString("nil")

	actual := visualL.String()
	expected := "3->2->1->4->17->19->5->nil"

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
