package grokking

import (
	"testing"
)

func TestFindMissingNumber1(t *testing.T) {
	r := findMissingNum([]int{4, 0, 3, 1})

	if r != 2 {
		panic("expected 2")
	}
}

func TestFindMissingNumber2(t *testing.T) {
	r := findMissingNum([]int{8, 3, 5, 2, 4, 6, 0, 1})

	if r != 7 {
		panic("expected 7")
	}
}

func TestFindMissingNumber3(t *testing.T) {
	r := findMissingNum([]int{2, 0, 3, 1})

	if r != 4 {
		panic("expected 4")
	}
}
