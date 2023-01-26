package grokking

import (
	"container/list"

	"github.com/vnscriptkid/go-fluency/grokking/grokking/shared"
)

func averageLevels(root *shared.TreeNode) []float32 {
	result := []float32{}

	//      1
	//     / \
	//    2   3
	//   /\   /\
	//  4  5 6  7

	queue := list.New()

	queue.PushBack(root)

	for queue.Len() > 0 {
		levelSize := queue.Len()
		oldLevelSize := levelSize
		sum := 0

		for levelSize > 0 {
			node := queue.Remove(queue.Front()).(*shared.TreeNode)

			sum += node.Value

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}

			levelSize--
		}

		result = append(result, float32(sum)/float32(oldLevelSize))
	}

	return result
}
