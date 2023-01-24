package shared

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) GetPrinted() string {
	cur := n

	arr := []string{}

	for cur != nil {
		arr = append(arr, strconv.Itoa(cur.Value))
		cur = cur.Next
	}

	return strings.Join(arr, "->")
}

func (n *Node) Print() {
	fmt.Println(n.GetPrinted())
}

func ParseLinkedList(str string) *Node {
	splitted := strings.Split(str, "->")

	var head *Node
	var prev *Node

	for _, strVal := range splitted {
		if val, err := strconv.Atoi(strVal); err == nil {
			node := Node{Value: val}

			if prev == nil {
				head = &node
			} else {
				prev.Next = &node
			}

			prev = &node
		} else {
			// end of list
		}
	}

	return head
}
