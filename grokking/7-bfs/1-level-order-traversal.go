package grokking

import (
	"container/list"
	"fmt"

	"github.com/vnscriptkid/go-fluency/grokking/shared"
)

func levelOrderTraverse(root *shared.TreeNode) [][]int {

	result := [][]int{}

	//     1
	//   /   \
	//  2     3
	// /\     /\
	//4  5   6  7

	queue := list.New()

	queue.PushBack(root)

	for queue.Len() > 0 {
		// Store size of queue now, which is number of items in current level
		levelSize := queue.Len()
		levelItems := []int{}

		// Process one level, push next-level items to queue
		for levelSize > 0 {
			queueItem := queue.Front()

			treeNode := queueItem.Value.(*shared.TreeNode)

			levelItems = append(levelItems, treeNode.Value)

			if treeNode.Left != nil {
				queue.PushBack(treeNode.Left)
			}

			if treeNode.Right != nil {
				queue.PushBack(treeNode.Right)
			}

			queue.Remove(queueItem)

			levelSize--
		}

		result = append(result, levelItems)
	}

	// queue: []
	// level: [4,5,6,7]
	// result: [[1], [2,3], [4,5,6,7]]

	return result
}

func main() {
	node4 := shared.TreeNode{Value: 4}
	node5 := shared.TreeNode{Value: 5}
	node6 := shared.TreeNode{Value: 6}
	node7 := shared.TreeNode{Value: 7}
	node2 := shared.TreeNode{Value: 2, Left: &node4, Right: &node5}
	node3 := shared.TreeNode{Value: 3, Left: &node6, Right: &node7}
	node1 := shared.TreeNode{Value: 1, Left: &node2, Right: &node3}

	result := levelOrderTraverse(&node1)

	fmt.Println(result)
}
