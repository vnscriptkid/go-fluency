package grokking

import "container/heap"

type CustomNode struct {
	Value   int
	ListIdx int
	ItemIdx int
}

type CustomMinHeap []CustomNode

func (h *CustomMinHeap) Push(x any) {
	*h = append(*h, x.(CustomNode))
}
func (h *CustomMinHeap) Pop() any {
	prev := *h
	prevLen := len(prev)
	lastEle := prev[prevLen-1]
	*h = prev[0 : prevLen-1]
	return lastEle
}
func (h CustomMinHeap) Len() int {
	return len(h)
}
func (h CustomMinHeap) Less(i, j int) bool {
	return h[i].Value < h[j].Value
}
func (h CustomMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h CustomMinHeap) Top() interface{} {
	return h[0]
}
func (h CustomMinHeap) Empty() bool {
	return len(h) == 0
}

func findKthSmallestFromMatrix(matrix [][]int, k int) int {
	aMinHeap := CustomMinHeap{}
	heap.Init(&aMinHeap)

	// push first k node to the min heap
	for i, list := range matrix {
		node := CustomNode{
			Value:   list[0],
			ListIdx: i,
			ItemIdx: 0,
		}

		heap.Push(&aMinHeap, node)
	}

	count := 0

	for !aMinHeap.Empty() {
		nextNode := heap.Pop(&aMinHeap).(CustomNode)

		count++

		if count == k {
			return nextNode.Value
		}

		curList := matrix[nextNode.ListIdx]
		if nextNode.ItemIdx+1 < len(curList) {
			heap.Push(&aMinHeap, CustomNode{
				Value:   curList[nextNode.ItemIdx+1],
				ListIdx: nextNode.ListIdx,
				ItemIdx: nextNode.ItemIdx + 1,
			})
		}
	}

	return 0
}
