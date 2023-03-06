package topk

import "container/heap"

type KthLargest struct {
	// Write your code here
	minHeap *MinHeap
	k       int
}

// KthLargestInit is a constructor to initialize heap and store values in it
func (this *KthLargest) KthLargestInit(k int, nums []int) {
	this.minHeap = newMinHeap()
	this.k = k

	for _, num := range nums {
		this.Add(num)
	}
}

// Add function adds element in the heap
func (this *KthLargest) Add(value int) int {
	heap.Push(this.minHeap, value)

	if this.minHeap.Len() > this.k {
		heap.Pop(this.minHeap)
	}
	// Write your code here
	// Your code will replace this placeholder return statement
	return value
}

// Returns the Kth largest element from the heap
func (this *KthLargest) ReturnKthLargest() int {
	// Write your code here
	// Your code will replace this placeholder return statement
	return this.minHeap.Top()
}

// MinHeap structure initialization
type MinHeap []int

// newMinHeap function initializes an instance of MinHeap
func newMinHeap() *MinHeap {
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

// Less function compares two elements of the MinHeap given their indexes
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap function swaps the values of the elements whose indices are given
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MinHeap
func (h MinHeap) Top() int {
	return h[0]
}

// Push function pushes an element into the MinHeap
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop function pops the element at the top of the MinHeap
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
