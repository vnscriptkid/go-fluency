package grokking

import "github.com/vnscriptkid/go-fluency/grokking/shared"

func pathWithGivenSeqRecur(root *shared.TreeNode, seq []int, curIdx int) bool {
	if root == nil {
		return false
	}

	if curIdx == len(seq) {
		return false
	}

	if root.Value != seq[curIdx] {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return curIdx == len(seq)-1 && seq[curIdx] == root.Value
	}

	return pathWithGivenSeqRecur(root.Left, seq, curIdx+1) || pathWithGivenSeqRecur(root.Right, seq, curIdx+1)
}

func pathWithGivenSeq(root *shared.TreeNode, seq []int) bool {
	return pathWithGivenSeqRecur(root, seq, 0)
}
