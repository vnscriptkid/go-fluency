package grokking

import (
	"testing"
)

func TestKthSmallestInMSortedLists1(t *testing.T) {
	r := kthSmallestInMsortedLists([][]int{{2, 6, 8}, {2, 6, 7}, {1, 3, 4}}, 5)

	if r != 4 {
		panic("Expect 4")
	}
}

func TestKthSmallestInMSortedLists2(t *testing.T) {
	r := kthSmallestInMsortedLists([][]int{{5, 8, 9}, {1, 7, 3}}, 3)

	if r != 7 {
		panic("Expect 7")
	}
}
