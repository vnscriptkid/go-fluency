package grokking

import "testing"

func TestMaxSumSubArr(t *testing.T) {
	r := maxSumSubArr([]int{2, 1, 5, 1, 3, 2}, 3)

	if r != 9 {
		panic("Wrong result")
	}
}
