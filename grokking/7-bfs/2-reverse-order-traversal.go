package grokking

import (
	"container/list"

	"github.com/vnscriptkid/go-fluency/grokking/shared"
)

func reverseLevelOrderTraverse(root *shared.TreeNode) [][]int {
	//     1
	//   /   \
	//  2     3
	// /\     /\
	//4  5   6  7

	// q: [] levelProcessed: 0
	// leftToRight := true
	// out: 7
	// [4 5 6 7]
	// result: [ [1] [3,2] ]

	r := [][]int{}

	q := list.List{}
	q.PushBack(root)

	lToR := true

	for q.Len() > 0 {
		c := q.Len()
		l := []int{}

		for c > 0 {
			o := q.Remove(q.Front()).(*shared.TreeNode)

			if lToR {
				l = append(l, o.Value)
			} else {
				// like prepend
				l = append([]int{o.Value}, l...)
			}

			if o.Left != nil {
				q.PushBack(o.Left)
			}

			if o.Right != nil {
				q.PushBack(o.Right)
			}

			c--
		}

		lToR = !lToR

		r = append(r, l)
	}

	return r
}
