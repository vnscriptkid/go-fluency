package grokking

import (
	"container/heap"
	"sort"
)

type RoomMinHeap [][2]int

func (h RoomMinHeap) Len() int {
	return len(h)
}

func (h RoomMinHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h RoomMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h RoomMinHeap) Empty() bool {
	return len(h) == 0
}

func (h RoomMinHeap) Peek() any {
	return h[0]
}

func (h *RoomMinHeap) Pop() any {
	prev := *h
	prevLen := len(prev)

	lastEle := prev[prevLen-1]

	*h = prev[0 : prevLen-1]

	return lastEle
}

func (h *RoomMinHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func minMeetingRooms(meetings [][2]int) int {

	// Sort meetings by startTime in asc order
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	roomMinHeap := &RoomMinHeap{}
	heap.Init(roomMinHeap)

	minMeetingRooms := 0
	// [][2]int{{4, 5}, {2, 3}, {2, 4}, {3, 5}},

	// 1  2  3  4  5  6
	//    [  ]
	//    {     }
	//       {     }
	//          [  ]

	for _, meeting := range meetings {
		heap.Push(roomMinHeap, meeting)

		// Remove all the meetings that has ended
		for !roomMinHeap.Empty() && roomMinHeap.Peek().([2]int)[1] <= meeting[0] {
			heap.Pop(roomMinHeap)
		}

		// How many meetings that are happening now
		if roomMinHeap.Len() > minMeetingRooms {
			minMeetingRooms = roomMinHeap.Len()
		}
	}

	return minMeetingRooms
}
