package grokking

import "github.com/vnscriptkid/go-fluency/grokking/shared"

func dfs(head *shared.TreeNode, curNum int, result map[string]int) {
	if head == nil {
		return
	}

	// visit cur node
	newNum := curNum*10 + head.Value

	// lead node, done cur path
	if head.Left == nil && head.Right == nil {
		result["total"] += newNum
		return
	}

	if head.Left != nil {
		dfs(head.Left, newNum, result)
	}

	if head.Right != nil {
		dfs(head.Right, newNum, result)
	}
}

func sumOfPathNumbers(head *shared.TreeNode) int {
	result := map[string]int{}

	result["total"] = 0

	dfs(head, 0, result)

	return result["total"]
}
