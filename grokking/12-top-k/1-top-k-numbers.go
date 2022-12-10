package grokking

import (
	"container/heap"
)

type IntHeap []int

// ///////////////////////
// Impl Heap interface //
// ///////////////////////
func (intHeap *IntHeap) Push(x interface{}) {
	*intHeap = append(*intHeap, x.(int))
}

func (intHeap *IntHeap) Pop() interface{} {
	prev := *intHeap

	lastEle := prev[len(prev)-1]

	*intHeap = prev[0 : len(prev)-1]

	return lastEle
}

// ///////////////////////
// Impl Sort interface //
// ///////////////////////
func (intHeap IntHeap) Len() int {
	return len(intHeap)
}

// Min Heap (To find top k largest)
func (intHeap IntHeap) Less(i, j int) bool {
	return intHeap[i] < intHeap[j]
}

func (intHeap IntHeap) Swap(i, j int) {
	intHeap[i], intHeap[j] = intHeap[j], intHeap[i]
}

func topKLargestNumbers(arr []int, k int) []int {
	intHeap := &IntHeap{}
	heap.Init(intHeap)

	for _, num := range arr {
		heap.Push(intHeap, num)

		if intHeap.Len() > k {
			// Len now is k + 1, smallest out
			heap.Pop(intHeap)
		}
	}

	return *intHeap
}
