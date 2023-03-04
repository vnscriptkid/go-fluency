package slidingwindow

import (
	"container/list"
	"strconv"
)

func findMaxSlidingWindow(nums []int, windowSize int) []int {
	// [1,2,3,4,5,6,7,8,9,10] | 3
	//  ^
	//      $
	// queue: [2]
	// result: [3]

	result := make([]int, 0)
	queue := NewDeque()
	start := 0
	for end := 0; end < len(nums); end++ {
		curNum := nums[end]

		for !queue.Empty() && curNum > nums[queue.Back()] {
			queue.PopBack()
		}

		queue.PushBack(end)

		if end-start+1 > windowSize {
			// pop front
			if start == queue.Front() {
				queue.PopFront()
			}

			start++
		}

		if end-start+1 == windowSize {
			result = append(result, nums[queue.Front()])
		}
	}
	// Write your code here
	// Your code will replace this placeholder return statement
	return result
}

/* ---------- Template for Dequeue ---------- */

type Deque struct {
	items *list.List
}

// NewDeque is a constructor that will declare and return the Deque type object
func NewDeque() *Deque {
	return &Deque{list.New()}
}

// PushFront will push an element at the front of the dequeue
func (d *Deque) PushFront(val int) {
	d.items.PushFront(val)
}

// PushBack will push an element at the back of the dequeue
func (d *Deque) PushBack(val int) {
	d.items.PushBack(val)
}

// PopFront will pop an element from the front of the dequeue
func (d *Deque) PopFront() int {
	return d.items.Remove(d.items.Front()).(int)
}

// PopBack will pop an element from the back of the dequeue
func (d *Deque) PopBack() int {
	return d.items.Remove(d.items.Back()).(int)
}

// Front will return the element from the front of the dequeue
func (d *Deque) Front() int {
	return d.items.Front().Value.(int)
}

// Back will return the element from the back of the dequeue
func (d *Deque) Back() int {
	return d.items.Back().Value.(int)
}

// Empty will check if the dequeue is empty or not
func (d *Deque) Empty() bool {
	return d.items.Len() == 0
}

// Len will return the length of the dequeue
func (d *Deque) Len() int {
	return d.items.Len()
}

func (d *Deque) Print() string {
	temp := d.items.Front()
	s := "["
	for temp != nil {
		temp2, _ := temp.Value.(int)
		s += strconv.Itoa(temp2)
		temp = temp.Next()
		if temp != nil {
			s += " , "
		}
	}
	s += "]"
	return s
}
