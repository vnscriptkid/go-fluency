package main

import (
	"fmt"
	"time"
)

func RowTraverse() int {
	rows := 1000
	cols := 1000
	matrix := make([][]byte, rows)

	for i := range matrix {
		matrix[i] = make([]byte, cols)
	}

	var ctr int
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0x00 {
				ctr++
			}
		}
	}
	return ctr
}

func ColumnTraverse() int {
	rows := 1000
	cols := 1000
	matrix := make([][]byte, rows)

	for i := range matrix {
		matrix[i] = make([]byte, cols)
	}

	var ctr int
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if matrix[row][col] == 0x00 {
				ctr++
			}
		}
	}
	return ctr
}

type ListNode struct {
	Value int
	Next  *ListNode
}

func LinkedListTraverse() {

	// Creating a linked list of 1 million nodes.
	var head, current *ListNode
	for i := 0; i < 1000000; i++ {
		newNode := &ListNode{Value: i}
		if head == nil {
			head = newNode
			current = head
		} else {
			current.Next = newNode
			current = newNode
		}
	}

	// Traversing the linked list.
	var sum int
	current = head
	for current != nil {
		sum += current.Value
		current = current.Next
	}
	fmt.Println("Sum:", sum)
}

func main() {
	start := time.Now()
	RowTraverse()
	elapsed := time.Since(start)
	println("[RowTraverse] Elapsed in Nanoseconds: ", elapsed.Nanoseconds())

	start = time.Now()
	ColumnTraverse()
	elapsed = time.Since(start)
	println("[ColumnTraverse] Elapsed in Nanoseconds: ", elapsed.Nanoseconds())

	start = time.Now()
	LinkedListTraverse()
	elapsed = time.Since(start)
	println("[LinkedListTraverse] Elapsed in Nanoseconds: ", elapsed.Nanoseconds())

}
