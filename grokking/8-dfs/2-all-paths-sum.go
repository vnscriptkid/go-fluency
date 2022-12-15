package grokking

import (
	"github.com/vnscriptkid/go-fluency/grokking/shared"
)

func allPathsForSum(head *shared.TreeNode, targetSum int) int {
	if head == nil {
		return 0
	}

	if head.Left == nil && head.Right == nil {
		// Leaf node
		if head.Value == targetSum {
			return 1
		} else {
			return 0
		}
	}

	return allPathsForSum(head.Left, targetSum-head.Value) + allPathsForSum(head.Right, targetSum-head.Value)
}
