package grokking

import (
	"container/heap"
	"math"
)

type MaxHeapPoint [][2]int

func (h *MaxHeapPoint) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *MaxHeapPoint) Pop() any {
	prev := *h

	lastPoint := prev[len(prev)-1]

	*h = prev[:len(prev)-1]

	return lastPoint
}

func (h MaxHeapPoint) Len() int {
	return len(h)
}

func (h MaxHeapPoint) Less(i, j int) bool {
	return DistanceToOrigin(h[j]) < DistanceToOrigin(h[i])
}

func DistanceToOrigin(point [2]int) float64 {
	return math.Sqrt(float64(point[0]*point[0] + point[1]*point[1]))
}

func (h MaxHeapPoint) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func findKClosestPoints(points [][2]int, k int) [][2]int {
	maxHeapPoint := MaxHeapPoint{}
	heap.Init(&maxHeapPoint)

	// Add first k points to max heap
	for i := 0; i < k; i++ {
		heap.Push(&maxHeapPoint, points[i])
	}

	// From now on, be selective, push and pop if cur point is farther than cur farthest point
	for i := k; i < len(points); i++ {
		if DistanceToOrigin(points[i]) < DistanceToOrigin(maxHeapPoint[0]) {
			heap.Pop(&maxHeapPoint)
			heap.Push(&maxHeapPoint, points[i])
		}
	}

	return maxHeapPoint
}
