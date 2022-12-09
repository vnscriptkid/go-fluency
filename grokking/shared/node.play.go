package main

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

func main() {
	node10 := Node{Value: 10}
	node8 := Node{Value: 8, Next: &node10}
	node6 := Node{Value: 6, Next: &node8}
	node4 := Node{Value: 4, Next: &node6}
	node2 := Node{Value: 2, Next: &node4}

	node2.Print()
}
