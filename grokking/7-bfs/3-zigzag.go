package grokking

import (
	"container/list"

	"github.com/vnscriptkid/go-fluency/grokking/shared"
)

func zigzagTraversal(head *shared.TreeNode) [][]int {
	//     1
	//   /   \
	//  2     3
	// /\     /\
	//4  5   6  7

	// queue: | (4) (5) (6) (7)
	// levelSize: 0
	// out (3)

	// level:
	// result: [ [2 3], [1], _ ]

	queue := list.New()

	queue.PushBack(head)
	result := [][]int{}

	for queue.Len() > 0 {
		levelSize := queue.Len()

		level := []int{}

		for levelSize > 0 {

			nodeOut := queue.Remove(queue.Front()).(*shared.TreeNode)
			level = append(level, nodeOut.Value)

			if nodeOut.Left != nil {
				queue.PushBack(nodeOut.Left)
			}

			if nodeOut.Right != nil {
				queue.PushBack(nodeOut.Right)
			}

			levelSize--
		}

		result = append([][]int{level}, result...)
	}

	return result
}
