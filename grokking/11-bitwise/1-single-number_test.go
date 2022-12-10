package grokking

import "testing"

func TestSingleNumber(t *testing.T) {
	r := singleNumber([]int{1, 4, 2, 1, 3, 2, 3})

	if r != 4 {
		panic("Expect r == 4")
	}
}
