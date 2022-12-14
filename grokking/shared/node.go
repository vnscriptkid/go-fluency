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
