package grokking

import "container/heap"

type AMinHeap []ANode

type ANode struct {
	Value   int
	ListIdx int
	EleIdx  int
}

func (a *AMinHeap) Push(x interface{}) {
	*a = append(*a, x.(ANode))
}

func (a *AMinHeap) Pop() interface{} {
	prev := *a

	lastEle := prev[len(prev)-1]

	*a = prev[0 : len(prev)-1]

	return lastEle
}

func (a AMinHeap) Less(i, j int) bool {
	return a[i].Value < a[j].Value
}

func (a AMinHeap) Len() int {
	return len(a)
}

func (a AMinHeap) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func kthSmallestInMsortedLists(lists [][]int, k int) int {
	aMinHeap := &AMinHeap{}
	heap.Init(aMinHeap)

	asc := []int{}

	// init min heap
	for listIdx, list := range lists {
		node := ANode{
			Value:   list[0],
			ListIdx: listIdx,
			EleIdx:  0,
		}
		heap.Push(aMinHeap, node)
	}

	for len(asc) < k && len(*aMinHeap) > 0 {
		node := heap.Pop(aMinHeap).(ANode)

		asc = append(asc, node.Value)

		curList := lists[node.ListIdx]
		nextIdx := node.EleIdx + 1

		if nextIdx < len(curList) {
			heap.Push(aMinHeap, ANode{
				Value:   curList[nextIdx],
				ListIdx: node.ListIdx,
				EleIdx:  nextIdx,
			})
		}
	}

	return asc[len(asc)-1]
}
