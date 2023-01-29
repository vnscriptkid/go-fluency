package grokking

import (
	"container/heap"
	"math"
)

type NumberNode struct {
	Value  int
	NumIdx int
	ArrIdx int
}

type NumberMinHeap []NumberNode

func (h NumberMinHeap) Len() int {
	return len(h)
}

func (h NumberMinHeap) Less(i, j int) bool {
	return h[i].Value < h[j].Value
}

func (h NumberMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h NumberMinHeap) Empty() bool {
	return len(h) == 0
}

func (h NumberMinHeap) Peek() any {
	return h[0]
}

func (h *NumberMinHeap) Push(x any) {
	*h = append(*h, x.(NumberNode))
}

func (h *NumberMinHeap) Pop() any {
	old := *h
	oldLen := len(old)
	lastEle := old[oldLen-1]

	*h = old[0 : oldLen-1]

	return lastEle
}

func smallestNumRange(sortedArr [][]int) [2]int {
	minHeap := NumberMinHeap{}
	k := len(sortedArr)

	heap.Init(&minHeap)

	// Keep track of cur max
	curMax := math.MinInt
	// Keep track of cur smallest range
	curSmallestRange := [2]int{0, math.MaxInt}
	// Add arrs[i][0] to min heap
	for arrIdx, arr := range sortedArr {
		heap.Push(&minHeap, NumberNode{
			Value:  arr[0],
			NumIdx: 0,
			ArrIdx: arrIdx,
		})

		if arr[0] > curMax {
			curMax = arr[0]
		}
	}

	// Take out one by one until minHeap.Len() < k
	for minHeap.Len() == k {
		// Decide to update cur smallest range
		peek := minHeap.Peek().(NumberNode)

		if curMax-peek.Value < curSmallestRange[1]-curSmallestRange[0] {
			curSmallestRange[0], curSmallestRange[1] = peek.Value, curMax
		}

		out := heap.Pop(&minHeap).(NumberNode)
		// Add next one to heap, update cur max
		curArr := sortedArr[out.ArrIdx]

		if out.NumIdx+1 < len(curArr) {
			newValue := curArr[out.NumIdx+1]

			heap.Push(&minHeap, NumberNode{
				Value:  newValue,
				NumIdx: out.NumIdx + 1,
				ArrIdx: out.ArrIdx,
			})

			if newValue > curMax {
				curMax = newValue
			}
		}
	}

	return curSmallestRange
}
