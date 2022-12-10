package grokking

import "github.com/vnscriptkid/go-fluency/grokking/shared"

func reverse(head *shared.Node) *shared.Node {
	// nil<-2 <- 4    6 -> 8 -> 10 -> nil
	// 		     p    c    t

	var prev *shared.Node = nil
	cur := head

	for cur != nil {
		// Save next so not to loose track
		temp := cur.Next

		// Point cur to prev
		cur.Next = prev

		// Update prev, then cur
		prev = cur
		cur = temp
	}

	return prev
}

func main() {
	node10 := shared.Node{Value: 10}
	node8 := shared.Node{Value: 8, Next: &node10}
	node6 := shared.Node{Value: 6, Next: &node8}
	node4 := shared.Node{Value: 4, Next: &node6}
	node2 := shared.Node{Value: 2, Next: &node4}

	newHead := reverse(&node2)

	(*newHead).Print()
}
