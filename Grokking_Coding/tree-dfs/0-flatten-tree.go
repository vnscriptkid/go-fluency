package TreeDFS

// flattenBinaryTree is used to populate same level pointers
func flattenBinaryTree(root *BinaryTreeNode) *BinaryTreeNode {
	preOrder(root, &Prev{})

	return root
}

type Prev struct {
	value *BinaryTreeNode
}

func preOrder(cur *BinaryTreeNode, prev *Prev) {
	if cur == nil {
		return
	}

	// visit cur
	if prev.value != nil {
		prev.value.next = cur
	}

	prev.value = cur

	// visit left
	preOrder(cur.left, prev)
	// visit right
	preOrder(cur.right, prev)
}
