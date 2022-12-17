package grokking

import (
	"container/heap"
)

type MaxIntHeap []int

func (h *MaxIntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxIntHeap) Pop() interface{} {
	prev := *h

	lastEle := prev[len(prev)-1]

	*h = prev[0 : len(prev)-1]

	return lastEle
}

func (h MaxIntHeap) Len() int {
	return len(h)
}

func (h MaxIntHeap) Peek() int {
	return h[0]
}

func (h MaxIntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxIntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func findKthSmallestNum(arr []int, k int) int {
	maxHeap := MaxIntHeap{}
	heap.Init(&maxHeap)

	// add k nums to maxHeap
	for i := 0; i < k; i++ {
		heap.Push(&maxHeap, arr[i])
	}

	for i := k; i < len(arr); i++ {
		if arr[i] < maxHeap.Peek() {
			heap.Pop(&maxHeap)
			heap.Push(&maxHeap, arr[i])
		}
	}

	return maxHeap.Peek()
}
