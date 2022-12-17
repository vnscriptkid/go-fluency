package grokking

import "testing"

func TestFindSingleNumbers1(t *testing.T) {
	n1, n2 := findSingleNumbers([]int{1, 4, 2, 1, 3, 5, 6, 2, 3, 5})

	if n1 != 4 || n2 != 6 {
		panic("Ooops")
	}
}

func TestFindSingleNumbers2(t *testing.T) {
	n1, n2 := findSingleNumbers([]int{2, 1, 3, 2})

	if n1 != 1 || n2 != 3 {
		panic("Ooops")
	}
}
