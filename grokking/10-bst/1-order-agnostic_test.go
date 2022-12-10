package grokking

import "testing"

func TestOrderAgnosticBst(t *testing.T) {
	arr := []int{4, 6, 10}
	r := orderAgnosticBst(arr, 10)

	if r != 2 {
		panic("Expect r == 2")
	}

	// TC 2
	r = orderAgnosticBst([]int{1, 2, 3, 4, 5, 6, 7}, 5)

	if r != 4 {
		panic("Expect r == 4")
	}

	// TC 2
	r = orderAgnosticBst([]int{10, 6, 4}, 10)

	if r != 0 {
		panic("Expect r == 0")
	}
}
