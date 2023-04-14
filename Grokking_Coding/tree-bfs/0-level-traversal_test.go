package TreeBFS

import "testing"

func Test_levelOrderTraversal(t *testing.T) {
	node100 := NewBinaryTreeNode(100)
	node50 := NewBinaryTreeNode(50)
	node200 := NewBinaryTreeNode(200)
	node25 := NewBinaryTreeNode(25)
	node75 := NewBinaryTreeNode(75)
	node350 := NewBinaryTreeNode(350)

	node100.left = node50
	node100.right = node200

	node50.left = node25
	node50.right = node75

	node200.right = node350

	actual := levelOrderTraversal(node100)
	expected := "100:50,200:25,75,350"

	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
