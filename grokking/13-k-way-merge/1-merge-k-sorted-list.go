package grokking

import "container/heap"

type HeapNode struct {
	Value   int
	ListIdx int
	EleIdx  int
}

type IntHeap []HeapNode

// Impl Heap interface
func (pIntHeap *IntHeap) Push(x interface{}) {
	*pIntHeap = append(*pIntHeap, x.(HeapNode))
}

func (pIntHeap *IntHeap) Pop() interface{} {
	prev := *pIntHeap

	lastEle := prev[len(prev)-1]

	*pIntHeap = prev[0 : len(prev)-1]

	return lastEle
}

// Impl Sort interface
func (intHeap IntHeap) Len() int {
	return len(intHeap)
}

func (intHeap IntHeap) Less(i, j int) bool {
	return intHeap[i].Value < intHeap[j].Value
}

func (intHeap IntHeap) Swap(i, j int) {
	intHeap[i], intHeap[j] = intHeap[j], intHeap[i]
}

func mergeKSortedLists(lists [][]int) []int {
	minHeap := IntHeap{}
	heap.Init(&minHeap)

	// Go through each list, take first ele of each list and add to minHeap
	for listIdx, list := range lists {
		heap.Push(&minHeap, HeapNode{Value: list[0], ListIdx: listIdx, EleIdx: 0})
	}

	result := []int{}

	for minHeap.Len() > 0 {
		nodeOut := heap.Pop(&minHeap).(HeapNode)

		result = append(result, nodeOut.Value)

		// Consider adding next of nodeOut to minHeap
		if nodeOut.EleIdx+1 < len(lists[nodeOut.ListIdx]) {
			heap.Push(&minHeap, HeapNode{
				Value:   lists[nodeOut.ListIdx][nodeOut.EleIdx+1],
				ListIdx: nodeOut.ListIdx,
				EleIdx:  nodeOut.EleIdx + 1,
			})
		}
	}

	return result
}
