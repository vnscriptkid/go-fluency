package grokking

import "container/heap"

type RopeMinHeap []int

func (h RopeMinHeap) Len() int {
	return len(h)
}

func (h RopeMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h RopeMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h RopeMinHeap) Empty() bool {
	return len(h) == 0
}

func (h RopeMinHeap) Peek() any {
	return h[0]
}

func (h *RopeMinHeap) Pop() any {
	old := *h
	oldLen := len(old)

	lastEle := old[oldLen-1]

	*h = old[0 : oldLen-1]

	return lastEle
}

func (h *RopeMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func connectRopes(ropes []int) int {
	minHeap := RopeMinHeap{}
	heap.Init(&minHeap)

	// Push everything to min heap
	for _, rope := range ropes {
		heap.Push(&minHeap, rope)
	}

	// Loop until there's at least 2 in heap
	totalCost := 0

	for minHeap.Len() >= 2 {
		// Take 2 out, add together, add back to min heap
		first := heap.Pop(&minHeap).(int)
		second := heap.Pop(&minHeap).(int)

		final := first + second

		totalCost += final

		heap.Push(&minHeap, final)
	}

	// Return cost
	return totalCost
}
