package twoheaps

import (
	"container/heap"
)

func maximumCapital(c int, k int, capitals []int, profits []int) int {

	// 1 , 2 , [1, 2, 2, 3] , [2, 4, 6, 8]
	// minHeap  (2,1), (2,2), (3,3)
	// maxHeap (2,0)
	// curCap: 3
	curCapital := c

	// putting projects in min heap (by capital)
	minHeap := NewMinHeap()
	maxHeap := NewMaxHeap()

	for i, cap := range capitals {
		minHeap.Push(Set{n1: -cap, n2: i})
	}

	// have k projects, handle one by one
	for k > 0 {
		// for each project, take all candidates in the minHeap
		// puting into a maxHeap (by capital)
		for !minHeap.Empty() && curCapital >= capitals[minHeap.Top().(Set).n2] {
			proj := heap.Pop(minHeap).(Set)

			idx := proj.n2

			maxHeap.Push(Set{n1: profits[idx], n2: idx})
		}

		// take best by capital from maxHeap

		if !maxHeap.Empty() {
			proj := heap.Pop(maxHeap).(Set)

			curCapital += proj.n1
		}

		// add to finalProfit
		k--
	}

	return curCapital
}

type Set struct {
	n1 int
	n2 int
}

// MinHeap structure intialization
type MinHeap []Set

// newMinHeap function initializes an instance of MinHeap
func NewMinHeap() *MinHeap {
	min := &MinHeap{}
	heap.Init(min)
	return min
}

// Len function returns the length of MinHeap
func (h MinHeap) Len() int {
	return len(h)
}

// Empty function returns true if empty, false otherwise
func (h MinHeap) Empty() bool {
	return len(h) == 0
}

// Less function compares two elements of MinHeap given their indices
func (h MinHeap) Less(i, j int) bool {
	return h[i].n1 < h[j].n1
}

// Swap function swaps the value of the elements whose indices are given
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MinHeap
func (h MinHeap) Top() interface{} {
	return h[0]
}

// Push function pushes an element into the MinHeap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Set))
}

// Pop function pops the element at the top of the MinHeap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MaxHeap structure intialization
type MaxHeap []Set

// newMaxHeap function initializes an instance of MaxHeap
func NewMaxHeap() *MaxHeap {
	max := &MaxHeap{}
	heap.Init(max)
	return max
}

// Len function returns the length of MaxHeap
func (h MaxHeap) Len() int {
	return len(h)
}

// Empty function returns true if the MaxHeap empty, false otherwise
func (h MaxHeap) Empty() bool {
	return len(h) == 0
}

// Less function compares two elements of MaxHeap given their indices
func (h MaxHeap) Less(i, j int) bool {
	return h[i].n1 < h[j].n1
}

// Swap function swaps the value of the elements whose indices are given
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MaxHeap
func (h MaxHeap) Top() interface{} {
	return h[0]
}

// Push function pushes an element into the MaxHeap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Set))
}

// Pop function pops the element at the top of the MaxHeap
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
