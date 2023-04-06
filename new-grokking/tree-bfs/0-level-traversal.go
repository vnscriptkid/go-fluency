package TreeBFS

import (
	"container/list"
	"fmt"
	"strings"
)

func levelOrderTraversal(root *BinaryTreeNode) string {

	// Write your code here
	// your code will replace this placeholder return statement
	queue := list.New()

	queue.PushBack(root)

	result := strings.Builder{}

	for queue.Len() > 0 {
		levelSize := queue.Len()

		level := strings.Builder{}

		for levelSize > 0 {
			node := queue.Remove(queue.Front()).(*BinaryTreeNode)

			level.WriteString(fmt.Sprint(node.data))

			// not last item in the level
			if levelSize > 1 {
				level.WriteString(",")
			}

			if node.left != nil {
				queue.PushBack(node.left)
			}

			if node.right != nil {
				queue.PushBack(node.right)
			}

			levelSize--
		}

		// not the last level?
		if queue.Len() > 0 {
			level.WriteString(":")
		}

		result.WriteString(level.String())
	}

	return result.String()
}
