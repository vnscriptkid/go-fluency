package topoSort

import (
	"fmt"
)

// Node type struct
type Node struct {
	value interface{}
	next  *Node
}

// Queue type struct
type Queue struct {
	head *Node
	tail *Node
	size int
}

// Size returns the size of the queue
func (q *Queue) Size() int {
	return q.size
}

// IsEmpty checks whether the queue is empty or not
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Peek returns the top element of the queue
func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return nil
	}
	return q.head.value
}

// Enqueue push an element into the queue
func (q *Queue) Enqueue(value interface{}) {
	temp := &Node{value, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
	} else {
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

// Dequeue remove an element from the queue
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return nil
	}
	value := q.head.value
	q.head = q.head.next
	q.size--
	return value
}

// Print prints the queue on to the console
func (q *Queue) Print() string {
	if q.Size() == 0 {
		return "[]"
	}
	temp := q.head
	out := "["
	for i := 0; i < q.Size(); i++ {
		out = out + "'" + string(temp.value.(rune)) + "'" + ", "
		temp = temp.next
	}
	out = out[0 : len(out)-2]
	out += "]"
	return out
}
