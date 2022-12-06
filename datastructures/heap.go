package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

// ////////////////////////////
// Implementing sort interface
// ////////////////////////////
func (iHeap IntHeap) Len() int {
	return len(iHeap)
}

// Min Heap
func (iHeap IntHeap) Less(i, j int) bool {
	return iHeap[i] < iHeap[j]
}

func (iHeap IntHeap) Swap(i, j int) {
	iHeap[i], iHeap[j] = iHeap[j], iHeap[i]
}

// ////////////////////////////
// Implementing heap interface
// ////////////////////////////
func (iHeap *IntHeap) Push(x interface{}) {
	*iHeap = append(*iHeap, x.(int))
}

func (iHeap *IntHeap) Pop() interface{} {
	prev := *iHeap

	lastEle := prev[len(prev)-1]

	// Extract out last ele
	*iHeap = prev[0 : len(prev)-1]

	return lastEle
}

func main() {
	intHeap := &IntHeap{2, 5, 4, 1, 8, 7}

	heap.Init(intHeap)
	heap.Push(intHeap, 10)

	fmt.Println("Min now: ", (*intHeap)[0])

	fmt.Println("Pop one: ", heap.Pop(intHeap).(int))
	fmt.Println("Pop one: ", heap.Pop(intHeap).(int))
	fmt.Println("Pop one: ", heap.Pop(intHeap).(int))

	fmt.Println("Min now: ", (*intHeap)[0])
}
