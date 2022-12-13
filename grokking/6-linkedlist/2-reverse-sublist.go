package grokking

import "github.com/vnscriptkid/go-fluency/grokking/shared"

func reverseSublist(h *shared.Node, s int, e int) *shared.Node {
	// 1 -> 2 -> 3 -> 4 -> 5 -> nil
	//      ^         $

	// 1 -> 2 <- 3 <- 4 -> 5 -> nil

	var p *shared.Node = nil
	c := h

	i := 1

	for ; c != nil && i < s; i++ {
		p = c
		c = c.Next
	}

	// i === s: cur now is at start of sublist
	pos := p
	hos := c

	// start reverse from cur
	for i <= e && c != nil {
		t := c.Next

		c.Next = p

		// Update p, c
		p = c
		c = t

		i++
	}

	// i > e: cur now is at right after sublist
	pos.Next = p
	hos.Next = c

	return h
}
