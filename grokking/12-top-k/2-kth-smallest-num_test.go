package grokking

import "testing"

func TestFindKthSmallestNum1(t *testing.T) {
	r := findKthSmallestNum([]int{1, 5, 12, 2, 11, 5}, 3)

	if r != 5 {
		panic("Oops")
	}
}

func TestFindKthSmallestNum2(t *testing.T) {
	r := findKthSmallestNum([]int{1, 5, 12, 2, 11, 5}, 4)

	if r != 5 {
		panic("Oops")
	}
}

func TestFindKthSmallestNum3(t *testing.T) {
	r := findKthSmallestNum([]int{5, 12, 11, -1, 12}, 3)

	if r != 11 {
		panic("Oops")
	}
}
