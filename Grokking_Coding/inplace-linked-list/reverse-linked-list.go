package inplacelinkedlist1

import (
	"fmt"
	"strings"
)

func reverse(head *EduLinkedListNode) *EduLinkedListNode {
	cur := head

	// nil <-1   2 -> 3 -> nil
	//  $    ^

	var prev *EduLinkedListNode = nil

	for cur != nil {
		temp := cur.next

		cur.next = prev

		prev = cur
		cur = temp
	}

	return prev
}

type EduLinkedListNode struct {
	data int
	next *EduLinkedListNode
}

func NewLinkedListNode(data int, next *EduLinkedListNode) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = next
	return node
}

func InitLinkedListNode(data int) *EduLinkedListNode {
	node := new(EduLinkedListNode)
	node.data = data
	node.next = nil
	return node
}

type EduLinkedList struct {
	head *EduLinkedListNode
}

/*
InsertNodeAtHead method will insert a LinkedListNode at head

	of a linked list.
*/
func (l *EduLinkedList) InsertNodeAtHead(node *EduLinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

/*
	CreateLinkedList method will create the linked list using

the given integer array with the help of InsertAthead method.
*/
func (l *EduLinkedList) CreateLinkedList(lst []int) {
	for i := len(lst) - 1; i >= 0; i-- {
		newNode := InitLinkedListNode(lst[i])
		l.InsertNodeAtHead(newNode)
	}
}

// DisplayLinkedList method will display the elements of linked list.
func (l *EduLinkedList) DisplayLinkedList() {
	temp := l.head
	fmt.Print("[")
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
}

func (l *EduLinkedList) GetLinkedListPrinted() string {
	temp := l.head
	var sBuilder strings.Builder
	sBuilder.WriteString("[")

	for temp != nil {
		// fmt.Print(temp.data)
		sBuilder.WriteString(fmt.Sprint(temp.data))

		temp = temp.next
		if temp != nil {
			sBuilder.WriteString(", ")
		}
	}
	sBuilder.WriteString("]")

	return sBuilder.String()
}
