package main

import (
	"fmt"

	"github.com/vnscriptkid/go-fluency/grokking/shared"
)

func hasPath(head *shared.TreeNode, targetSum int) bool {
	//   1
	// /   \
	// 2    3

	// base case: leaf node
	if head == nil {
		return false
	}

	// leaf node
	if head.Left == nil && head.Right == nil {
		return head.Value == targetSum
	}

	// 2 choices: left, right
	// use recursion, update root node and target sum along the way
	return hasPath(head.Left, targetSum-head.Value) || hasPath(head.Right, targetSum-head.Value)
}

func main() {
	node4 := shared.TreeNode{Value: 4}
	node5 := shared.TreeNode{Value: 5}
	node6 := shared.TreeNode{Value: 6}
	node7 := shared.TreeNode{Value: 7}
	node2 := shared.TreeNode{Value: 2, Left: &node4, Right: &node5}
	node3 := shared.TreeNode{Value: 3, Left: &node6, Right: &node7}
	node1 := shared.TreeNode{Value: 1, Left: &node2, Right: &node3}

	fmt.Println(hasPath(&node1, 10))
	fmt.Println(hasPath(&node1, 11))
	fmt.Println(hasPath(&node1, 12))
}
