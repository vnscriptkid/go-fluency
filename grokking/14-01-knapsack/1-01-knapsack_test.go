package grokking

import (
	"strconv"
	"testing"
)

func TestTopDown(t *testing.T) {
	r := topDown([]int{2, 3, 1, 4}, []int{4, 5, 3, 7}, 5)

	if r != 10 {
		panic("Expected " + strconv.Itoa(10) + " but got " + strconv.Itoa(r))
	}

	r = topDown([]int{1, 2, 3, 5}, []int{1, 6, 10, 16}, 7)

	if r != 22 {
		panic("Expected " + strconv.Itoa(22) + " but got " + strconv.Itoa(r))
	}

	r = topDown([]int{1, 2, 3, 5}, []int{1, 6, 10, 16}, 6)

	if r != 17 {
		panic("Expected " + strconv.Itoa(17) + " but got " + strconv.Itoa(r))
	}
}
